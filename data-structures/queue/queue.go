package queue

import "errors"

type Queue struct {
	slice []int
}

var (
	ErrEmptyQueue = errors.New("Queue is empty")
)

// Enqueue Adds an Item to the Queue
func (q *Queue) Enqueue(i int) {
	// Add New Value to End of Slice
	q.slice = append(q.slice, i)
}

// Dequeue Returns an Item from the Queue
func (q *Queue) Dequeue() (int, error) {
	// Check if Queue is Empty
	if q.Len() == 0 {
		return 0, ErrEmptyQueue
	}

	// The Front of the Queue is at Index 0
	val := q.slice[0]

	// Removing Index 0 from slice
	q.slice = q.slice[1:]

	return val, nil
}

// Len Gives the Length of the Queue
func (q *Queue) Len() int {
	return len(q.slice)
}
