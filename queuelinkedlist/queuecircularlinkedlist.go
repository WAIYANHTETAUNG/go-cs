package queuecircularlinkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	next *Node
	prev *Node
	data int
}

type Queue struct {
	Head *Node
	Tail *Node
	size int
}

func NewQueue() *Queue {
	return &Queue{}
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

func (queue *Queue) Enqueue(node *Node) {
	if queue.Head == nil {
		queue.Head = node
		queue.Tail = node
		queue.size++
		return
	}
	currentTail := queue.Tail
	queue.Tail = node
	currentTail.next = node
	queue.Tail.prev = currentTail
	queue.size++
}

func (queue *Queue) Dequeue() (*Node, error) {
	if queue.Empty() {
		fmt.Println("Queue is empty")
		return nil, errors.New("Queue is empty")
	}

	if queue.Tail.prev == nil {
		item := queue.Tail
		queue.Tail = nil
		queue.Head = nil

		queue.size--
		return item, nil
	}
	item := queue.Head
	queue.Head = queue.Head.next
	queue.Head.prev = nil
	queue.size--
	return item, nil
}

func (queue *Queue) Empty() bool {
	return queue.Head == nil
}

func (queue *Queue) Size() int {
	return queue.size
}
