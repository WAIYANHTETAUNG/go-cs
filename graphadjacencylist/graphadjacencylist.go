package graphadjacencylist

import (
	"fmt"

	"github.com/pyaesone17/go-cs/singlylinkedlist"
)

type GraphAdjacency struct {
	adjList map[int]*singlylinkedlist.LinkedList
}

func NewGraphAdjacency() *GraphAdjacency {
	return &GraphAdjacency{}
}

func (graph *GraphAdjacency) AddEdge(src int, dest int) {

	if graph.adjList == nil {
		graph.adjList = make(map[int]*singlylinkedlist.LinkedList)
	}

	_, ok := graph.adjList[src]

	if !ok {
		graph.adjList[src] = singlylinkedlist.NewLinkedList()
	}
	graph.adjList[src].PushFront(singlylinkedlist.NewNode(dest))

	_, okDest := graph.adjList[dest]
	if !okDest {
		graph.adjList[dest] = singlylinkedlist.NewLinkedList()
	}

	graph.adjList[dest].PushFront(singlylinkedlist.NewNode(src))
}

func (graph *GraphAdjacency) PrintGraph() {
	fmt.Println(graph.adjList)
	for key, row := range graph.adjList {
		node := row.Head
		if node != nil {
			fmt.Printf("%d >> ", key)
		}
		for node != nil {
			if node.Next == nil {
				fmt.Printf("%d", node.Data)
			} else {
				fmt.Printf("%d >> ", node.Data)
			}
			node = node.Next
		}
		fmt.Println()
	}
}
