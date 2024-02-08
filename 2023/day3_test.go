package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestSumOfValidPartNumbers(t *testing.T) {
	inputText := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	input := strings.Split(inputText, "\n")

	expectedSum := 4361

	e := new(engineSchematic)
	e.new(input)
	actualSum := day3part1(e)
	if expectedSum != actualSum {
		t.Errorf("expected: " + strconv.Itoa(expectedSum) + ", actual: " + strconv.Itoa(actualSum))
	}
}

func TestSumOfGearRatios(t *testing.T) {
	inputText := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	input := strings.Split(inputText, "\n")

	expectedSum := 467835

	e := new(engineSchematic)
	e.new(input)
	actualSum := day3part2(e)
	if expectedSum != actualSum {
		t.Errorf("expected: " + strconv.Itoa(expectedSum) + ", actual: " + strconv.Itoa(actualSum))
	}
}
