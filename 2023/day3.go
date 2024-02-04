// this file implements the 2023 Advent of Code Day 2 assignment.
// See the readme.md for details on this assignment or visit the Advent
// of Code website: https://adventofcode.com/2023/day/2

package main

import (
	"fileprocessing"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// day3() has the Elf and I reaching a gondola lift station which will take us up to
// the water source, but the gondolas aren't moving and we need to use the engine
// schematic to fix it.
func day3(name string, inputFile string) string {
	fileContents, _, err := fileprocessing.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	e, partNumbers := processInput(fileContents)

	sumOfPartNumbers := strconv.Itoa(day3part1(e, partNumbers))

	var output strings.Builder
	output.WriteString("Part 1:\n")
	output.WriteString("The sum of all valid part numbers is: " + sumOfPartNumbers)
	output.WriteString("\n")

	sumOfGears := strconv.Itoa(day3part2(e, partNumbers))

	output.WriteString("Part 2:\n")
	output.WriteString("The sum of all gear ratios is: " + sumOfGears)

	return output.String()
}

// day3part1() calculates the sum of partNumbers by looking for partNumbers that
// have an adjacent symbol (which confirms it is a part number)
func day3part1(e engineSchematic, partNumbers []partNumber) int {
	sum := 0

	for _, partNumber := range partNumbers {
		if e.partNumberHasAdjacentSymbol(partNumber) {
			sum += partNumber.value
		}
	}

	return sum
}

const gear_symbol = '*'

func day3part2(e engineSchematic, partNumbers []partNumber) int {
	sum := 0

	// get each '*'
	for row, r := range e {
		for col, _ := range r {
			if e[row][col] == gear_symbol {
				firstGear := 0
				secondGear := 0

				// 	parse the partNumbers and find out if there are two gears that are adjacent
				for _, partNumber := range partNumbers {
					if row < partNumber.row-1 {
						continue
					}

					if row > partNumber.row+1 {
						continue
					}

					if partNumber.isAdjacentTo(row, col) {
						if firstGear == 0 {
							firstGear = partNumber.value
						} else {
							if secondGear == 0 {
								secondGear = partNumber.value
							} else {
								log.Fatal("found more than two numbers adjacent to the gear at " + strconv.Itoa(row) + ", " + strconv.Itoa(col))
							}
						}
					}
				}

				// 	if so, multiply and add to the sum (if no, the product will be 0)
				sum += firstGear * secondGear
			}
		}
	}

	return sum
}

type engineSchematic [][]rune
type partNumber struct {
	value           int
	row, start, end int
}

// print() prints a given engine schematic
func (e engineSchematic) print() {
	fmt.Printf("engine schematic (input):")
	fmt.Println()

	for _, row := range e {
		for _, col := range row {
			fmt.Printf("%c", col)
		}
		fmt.Println()
	}
}

// getDimensions() returns the number of rows and columns in a given engine schematic
func (e engineSchematic) getDimensions() (int, int) {
	numRows := len(e)
	numCols := 0
	if numRows > 0 {
		numCols = len(e[0])
	}

	return numRows, numCols
}

// partNumberHasAdjacentSymbol() inspects a partNumber to determine if, on the
// engineering schematic it is contained, it has an adjacent symbol
func (e engineSchematic) partNumberHasAdjacentSymbol(p partNumber) bool {
	numRows, numCols := e.getDimensions()
	startIndex := p.start
	if startIndex > 0 {
		// check the column before where the partNumber starts
		startIndex = startIndex - 1
	}

	endIndex := p.end
	if endIndex == numCols {
		// don't go past the last column
		endIndex = endIndex - 1
	}

	// check above
	if p.row > 0 {
		for i := startIndex; i <= endIndex; i++ {
			if !runeIsNumberOrDot(e[p.row-1][i]) {
				return true
			}
		}
	}

	// check current row
	if startIndex < p.start {
		// check left, if there is a left
		if !runeIsNumberOrDot(e[p.row][startIndex]) {
			return true
		}
	}

	if endIndex == p.end {
		// check right, if there is a right
		if !runeIsNumberOrDot(e[p.row][endIndex]) {
			return true
		}
	}

	// check below
	if p.row < numRows-1 {
		for i := startIndex; i <= endIndex; i++ {
			if !runeIsNumberOrDot(e[p.row+1][i]) {
				return true
			}
		}
	}

	return false
}

// isAdjacentTo() the (x,y) coordinate provided is adjacent to the location of
// the partNumber in the grid
func (p partNumber) isAdjacentTo(row int, col int) bool {
	if row >= p.row-1 && row <= p.row+1 {
		if col >= p.start-1 && col <= p.end {
			return true
		}
	}

	return false
}

// processInput() takes an array of strings and parses it into an engineSchematic and
// an array of partNumbers
func processInput(input []string) (engineSchematic, []partNumber) {
	var e engineSchematic
	var partNumbers []partNumber
	for row, s := range input {
		e = append(e, []rune(s))

		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllStringIndex(s, -1)
		for _, match := range matches {
			var p partNumber
			p.start = match[0]
			p.end = match[1]
			value, err := strconv.Atoi(s[p.start:p.end])
			if err != nil {
				log.Fatal(err)
			}

			p.value = value
			p.row = row

			partNumbers = append(partNumbers, p)
		}
	}

	return e, partNumbers
}

// runeIsNumberOrDot() determines whether the specified rune is a digit or '.' value
func runeIsNumberOrDot(r rune) bool {
	return (r >= '0' && r <= '9') || r == '.'
}
