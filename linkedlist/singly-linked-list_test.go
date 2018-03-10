package linkedlist

import (
	"reflect"
	"testing"
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

	reflect.DeepEqual(list.Size(), 4)
}
