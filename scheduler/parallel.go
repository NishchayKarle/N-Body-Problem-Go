package scheduler

import (
	"fmt"
	"os"
	"proj3/concurrent"
	"proj3/nbody"
)

type NbodyTask struct {
	id              int
	bodies          []*nbody.Body
	dt              float32
	numBodies       int
	softeningFactor float32
	typeOfTask      string
}

func NewNbodyTask(id int, bodies []*nbody.Body, dt float32,
	numBodies int, softeningFactor float32, typeOfTask string) concurrent.Runnable {
	return &NbodyTask{
		id:              id,
		bodies:          bodies,
		dt:              dt,
		numBodies:       numBodies,
		softeningFactor: softeningFactor,
		typeOfTask:      typeOfTask,
	}
}

func (task *NbodyTask) Run() {
	if task.typeOfTask == "ComputeForce" {
		// COMPUTE INTERBODY FORCES
		nbody.ComputeBodyForce(
			task.id,
			task.bodies,
			task.dt,
			task.numBodies,
			task.softeningFactor,
		)
	} else if task.typeOfTask == "IntegratePositions" {
		// INTEGRATE POSITIONS
		nbody.IntegratePositions(
			task.id,
			task.bodies,
			task.numBodies,
			task.dt,
		)
	} else if task.typeOfTask == "InitPositionsAndVelocities" {
		nbody.InitPositionsAndVelocities(
			task.id,
			task.bodies,
			task.numBodies,
		)
	}
}

func RunParallel(numBodies, iterations int, dt float32, record string, threads int, mode string) {
	bodies := make([]*nbody.Body, numBodies)

	var file *os.File
	var err error

	// WRITE POSITIONS
	if record == "yes" {
		file, err = os.Create("../scheduler/parallel/nbody.csv")
		if err != nil {
			fmt.Println("ERROR WHEN OPENING FILE \"nbody.csv\"")
			panic(err)
		}
		defer file.Close()
	}

	var executor concurrent.ExecutorService
	if mode == "ws" {
		executor = concurrent.NewWorkStealingExecutor(threads, numBodies/(10*threads))
	} else {
		executor = concurrent.NewWorkBalancingExecutor(threads, numBodies/(10*threads), numBodies/(50*threads))
	}

	futures := make([]concurrent.Future, numBodies)
	for i := 0; i < numBodies; i++ {
		futures[i] = executor.Submit(NewNbodyTask(i, bodies, dt, numBodies, 1e-4, "InitPositionsAndVelocities"))
	}

	for _, f := range futures {
		f.Get()
	}

	for iter := 0; iter < iterations+1; iter++ {
		if record == "yes" {
			if iterations%(iterations/10) == 0 {
				nbody.ParticlePositionsToCSV(file, iter/(iterations/10), bodies, numBodies) // WRITE POSITIONS AFTER ITERATION
			}
		}

		for i := 0; i < numBodies; i++ {
			futures[i] = executor.Submit(NewNbodyTask(i, bodies, dt, numBodies, 1e-4, "ComputeForce"))
		}

		for _, f := range futures {
			f.Get()
		}

		for i := 0; i < numBodies; i++ {
			futures[i] = executor.Submit(NewNbodyTask(i, bodies, dt, numBodies, 1e-4, "IntegratePositions"))
		}

		for _, f := range futures {
			f.Get()
		}
	}
	executor.Shutdown()
}
