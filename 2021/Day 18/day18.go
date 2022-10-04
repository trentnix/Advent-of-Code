// Package day17 implements the 2021 Advent of Code Day 17 assignment.
// See the readme.md for details on this assignment or visit the Advent
// of Code website: https://adventofcode.com/2021/day/17
package main

import (
	"fmt"
	"log"
	"os"

	"fileprocessing"
)

// main() prints the Part 1 and Part 2 solutions. It receives a single
// argument specifying the name of the data file containing the snapfish
// number input and is used to calculate the Part 1 and Part 2 solutions.
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
}
