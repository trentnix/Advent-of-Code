package graph

import (
	"math"
	"sort"
)

// the code below follows the algorithm specified for Dijkstra's Shortest Path
// which finds the shortest path between two nodes on a connected graph
// more information about Dijkstra's Shortest Path algorithm can be found at:
// https://en.wikipedia.org/wiki/Dijkstra's_algorithm

func (g *Graph) ShortestPathFromOrigin(origin int, destination int) int {
	graphSize := g.NumNodes
	if graphSize <= 1 {
		return -1
	}

	if destination < 0 || destination >= graphSize {
		return -1
	}

	visited := make([]bool, graphSize)
	weights := make([]int, graphSize)
	weights[origin] = 0
	for i := 1; i < graphSize; i++ {
		weights[i] = math.MaxInt
		visited[i] = false
	}

	g.CalculateShortestPath(origin, weights, visited)

	return weights[destination]
}

func (g *Graph) CalculateShortestPath(position int, distances []int, visited []bool) {
	var edges PairList
	edges = append(edges, Pair{position, 0})

	for len(edges) > 0 {
		position = edges[0].Key

		for _, edge := range g.Edges[position] {
			if !visited[edge.To] {
				if distances[edge.To] == math.MaxInt {
					edges = append(edges, Pair{edge.To, edge.Weight})
				}

				weight := distances[position] + edge.Weight
				if weight < distances[edge.To] {
					distances[edge.To] = weight
				}
			}
		}

		visited[position] = true
		edges = edges[1:]

		for i := 0; i < len(edges); i++ {
			edges[i].Value = distances[edges[i].Key]
		}

		sort.Sort(edges)
	}
}

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}
