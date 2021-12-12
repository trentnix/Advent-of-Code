package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"sciencerocketry.com/fileprocessing"
	"sciencerocketry.com/graph"
)

func main() {
	inputFile := os.Args[1]

	fileContents, err := fileprocessing.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	numLines := len(fileContents)
	if numLines <= 0 {
		// invalid input
		log.Fatal(fmt.Errorf("invalid input in %s", inputFile))
	}

	var g graph.ItemGraph

	for _, line := range fileContents {
		path := strings.Split(line, "-")
		start := path[0]
		destination := path[1]

		g.AddEdge(start, destination)
	}

	fmt.Printf("Nodes: %d\n", g.NodeCount())

	var visited []string
	var path []string

	TraverseGraph(g, "start", path, visited, false)
	fmt.Printf("Part One - Number Distinct Paths: %d\n", len(traversals))

	traversals, visited, path = nil, nil, nil
	TraverseGraph(g, "start", path, visited, true)
	fmt.Printf("Part Two - Number Distinct Paths: %d\n", len(traversals))
}

var traversals [][]string

func TraverseGraph(g graph.ItemGraph, currentNode string, currentPath []string, visited []string, allowAnotherSmallVisit bool) {
	if currentNode == "start" && len(currentPath) > 0 {
		return
	}

	if currentNode == "end" {
		currentPath = append(currentPath, currentNode)
		traversals = append(traversals, currentPath)
		return
	}

	if elementExists(visited, currentNode) {
		if !allowAnotherSmallVisit {
			return
		}

		allowAnotherSmallVisit = false
	}

	currentPath = append(currentPath, currentNode)

	if !isUpper(currentNode) {
		visited = append(visited, currentNode)
	}

	for _, destination := range g.GetDestinations(currentNode) {
		TraverseGraph(g, destination, currentPath, visited, allowAnotherSmallVisit)
	}
}

func elementExists(a []string, s string) bool {
	for _, current := range a {
		if s == current {
			return true
		}
	}

	return false
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
