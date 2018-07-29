package hashwithchaining

import (
	"fmt"
)

// Node for linked list
type Node struct {
	key  string
	val  int
	Next *Node
}

// NewNode is creating node for linked list
func NewNode(key string, val int) *Node {
	return &Node{key: key, val: val, Next: nil}
}

// LinkedList is to store head
type LinkedList struct {
	Head *Node
}

// NewLinkedList is creating linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// PushFront to the LinkedList
func (linkedlist *LinkedList) PushFront(node *Node) {
	tempHead := linkedlist.Head
	linkedlist.Head = node

	if tempHead == nil {
		return
	}

	linkedlist.Head.Next = tempHead
}

// PopFront to the LinkedList
func (linkedlist *LinkedList) PopFront() *Node {
	tempHead := linkedlist.Head

	if tempHead == nil {
		return nil
	}

	if tempHead.Next == nil {
		linkedlist.Head = nil
	}

	linkedlist.Head = tempHead.Next
	return tempHead
}

// Size to the LinkedList
func (linkedlist *LinkedList) Size() int {
	temp := linkedlist.Head
	counter := 0

	for temp != nil {
		counter++
		temp = temp.Next
	}

	return counter
}

// Empty is to check whether linkedlist is empty or not
func (linkedlist LinkedList) Empty() bool {
	return linkedlist.Head == nil
}

// PushBack is to push at the end
func (linkedlist *LinkedList) PushBack(node *Node) {

	if linkedlist.Head == nil {
		linkedlist.Head = node
		return
	}

	temp := linkedlist.Head
	for temp != nil {
		if temp.Next == nil {
			temp.Next = node
			return
		}
		temp = temp.Next
	}
}

// PopBack is to pop node at the end
func (linkedlist LinkedList) PopBack() *Node {

	if linkedlist.Head == nil {
		return nil
	}

	temp := linkedlist.Head
	for temp != nil {
		nextNode := temp.Next
		if nextNode.Next == nil {
			temp.Next = nil
			return nextNode
		}
		temp = nextNode
	}

	return nil
}

// Back is to get node at the end
func (linkedlist LinkedList) Back() *Node {

	if linkedlist.Head == nil {
		return nil
	}

	temp := linkedlist.Head
	for temp != nil {
		nextNode := temp.Next
		if nextNode.Next == nil {
			return nextNode
		}
		temp = nextNode
	}

	return nil
}

// Front is to get node at the front
func (linkedlist LinkedList) Front() *Node {
	return linkedlist.Head
}

// ValueAt is to get node at the nth element
func (linkedlist LinkedList) ValueAt(n int) *Node {
	current := 0
	temp := linkedlist.Head
	for temp != nil {
		if current == n {
			return temp
		}
		current++
		temp = temp.Next
	}
	return nil
}

// Erase is to delete node at the given nth index
func (linkedlist LinkedList) Erase(n int) *Node {
	current := 0
	temp := linkedlist.Head
	for temp != nil {
		if current+1 == n {
			temp.Next = temp.Next.Next
			return temp
		}
		current++
		temp = temp.Next
	}
	return nil
}

// Insert is to insert node at the given nth index
func (linkedlist LinkedList) Insert(n int, node *Node) *Node {
	current := 0
	temp := linkedlist.Head
	for temp != nil {
		if current+1 == n {
			node.Next = temp.Next
			temp.Next = node
			return temp
		}
		current++
		temp = temp.Next
	}
	return nil
}

// PrintList to console
func (linkedlist *LinkedList) PrintList() {
	if linkedlist.Head == nil {
		fmt.Println("No list to print")
		return
	}
	temp := linkedlist.Head
	for temp != nil {
		temp = temp.Next
	}
}
