package main

import (
	"fmt"

	"./linkedlist"
)

func main() {
	list := linkedlist.NewLinkedList()
	node1 := linkedlist.NewNode(1)
	node2 := linkedlist.NewNode(2)
	node3 := linkedlist.NewNode(3)
	node4 := linkedlist.NewNode(4)
	node5 := linkedlist.NewNode(5)

	list.PushFront(node1)
	list.PushFront(node2)
	list.PushFront(node3)
	list.PushFront(node4)
	list.PushBack(node5)

	data := list.PopFront().Data
	data1 := list.PopBack().Data

	fmt.Println(data)
	fmt.Println(data1)

	list.PrintList()
}
