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
func (linkedlist *LinkedList) PopFront() interface{} {
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
