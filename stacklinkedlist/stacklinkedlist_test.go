package stacklinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPush
func TestPush(t *testing.T) {
	list := NewLinkedList()
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)

	list.Push(node1)
	list.Push(node2)
	list.Push(node3)
	list.Push(node4)

	assert.Equal(t, list.Size(), 4)
}

// TestPushPop
func TestPop(t *testing.T) {
	list := NewLinkedList()
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)

	list.Push(node1)
	list.Push(node2)
	list.Push(node3)
	list.Push(node4)

	list.Pop()

	assert.Equal(t, list.Size(), 3)
}

// TestPushPop
func TestEmpty(t *testing.T) {
	list := NewLinkedList()
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)

	list.Push(node1)
	list.Push(node2)
	list.Push(node3)
	list.Push(node4)

	list.Pop()

	assert.False(t, list.IsEmpty())
}
