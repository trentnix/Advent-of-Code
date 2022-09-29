// Package day17 implements the 2021 Advent of Code Day 17 assignment.
// See the readme.md for details on this assignment or visit the Advent
// of Code website: https://adventofcode.com/2021/day/17
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"fileprocessing"
)

const (
	Drag    int = -1
	Gravity int = -1
)

type TargetArea struct {
	Xmin, Xmax int
	Ymin, Ymax int
}

// main() prints the Part 1 and Part 2 solutions. It receives a single
// argument specifying the name of the data file containing the hexidecimal
// string input, which is converted to a binary string and then used to
// calculate the Part 1 and Part 2 solutions.
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

	ta := ParseInput(fileContents[0])
	launchers := FindAllIntersectingPaths(ta)
	maxHeight, maxHeightLauncher := FindMaximumHeight(launchers)

	fmt.Printf("Part 1 - (Max Height): %d using intial velocity (%d, %d)\n", maxHeight, maxHeightLauncher.xVelocity, maxHeightLauncher.yVelocity)
	fmt.Printf("Part 2 - Number of Paths: %d\n", len(launchers))
}

// parseInput() parses the input data into a TargetArea structure. The input
// will be in the following format:
// target area: x=20..30, y=-10..-5
func ParseInput(input string) TargetArea {
	data := input[len("target area: x="):]
	values := strings.Split(data, ", ")
	xVals := strings.Split(values[0], "..")
	yVals := strings.Split(values[1][2:], "..")

	var ta TargetArea
	ta.Xmin, _ = strconv.Atoi(xVals[0])
	ta.Xmax, _ = strconv.Atoi(xVals[1])
	ta.Ymin, _ = strconv.Atoi(yVals[0])
	ta.Ymax, _ = strconv.Atoi(yVals[1])

	return ta
}

type Coordinate struct {
	x, y int
}

type Launcher struct {
	xVelocity, yVelocity int
	path                 []Coordinate
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// FindAllIntersectingPaths() determines all the possible initial x and y velocities that
// have paths that intersect with the specified TargetArea.
func FindAllIntersectingPaths(ta TargetArea) []Launcher {
	var paths []Launcher
	var path []Coordinate

	// determine an acceptible x and y range to check
	xMin, xMax, yMin, yMax := 0, ta.Xmax, -1*absInt(ta.Ymin), absInt(ta.Ymin)
	for x := xMin; x <= xMax; x++ {
		var isHit bool

		for y := yMin; y <= yMax; y++ {
			isHit, path = CalculatePath(x, y, ta)
			if isHit {
				// this path hit the TargetArea
				paths = append(paths, Launcher{xVelocity: x, yVelocity: y, path: path})
			}
		}
	}

	return paths
}

// CalculatePath() takes a given set of x,y velocities and calculates whether or not the
// velocities will result in a hit inside the specified target area. A flag indicating the
// target was hit hit and the x,y coordinates of that position are returned.
func CalculatePath(xVelocity, yVelocity int, ta TargetArea) (bool, []Coordinate) {
	t := 0
	var xHit, yHit = false, false
	var xVal, yVal = 0, 0

	var path []Coordinate

	for {
		xHit, yHit = false, false

		if (xVelocity - t) > 0 {
			xVal = xVal + (xVelocity + (Drag * t))
		} else {
			if xVal < ta.Xmin {
				// this trajectory never gets over the specified TargetArea
				break
			}
		}

		yVal = yVal + (yVelocity + (Gravity * t))
		if yVal < ta.Ymin {
			// this trajectory is below the specified TargetArea
			break
		}

		path = append(path, Coordinate{x: xVal, y: yVal})

		if ta.Xmin <= xVal && xVal <= ta.Xmax {
			xHit = true
		}

		if ta.Ymin <= yVal && yVal <= ta.Ymax {
			yHit = true
		}

		if xHit && yHit {
			break
		}

		t++
	}

	if xHit && yHit {
		return true, path
	}

	return false, nil
}

// FindMaximumHeight() determines the highest height value for any of the specified paths
func FindMaximumHeight(launchers []Launcher) (int, Launcher) {
	var maxHeight int
	var maxHeightLauncher Launcher

	for _, launcher := range launchers {
		for _, coordinate := range launcher.path {
			if coordinate.y > maxHeight {
				maxHeight = coordinate.y
				maxHeightLauncher = launcher
			}
		}
	}

	return maxHeight, maxHeightLauncher
}
