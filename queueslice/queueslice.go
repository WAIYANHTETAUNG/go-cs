package queueslice

import (
	"fmt"
)

// QueueSlice is struct to hold information of Q
type QueueSlice struct {
	Front int
	Rear  int
	Q     []int
}

// NewQueue is create Queue
func NewQueue() *QueueSlice {
	return &QueueSlice{}
}

// Enqueue is to add item at the end
func (queueslice *QueueSlice) Enqueue(item int) {
	queueslice.Q = append(queueslice.Q, item)
}

// Dequeue is to pop item at the first
func (queueslice *QueueSlice) Dequeue() int {
	length := len(queueslice.Q)
	if length == 0 {
		fmt.Println("Q is empty, Can't dequeue.")
		return -1
	}
	var firstItem int

	firstItem, queueslice.Q = queueslice.Q[0], queueslice.Q[1:]
	return firstItem
}

// QueueFront is to get item at the first without mutating Q
func (queueslice *QueueSlice) QueueFront() int {
	length := len(queueslice.Q)
	if length == 0 {
		fmt.Println("Q is empty")
		return -1
	}

	return queueslice.Q[0]
}

// QueueRear is to get item at the end without mutating Q
func (queueslice *QueueSlice) QueueRear() int {
	length := len(queueslice.Q)
	if length == 0 {
		fmt.Println("Q is empty")
		return -1
	}

	return queueslice.Q[length-1]
}

func (queueslice *QueueSlice) isEmpty() bool {
	return len(queueslice.Q) == 0
}

func (queueslice *QueueSlice) Size() int {
	return len(queueslice.Q)
}
