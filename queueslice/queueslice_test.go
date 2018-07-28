package queueslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	queueslice := NewQueue()
	queueslice.Enqueue(1)
	queueslice.Enqueue(2)

	assert.False(t, queueslice.isEmpty())
}

func TestDequeue(t *testing.T) {
	queueslice := NewQueue()
	queueslice.Enqueue(1)
	queueslice.Enqueue(2)
	queueslice.Enqueue(3)
	queueslice.Enqueue(4)

	assert.Equal(t, queueslice.Size(), 4)

	dequeuedItem := queueslice.Dequeue()
	assert.Equal(t, dequeuedItem, 1)
	assert.Equal(t, queueslice.Size(), 3)
}

func TestQueueFront(t *testing.T) {
	queueslice := NewQueue()
	queueslice.Enqueue(1)
	queueslice.Enqueue(2)
	queueslice.Enqueue(3)
	queueslice.Enqueue(4)

	assert.Equal(t, queueslice.QueueFront(), 1)
}

func TestQueueRear(t *testing.T) {
	queueslice := NewQueue()
	queueslice.Enqueue(1)
	queueslice.Enqueue(2)
	queueslice.Enqueue(3)
	queueslice.Enqueue(4)

	assert.Equal(t, queueslice.QueueRear(), 4)
}
