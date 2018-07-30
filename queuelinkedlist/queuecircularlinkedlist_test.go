package queuecircularlinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestQueueWithCircularLinkedlist correct amount of Node include in list
func TestQueueWithCircularLinkedlist(t *testing.T) {
	queue := NewQueue()

	node1 := NewNode(1)
	node2 := NewNode(2)

	queue.Enqueue(node1)
	queue.Enqueue(node2)

	item, _ := queue.Dequeue()
	item2, _ := queue.Dequeue()
	_, err := queue.Dequeue()

	assert.Equal(t, queue.Size(), 0)
	assert.Equal(t, item.data, 1)
	assert.Equal(t, item2.data, 2)
	assert.True(t, err != nil)
}
