package linkedlist

import (
	"fmt"
)

// Node for linked list
type Node struct {
	Data interface{}
	next *Node
}

// NewNode is creating node for linked list
func NewNode(data interface{}) *Node {
	return &Node{Data: data, next: nil}
}

// LinkedList is to store head
type LinkedList struct {
	head *Node
}

// NewLinkedList is creating linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// PushFront to the LinkedList
func (linkedlist *LinkedList) PushFront(node *Node) {
	tempHead := linkedlist.head
	linkedlist.head = node

	if tempHead == nil {
		return
	}

	linkedlist.head.next = tempHead
}

// PopFront to the LinkedList
func (linkedlist *LinkedList) PopFront() *Node {
	tempHead := linkedlist.head

	if tempHead == nil {
		return nil
	}

	if tempHead.next == nil {
		linkedlist.head = nil
	}

	linkedlist.head = tempHead.next
	return tempHead
}

// Size to the LinkedList
func (linkedlist *LinkedList) Size() int {
	temp := linkedlist.head
	counter := 0

	for temp != nil {
		counter++
		temp = temp.next
	}

	return counter
}

// Empty is to check whether linkedlist is empty or not
func (linkedlist LinkedList) Empty() bool {
	return linkedlist.head == nil
}

// PushBack is to push at the end
func (linkedlist *LinkedList) PushBack(node *Node) {

	if linkedlist.head == nil {
		linkedlist.head = node
		return
	}

	temp := linkedlist.head
	for temp != nil {
		if temp.next == nil {
			temp.next = node
			return
		}
		temp = temp.next
	}
}

// PopBack is to pop node at the end
func (linkedlist LinkedList) PopBack() *Node {

	if linkedlist.head == nil {
		return nil
	}

	temp := linkedlist.head
	for temp != nil {
		nextNode := temp.next
		if nextNode.next == nil {
			temp.next = nil
			return nextNode
		}
		temp = nextNode
	}

	return nil
}

// Back is to get node at the end
func (linkedlist LinkedList) Back() *Node {

	if linkedlist.head == nil {
		return nil
	}

	temp := linkedlist.head
	for temp != nil {
		nextNode := temp.next
		if nextNode.next == nil {
			return nextNode
		}
		temp = nextNode
	}

	return nil
}

// Front is to get node at the front
func (linkedlist LinkedList) Front() *Node {
	return linkedlist.head
}

// ValueAt is to get node at the nth element
func (linkedlist LinkedList) ValueAt(n int) *Node {
	current := 0
	temp := linkedlist.head
	for temp != nil {
		if current == n {
			return temp
		}
		current++
		temp = temp.next
	}
	return nil
}

// PrintList to console
func (linkedlist *LinkedList) PrintList() {
	if linkedlist.head == nil {
		fmt.Println("No list to print")
		return
	}
	temp := linkedlist.head
	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.next
	}
}
