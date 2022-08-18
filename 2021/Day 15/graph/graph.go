package graph

import (
	"fmt"
)

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

func (g *Graph) PrintAdjacentEdges() {
	fmt.Println("Printing all edges in the graph.")
	for _, adjacent := range g.Edges {
		for _, e := range adjacent {
			fmt.Printf("Edge: %d -> %d (%d)\n", e.From, e.To, e.Weight)
		}
	}
}

func (g *Graph) PrintGraph(numRows int, numColumns int, initialValue int) {
	values := make([]int, g.NumNodes)

	values[0] = initialValue
	for _, adjacent := range g.Edges {
		for _, e := range adjacent {
			values[e.To] = e.Weight
		}
	}

	fmt.Printf("\n")
	for y := 0; y < numRows; y++ {
		for x := 0; x < numColumns; x++ {
			fmt.Printf("%d", values[y*numColumns+x])
		}
		fmt.Printf("\n")
	}
}
