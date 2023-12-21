package concurrent

import (
	"math/rand"
	"sync"
	"time"
)

// NewWorkBalancingExecutor returns an ExecutorService that is implemented using the work-balancing algorithm.
// @param capacity - The number of goroutines in the pool
// @param threshold - The number of items that a goroutine in the pool can
// grab from the executor in one time period. For example, if threshold = 10
// this means that a goroutine can grab 10 items from the executor all at
// once to place into their local queue before grabbing more items. It's
// not required that you use this parameter in your implementation.
// @param thresholdBalance - The threshold used to know when to perform
// balancing. Remember, if two local queues are to be balanced the
// difference in the sizes of the queues must be greater than or equal to
// thresholdBalance. You must use this parameter in your implementation.
func NewWorkBalancingExecutor(capacity, thresholdQueue, thresholdBalance int) ExecutorService {
	globalQueue := NewUnBoundedDEQueue()
	localQueueList := make([]DEQueue, capacity)
	execService := &ExecService{
		capacity:       capacity,
		threshold:      thresholdQueue,
		globalQueue:    globalQueue,
		localQueueList: localQueueList,
		done:           false,
		wg:             &sync.WaitGroup{},
	}

	worker := func(execService *ExecService, workerId int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			// BREAKING CONDITION
			if execService.done &&
				execService.globalQueue.IsEmpty() &&
				execService.localQueueList[workerId].IsEmpty() {
				break
			}

			// POLL FROM GLOBAL QUEUE FOR MORE TASKS
			for i := 0; i < execService.threshold; i++ {
				f := execService.globalQueue.PopTop()
				if f == nil {
					break
				}
				execService.localQueueList[workerId].PushBottom(f)
			}

			// LOAD BALANCING ALGORITHM
			loadBalancer := func() {
				// RANDOM NUMBER GENERATOR
				random := func(max int, except int) int {
					for {
						rand.Seed(time.Now().UnixNano())
						n := rand.Intn(max)
						if n != except {
							return n
						}
					}
				}

				sizeLocal := execService.localQueueList[workerId].Size()
				if random(sizeLocal+1, -1) == sizeLocal {
					victim := random(execService.capacity, workerId)
					sizeVictim := execService.localQueueList[victim].Size()

					minQ := execService.localQueueList[victim]
					maxQ := execService.localQueueList[workerId]

					if sizeLocal < sizeVictim {
						minQ = execService.localQueueList[workerId]
						maxQ = execService.localQueueList[victim]
					}

					if maxQ.Size()-minQ.Size() > thresholdBalance {
						for i := 0; i < (maxQ.Size()-minQ.Size())/2; i++ {
							f := maxQ.PopTop()
							minQ.PushBottom(f)
						}
					}
				}
			}

			// PERFORM LOAD BALANCING
			loadBalancer()

			// WORK ON TASKS IN THE LOCAL QUEUE
			for !execService.localQueueList[workerId].IsEmpty() {
				f_ := execService.localQueueList[workerId].PopBottom()
				if f_ == nil {
					break
				}

				f := f_.(*future)
				if task, ok := f.Task.(interface{ Call() interface{} }); ok {
					f.Promise <- task.Call()
				} else {
					task := f.Task.(interface{ Run() })
					task.Run()
					f.Promise <- nil
				}
				close(f.Promise)

			}

		}
	}

	execService.wg.Add(1)
	{
		go func(e *ExecService) {
			var wg sync.WaitGroup
			wg.Add(capacity)

			for i := 0; i < capacity; i++ {
				localQueueList[i] = NewUnBoundedDEQueue()
			}

			// SPAWN WORKERS
			for i := 0; i < capacity; i++ {
				go worker(e, i, &wg)
			}

			// WAIT FOR EVERYONE TO FINISH WORKING
			wg.Wait()

			// NOTIFY SHUTDOWN ALL TASKS ARE COMPLETED
			execService.wg.Done()
		}(execService)
	}

	return execService
}
