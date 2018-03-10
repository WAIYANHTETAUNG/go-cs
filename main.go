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

	list.PushFront(node1)
	list.PushFront(node2)
	list.PushFront(node3)
	list.PushFront(node4)

	data := list.PopFront().(*linkedlist.Node).Data

	fmt.Println(data)
	list.PrintList()
}
