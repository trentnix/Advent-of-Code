// Package day19 implements the 2021 Advent of Code Day 19 assignment.
// See the readme.md for details on this assignment or visit the Advent
// of Code website: https://adventofcode.com/2021/day/19
package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"fileprocessing"
)

type Position struct {
	x, y, z int
}

type Scanner struct {
	ID                 int
	Beacons            []Position
	BeaconPermutations [][]Position
}

func (s *Scanner) New(id int, input []string) error {
	if len(input) == 0 {
		return fmt.Errorf("could not parse a new Scanner - input was empty")
	}

	s.ID = id

	for _, line := range input {
		if strings.Contains(line, "scanner") {
			// this is the first line
			parsedHeader := strings.Split(line, " ")
			scannerID, err := strconv.Atoi(parsedHeader[2])
			if err != nil {
				return err
			}

			s.ID = scannerID
		} else {
			coordinates := strings.Split(line, ",")
			x, errX := strconv.Atoi(coordinates[0])
			if errX != nil {
				return errX
			}
			y, errY := strconv.Atoi(coordinates[1])
			if errY != nil {
				return errY
			}
			z, errZ := strconv.Atoi(coordinates[2])
			if errZ != nil {
				return errZ
			}

			s.Beacons = append(s.Beacons, Position{x, y, z})
		}
	}

	orientations := getPositionPermutations()
	s.BeaconPermutations = make([][]Position, len(orientations))

	for i, orientation := range orientations {
		s.BeaconPermutations[i] = make([]Position, len(s.Beacons))
		for j, p := range s.Beacons {
			s.BeaconPermutations[i][j] = rotate_XYZ(p, orientation)
		}
	}

	return nil
}

// rotate_X
// Rotate a vector around the X axis
// x' = x
// y' = y * cos(angle) - z * sin(angle)
// z' = y * sin(angle) + z * cos(angle)
func rotate_X(beacon, rotation Position) Position {
	angle := float64(rotation.x)

	x := beacon.x
	y := (beacon.y * int(math.Cos(angle))) - (beacon.z * int(math.Sin(angle)))
	z := (beacon.y * int(math.Sin(angle))) + (beacon.z * int(math.Cos(angle)))

	return Position{x, y, z}
}

// rotate_Y
// Rotate a vector around the X axis
// x' = x * cos(angle) + z * sin(angle)
// y' = y
// z' = y * cos(angle) + x * sin(angle)
func rotate_Y(beacon, rotation Position) Position {
	angle := float64(rotation.y)

	x := (beacon.x * int(math.Cos(angle))) + (beacon.z * int(math.Sin(angle)))
	y := beacon.y
	z := (beacon.y * int(math.Cos(angle))) - (beacon.z * int(math.Sin(angle)))

	return Position{x, y, z}
}

// rotate_Z
// Rotate a vector around the X axis
// x' = x * cos(angle) - y * sin(angle)
// y' = x * sin(angle) + y * cos(angle)
// z' = z
func rotate_Z(beacon, rotation Position) Position {
	angle := float64(rotation.z)

	x := (beacon.x * int(math.Cos(angle))) + (beacon.z * int(math.Sin(angle)))
	y := (beacon.y * int(math.Cos(angle))) - (beacon.z * int(math.Sin(angle)))
	z := beacon.z

	return Position{x, y, z}
}

func rotate_XYZ(a, rotation Position) Position {
	a = rotate_X(a, rotation)
	a = rotate_Y(a, rotation)

	return rotate_Z(a, rotation)
}

func getPositionPermutations() [24]Position {
	return [24]Position{
		Position{0, 0, 0},
		Position{90, 0, 0},
		Position{180, 0, 0},
		Position{270, 0, 0},
		Position{0, 90, 0},
		Position{90, 90, 0},
		Position{180, 90, 0},
		Position{270, 90, 0},
		Position{0, 180, 0},
		Position{90, 180, 0},
		Position{180, 180, 0},
		Position{270, 180, 0},
		Position{0, 270, 0},
		Position{90, 270, 0},
		Position{180, 270, 0},
		Position{270, 270, 0},
		Position{0, 0, 90},
		Position{90, 0, 90},
		Position{180, 0, 90},
		Position{270, 0, 90},
		Position{0, 0, 270},
		Position{90, 0, 270},
		Position{180, 0, 270},
		Position{270, 0, 270},
	}
}

// main() prints the Part 1 and Part 2 solutions. It receives a single
// argument specifying the name of the data file containing the ...
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

	scanners := parseInput(fileContents)
	for i, scanner := range scanners {
		fmt.Printf("Scanner %d\nNumber of Beacons: %d, Number of Permutations: %d\n%v\n", i, len(scanner.Beacons), len(scanner.BeaconPermutations)*len(scanner.BeaconPermutations[0]), scanner)
	}
}

func parseInput(input []string) []Scanner {
	var scanners []Scanner

	id, start := 0, 0
	for i, line := range input {
		if i == len(input)-1 {
			// this is the last line of the input
			var scanner *Scanner = new(Scanner)
			scanner.New(id, input[start:i+1])
			scanners = append(scanners, *scanner)
		}

		if len(line) == 0 {
			// this is the last line of a section
			var scanner *Scanner = new(Scanner)
			scanner.New(id, input[start:i])
			scanners = append(scanners, *scanner)

			id++
		}

		if strings.Contains(line, "scanner") {
			start = i
		}
	}

	return scanners
}
