package concurrent

import (
	"sync"
)

/**** YOU CANNOT MODIFY ANY OF THE FOLLOWING INTERFACES/TYPES ********/
type Task interface{}

type DEQueue interface {
	PushBottom(task Task)
	IsEmpty() bool //returns whether the queue is empty
	PopTop() Task
	PopBottom() Task
	Size() int
}

/******** DO NOT MODIFY ANY OF THE ABOVE INTERFACES/TYPES *********************/

type queueNode struct {
	task Task
	next *queueNode
	prev *queueNode
}

func newQueueNode(task Task) *queueNode {
	return &queueNode{task: task, next: nil, prev: nil}
}

type unBoundedDEQueue struct {
	head *queueNode
	tail *queueNode
	lock *sync.Mutex
	size int
}

// NewUnBoundedDEQueue returns an empty UnBoundedDEQueue
func NewUnBoundedDEQueue() DEQueue {
	return &unBoundedDEQueue{lock: &sync.Mutex{}, head: nil, tail: nil}
}

func (q *unBoundedDEQueue) PushBottom(task Task) {
	q.lock.Lock()
	defer q.lock.Unlock()
	node := newQueueNode(task)
	if q.head == nil || q.tail == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		node.prev = q.tail
		q.tail = node
	}
	q.size++
}

func (q *unBoundedDEQueue) PopBottom() Task {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.head == nil || q.tail == nil {
		return nil
	} else if q.head == q.tail {
		task := q.tail.task
		q.head = nil
		q.tail = nil
		q.size--
		return task
	} else {
		task := q.tail.task
		q.tail = q.tail.prev
		q.tail.next = nil
		q.size--
		return task
	}
}

func (q *unBoundedDEQueue) PopTop() Task {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.head == nil || q.tail == nil {
		return nil
	} else if q.head == q.tail {
		task := q.head.task
		q.head = nil
		q.tail = nil
		q.size--
		return task
	} else {
		task := q.head.task
		q.head = q.head.next
		q.head.prev = nil
		q.size--
		return task
	}
}

func (q *unBoundedDEQueue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.head == nil || q.tail == nil
}

func (q *unBoundedDEQueue) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size
}
