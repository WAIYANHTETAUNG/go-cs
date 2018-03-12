package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCorrectAmountOfNodeAttach correct amount of Node include in list
func TestCorrectAmountOfNodeAttach(t *testing.T) {
	list := NewLinkedList()
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)

	list.PushFront(node1)
	list.PushFront(node2)
	list.PushFront(node3)
	list.PushFront(node4)

	assert.Equal(t, list.Size(), 4)
}

func TestCorrectNthElementReturn(t *testing.T) {
	list := NewLinkedList()
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)

	list.PushFront(node1)
	list.PushFront(node2)
	list.PushFront(node3)
	list.PushFront(node4)

	assert.Equal(t, list.ValueAt(0).Data, 4)
}
