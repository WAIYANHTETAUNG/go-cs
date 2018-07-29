package graphadjacencylist

import (
	"testing"
)

func TestHash(t *testing.T) {
	graph := NewGraphAdjacency()
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.PrintGraph()
}
