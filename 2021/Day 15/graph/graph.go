package graph

import (
	"fmt"
)

// this models a directed graph with weights on the edges.
type Graph struct {
	NumNodes int
	Edges    [][]Edge
}

type Edge struct {
	From   int
	To     int
	Weight int
}

func NewGraph(n int) *Graph {
	return &Graph{
		NumNodes: n,
		Edges:    make([][]Edge, n),
	}
}

func (g *Graph) AddEdge(from int, to int, weight int) {
	g.Edges[from] = append(g.Edges[from], Edge{From: from, To: to, Weight: weight})
}

// print a list of the nodes and the weights of the edges to the nodes they point to
func (g *Graph) PrintAdjacentEdges() {
	fmt.Println("Printing all edges in the graph.")
	for _, adjacent := range g.Edges {
		for _, e := range adjacent {
			fmt.Printf("Edge: %d -> %d (%d)\n", e.From, e.To, e.Weight)
		}
	}
}

// print the specified graph as a matrix of values
// you need to pass the initial value (should be 0), the number of rows to print,
// and the number of columns to print
func (g *Graph) PrintGraph(rowLength int, initialValue int) {
	values := make([]int, g.NumNodes)

	values[0] = initialValue
	for _, adjacent := range g.Edges {
		for _, e := range adjacent {
			values[e.To] = e.Weight
		}
	}

	numRows := g.NumNodes / rowLength
	if g.NumNodes%rowLength > 0 {
		numRows += 1
	}

	fmt.Printf("\n")
	for y := 0; y < numRows; y++ {
		for x := 0; x < rowLength; x++ {
			if y*rowLength+x < g.NumNodes {
				fmt.Printf("%d", values[y*rowLength+x])
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Printf("\n")
	}
}
