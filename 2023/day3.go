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

	sumOfPartNumbers := strconv.Itoa(day3part1(fileContents))

	var output strings.Builder
	output.WriteString("Part 1:")
	output.WriteString("The sum of all valid part numbers is: " + sumOfPartNumbers)

	return output.String()
}

// day3part1() calculates the sum of partNumbers by looking for partNumbers that
// have an adjacent symbol (which confirms it is a part number)
func day3part1(input []string) int {
	sum := 0
	e, partNumbers := processInput(input)
	for _, partNumber := range partNumbers {
		if e.partNumberHasAdjacentSymbol(partNumber) {
			sum += partNumber.value
		} else {
			// fmt.Println(strconv.Itoa(partNumber.value) + " is not a part number")
			// fmt.Println("start index: " + strconv.Itoa(partNumber.start) + " end index: " + strconv.Itoa(partNumber.end))
			// fmt.Println("row: " + strconv.Itoa(partNumber.row))
			// fmt.Println()
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
