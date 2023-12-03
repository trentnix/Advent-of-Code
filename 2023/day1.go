// this file implements the 2023 Advent of Code Day 1 assignment.
// See the readme.md for details on this assignment or visit the Advent
// of Code website: https://adventofcode.com/2023/day/1

package main

import (
	"fileprocessing"
	"log"
	"strconv"
	"strings"
)

// default value to return when parsing a line goes bad
const defaultRune rune = 'X'

// day1() wraps the two parts of the solution and returns a string (output) with the
// results of the two solves.
func day1(name string, inputFile string) string {
	fileContents, _, err := fileprocessing.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var output strings.Builder

	// part 1

	sumOfCalibrationValues := 0
	for _, line := range fileContents {
		sumOfCalibrationValues += getCalibrationValue(findFirstDigit(line), findLastDigit(line))
	}

	strSumOfCalibrationValues := strconv.Itoa(sumOfCalibrationValues)

	output.WriteString("Advent of Code 2023 - Day 1:\n")

	output.WriteString("  Part 1: The sum of the calibration values is ")
	output.WriteString(strSumOfCalibrationValues)
	output.WriteString("\n")

	// part 2

	sumOfCalibrationValues = 0
	for _, line := range fileContents {
		sumOfCalibrationValues += getCalibrationValue(findFirstDigitWithSubstitution(line), findLastDigitWithSubstitution(line))
	}

	strSumOfCalibrationValues = strconv.Itoa(sumOfCalibrationValues)

	output.WriteString("  Part 2: The sum of the calibration values is ")
	output.WriteString(strSumOfCalibrationValues)
	output.WriteString("\n")

	return output.String()
}

// getCalibrationValue(first rune, last rune) returns the integer value that results when
// appending the 'first' and 'last' digit provided.
//
//	ex: '1' and '2' = '12' = 12
//	ex: '5' and '5' = '55' = 55
//	ex: '4' and '9' = '49' = 49
func getCalibrationValue(first rune, last rune) int {
	if first == defaultRune || last == defaultRune {
		log.Fatal("The specified digits are invalid. first: " + string(first) + ", last: " + string(last))
	}

	calibrationVal, err := strconv.Atoi(string(first) + string(last))
	if err != nil {
		log.Fatal("There was an error converting ", string(first)+string(last), " to its calibration value: ", err)
	}

	return calibrationVal
}

// part 1 looks just for digits. The functions below start from the beginning and end,
// respectively looking for a digit from which to build the calibration value

// findFirstDigit(line string) returns the first digit encountered when parsing the input 'line'
// from left to right
func findFirstDigit(line string) rune {
	runes := []rune(line)
	for i := 0; i < len(runes); i++ {
		if isDigit(runes[i]) {
			return runes[i]
		}
	}

	return defaultRune
}

// findLastDigit(line string) returns the first digit encountered when parsing the input 'line'
// from right to left
func findLastDigit(line string) rune {
	runes := []rune(line)
	for i := len(runes) - 1; i >= 0; i-- {
		if isDigit(runes[i]) {
			return runes[i]
		}
	}

	return defaultRune
}

// part 2 looks just for either the digit or a numeral that is spelled out. The functions below
// start from the beginning and end, respectively looking for a digit, whether numeric or spelled
// out, from which to build the calibration value

// I made the mistake of navigating each line and replacing the spelled out version with a number
// and then finding the first and last numeric digit. If you have a line like '1twone', starting
// from the beginning gives you '12ne'. But if you start from the end, you get '1tw1'. So you
// have to start from the end when finding the second digit in a calibration value.

// findFirstDigitWithSubstitution(line string) returns the first digit encountered when parsing
// the input 'line' from left to right whether it is spelled out or provided numerically
func findFirstDigitWithSubstitution(line string) rune {
	digits := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	runes := []rune(line)
	for i := 0; i < len(runes); i++ {
		if isDigit(runes[i]) {
			return runes[i]
		}

		// this might shave a few cycles if the string 'line[i:]' is shorter than three. Any
		// check of the spelled-out digits would be impossible in that circumstance.
		//   - you could get cute and only check the digits when 'len(line[i:]) < len(digit)'
		//     (or similar logic) but the performance benefits would depend on the input

		for _, digit := range digits {
			found := strings.Index(line[i:], digit)
			if found == 0 {
				return convertStringDigitToRune(digit)
			}
		}
	}

	return defaultRune
}

// findLastDigitWithSubstitution(line string) returns the first digit encountered when parsing
// the input 'line' from right to left whether it is spelled out or provided numerically
func findLastDigitWithSubstitution(line string) rune {
	digits := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	runes := []rune(line)
	for i := len(runes) - 1; i >= 0; i-- {
		if isDigit(runes[i]) {
			return runes[i]
		}

		// this might shave a few cycles if the string 'line[i:]' is shorter than three. Any
		// check of the spelled-out digits would be impossible in that circumstance.
		//   - you could get cute and only check the digits when 'len(line[i:]) < len(digit)'
		//     (or similar logic) but the performance benefits would depend on the input

		for _, digit := range digits {
			found := strings.Index(line[i:], digit)
			if found == 0 {
				return convertStringDigitToRune(digit)
			}
		}
	}

	return defaultRune
}

// convertStringDigitToRune(digit string) converts a string spelled as a number to its rune
// (character) equivalent
func convertStringDigitToRune(digit string) rune {
	switch digit {
	case "one":
		return '1'
	case "two":
		return '2'
	case "three":
		return '3'
	case "four":
		return '4'
	case "five":
		return '5'
	case "six":
		return '6'
	case "seven":
		return '7'
	case "eight":
		return '8'
	case "nine":
		return '9'
	default:
		log.Fatal("Could not convert [", digit, "] to its numerical value")
	}

	return defaultRune
}

// isDigit(c rune) checks whether the rune value provided is a digit or not
func isDigit(c rune) bool {
	switch c {
	case '1':
		fallthrough
	case '2':
		fallthrough
	case '3':
		fallthrough
	case '4':
		fallthrough
	case '5':
		fallthrough
	case '6':
		fallthrough
	case '7':
		fallthrough
	case '8':
		fallthrough
	case '9':
		return true

	}

	return false
}
