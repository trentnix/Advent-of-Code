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

	actualSum := day3part1(processInput(input))
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
	actualSum := day3part2(processInput(input))
	if expectedSum != actualSum {
		t.Errorf("expected: " + strconv.Itoa(expectedSum) + ", actual: " + strconv.Itoa(actualSum))
	}
}
