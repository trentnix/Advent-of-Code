// Package day18 implements the 2021 Advent of Code Day 18 assignment.
// See the readme.md for details on this assignment or visit the Advent
// of Code website: https://adventofcode.com/2021/day/18
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

	snailfishNumber := ""
	for _, input := range fileContents {
		// first reduce the input string (just in case)
		input = reduceSnailfishNumber(input)
		// add the input to the existing snailfish number
		snailfishNumber = addSnailfishNumbers(snailfishNumber, input)
		// reduce the result
		snailfishNumber = reduceSnailfishNumber(snailfishNumber)
	}

	magnitude := calculateMagnitude(snailfishNumber)
	fmt.Printf("Part 1: %s = %d\n", snailfishNumber, magnitude)

	largestMagnitude := findLargestMagnitudeOfTwoNumbers(fileContents)
	fmt.Printf("Part 2: Largest magnitude of 2 input numbers: %d\n", largestMagnitude)
}

// addSnailfishNumbers() adds a new snailfish number to an existing
// snailfish number. To add two snailfish numbers, form a pair from
// the left and right parameters of the addition operator. For
// example, [1,2] + [[3,4],5] becomes [[1,2],[[3,4],5]].
func addSnailfishNumbers(left string, right string) string {
	if len(left) == 0 {
		return right
	}

	return "[" + left + "," + right + "]"
}

// reduceSnailfishNumber() takes a string and reduces it by using the
// 'explode' and 'split' functionality to reduce it to a resulting
// snailfish number.
//
// A key note - you should traverse the entire list to 'explode' all possibilities
// before doing a 'split'. And if a 'split' occurs, traverse the list to see if any
// new 'explode' opportunities have been presented.
func reduceSnailfishNumber(s string) string {
	depth := 0
	bExplode, bSplit := true, true

	for bExplode || bSplit {
		bExplode, bSplit = false, false
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '[':
				depth++
				continue
			case ']':
				depth--
				continue
			}

			if depth >= 5 {
				// explode
				bExplode = true
				s = explode(i, s)
				depth = 0
				i = -1
			}
		}

		for i := 0; i < len(s); i++ {
			if isDigit(s[i]) {
				if isDigit(s[i+1]) {
					bSplit = true
					s = split(i, s)
					break
				}
			}
		}
	}

	return s
}

// explode() takes the pair that starts at index start of string s
// and does an 'explode', where the pair's left value is added to the
// first regular number to the left of the exploding pair (if any),
// and the pair's right value is added to the first regular number to
// the right of the exploding pair (if any).
//
// Then, the entire exploding pair is replaced with the regular number 0.
func explode(start int, s string) (output string) {
	output = s

	var left, right string
	var end int
	for i := start; i < len(s); i++ {
		if s[i] == ']' {
			// found the closing bracket for the pair
			end = i
			// get the left and right side of the pair
			numbers := strings.Split(s[start:end], ",")
			if len(numbers) >= 2 {
				left = numbers[0]
				right = numbers[1]
			}

			break
		}
	}

	leftVal, _ := strconv.Atoi(left)
	rightVal, _ := strconv.Atoi(right)

	// add first number to previous number
	var prevNumber string
	var prevNumberVal, prevStart, prevEnd int
	for i := start - 1; i > 0; i-- {
		if isDigit(s[i]) {
			prevEnd = i + 1

			// found a number, now get all the digits
			for y := i - 1; y > 0; y-- {
				if !isDigit(s[y]) {
					prevStart = y + 1

					prevNumber = s[prevStart:prevEnd]
					prevNumberVal, _ = strconv.Atoi(prevNumber)

					break
				}
			}
		}

		if len(prevNumber) > 0 {
			break
		}
	}

	// add second number to next number
	var nextNumber string
	var nextNumberVal, nextStart, nextEnd int
	for i := end; i < len(s); i++ {
		if isDigit(s[i]) {
			nextStart = i

			// found a number, now get all the digits
			for y := i + 1; y < len(s); y++ {
				if !isDigit(s[y]) {
					nextEnd = y

					nextNumber = s[nextStart:nextEnd]
					nextNumberVal, _ = strconv.Atoi(nextNumber)

					break
				}
			}
		}

		if len(nextNumber) > 0 {
			break
		}
	}

	if len(nextNumber) > 0 && len(prevNumber) > 0 {
		// there is a previous and next number
		output = s[:prevStart] + strconv.Itoa(prevNumberVal+leftVal) + s[prevEnd:start-1] + "0" + s[end+1:nextStart] + strconv.Itoa(nextNumberVal+rightVal) + s[nextEnd:]
	}

	if len(nextNumber) > 0 && len(prevNumber) <= 0 {
		// there is no previous number
		output = s[:start-1] + "0" + s[end+1:nextStart] + strconv.Itoa(nextNumberVal+rightVal) + s[nextEnd:]
	}

	if len(nextNumber) <= 0 && len(prevNumber) > 0 {
		// there is no next number
		output = s[:prevStart] + strconv.Itoa(prevNumberVal+leftVal) + s[prevEnd:start-1] + "0" + s[end+1:]
	}

	// replace pair with 0

	return
}

// split() takes the value specified at s[start:start+2] and splits it into
// a snailfish value pair
func split(start int, s string) (output string) {
	// this number is 10 or larger, need to split
	num := s[start : start+2]
	numVal, _ := strconv.Atoi(num)

	leftVal := int(numVal / 2)
	rightVal := int(numVal/2) + (numVal % 2)

	output = s[:start] + "[" + strconv.Itoa(leftVal) + "," + strconv.Itoa(rightVal) + "]" + s[start+2:]

	return
}

// isDigit() determines whether the byte specified is a digit or not
func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// calculateMagnitude() calculates the magnitude of the input snailfish
// number. The magnitude of a pair is 3 times the magnitude of its left
// element plus 2 times the magnitude of its right element. The
// magnitude of a regular number is just that number.
func calculateMagnitude(s string) int {
	continueProcessing := true
	for continueProcessing {
		continueProcessing = false

		// find each comma and calculate the magnitude of each snailfish number until the string is
		// reduced to a single value
		for i := 0; i < len(s); i++ {
			if s[i] == ',' {
				if s[i+1] == '[' {
					continue
				}

				// we found a pair (since we found a comma), make sure we
				// try to run through the string again
				continueProcessing = true
				commaIndex := i

				// get end
				endBracketIndex := commaIndex + 2
				for {
					if s[endBracketIndex] == ']' {
						break
					}

					endBracketIndex++
				}

				// get beginning
				startBracketIndex := commaIndex - 1
				for {
					if s[startBracketIndex] == '[' {
						break
					}

					startBracketIndex--
				}

				left := s[startBracketIndex+1 : commaIndex]
				right := s[commaIndex+1 : endBracketIndex]

				leftVal, _ := strconv.Atoi(left)
				rightVal, _ := strconv.Atoi(right)

				magnitude := leftVal*3 + rightVal*2
				if endBracketIndex+1 < len(s) {
					s = s[:startBracketIndex] + strconv.Itoa(magnitude) + s[endBracketIndex+1:]
					i = 0
					continue
				} else {
					// this is the last substitution
					return magnitude
				}
			}
		}
	}

	return 0
}

// findLargestMagnitudeOfTwoNumbers() takes the input and tries all
// combinations of two snailfish numbers in the input to figure out
// the largest magnitude of the possible combinations.
func findLargestMagnitudeOfTwoNumbers(input []string) int {
	largestValue := 0
	for x, left := range input {
		for y, right := range input {
			if x != y {
				// first reduce the input string (just in case)
				snailfishNumber := reduceSnailfishNumber(left)
				// add the input to the existing snailfish number
				snailfishNumber = addSnailfishNumbers(snailfishNumber, right)
				// reduce the result
				snailfishNumber = reduceSnailfishNumber(snailfishNumber)

				magnitude := calculateMagnitude(snailfishNumber)
				if magnitude > largestValue {
					largestValue = magnitude
				}
			}
		}
	}

	return largestValue
}
