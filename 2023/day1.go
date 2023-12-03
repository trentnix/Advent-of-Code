package main

import (
	"fileprocessing"
	"log"
	"math"
	"strconv"
	"strings"
)

func day1(name string, inputFile string) string {
	fileContents, _, err := fileprocessing.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var output strings.Builder

	sumOfCalibrationValues := day1part1(fileContents)
	strSumOfCalibrationValues := strconv.Itoa(sumOfCalibrationValues)

	output.WriteString("Advent of Code 2023 - Day 1:\n")

	output.WriteString("  Part 1: The sum of the calibration values is ")
	output.WriteString(strSumOfCalibrationValues)
	output.WriteString("\n")

	sumOfCalibrationValues = day1part2(fileContents)
	strSumOfCalibrationValues = strconv.Itoa(sumOfCalibrationValues)

	output.WriteString("  Part 2: The sum of the calibration values is ")
	output.WriteString(strSumOfCalibrationValues)
	output.WriteString("\n")

	return output.String()
}

func day1part1(lines []string) int {
	sumCalibrationValues := 0
	for _, line := range lines {
		sumCalibrationValues += getCalibrationValue(findfirstDigit(line), findLastDigit(line))
	}

	return sumCalibrationValues
}

func day1part2(lines []string) int {
	sumCalibrationValues := 0
	for _, line := range lines {
		sumCalibrationValues += getCalibrationValue(findfirstDigitWithSubstitution(line), findLastDigitWithSubstitution(line))
	}

	return sumCalibrationValues
}

func replaceStringDigits(line string) string {
	digits := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	earliestMatch := -1
	earliestIndex := math.MaxInt

	for {
		for i := 0; i < len(digits); i++ {
			digit := digits[i]
			index := strings.Index(line, digit)
			if index >= 0 {
				if index < earliestIndex {
					earliestMatch = i
					earliestIndex = index
				}
			}
		}

		if earliestMatch >= 0 {
			digit := digits[earliestMatch]
			line = line[:earliestIndex] + convertStringDigitToInt(digit) + line[earliestIndex+len(digit):]

			earliestMatch = -1
			earliestIndex = math.MaxInt
		} else {
			break
		}
	}

	return line
}

var defaultRune rune = 'x'

func findfirstDigit(line string) rune {
	runes := []rune(line)
	for i := 0; i < len(runes); i++ {
		if isDigit(runes[i]) {
			return runes[i]
		}
	}

	// default value
	return defaultRune
}

func findLastDigit(line string) rune {
	runes := []rune(line)
	for i := len(runes); i >= 0; i-- {
		if isDigit(runes[i]) {
			return runes[i]
		}
	}

	// default value
	return defaultRune
}

func findfirstDigitWithSubstitution(line string) rune {
	digits := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	runes := []rune(line)
	for i := 0; i < len(runes); i++ {
		if isDigit(runes[i]) {
			return runes[i]
		}

		for _, digit := range digits {
			found := strings.Index(line[i:], digit)
			if found == 0 {
				return convertStringDigitToRune(digit)
			}
		}
	}

	// default value
	return defaultRune
}

func findLastDigitWithSubstitution(line string) rune {
	runes := []rune(line)
	for i := len(runes); i >= 0; i-- {
		if isDigit(runes[i]) {
			return runes[i]
		}
	}

	// default value
	return defaultRune
}

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

func convertStringDigitToInt(digit string) string {
	switch digit {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		log.Fatal("Could not convert [", digit, "] to its numerical value")
	}

	return ""
}

func getCalibrationValue(first rune, last rune) int {
	calibrationVal, err := strconv.Atoi(string(first) + string(last))
	if err != nil {
		log.Fatal("There was an error converting ", string(first)+string(last), " to its calibration value: ", err)
	}

	return calibrationVal
}

func getCalibrationValueOld(line string) int {
	var first, last rune

	first = '!'
	for _, val := range line {
		if isDigit(val) {
			if first == '!' {
				first = val
			}

			last = val
		}
	}

	calibrationVal, err := strconv.Atoi(string(first) + string(last))
	if err != nil {
		log.Fatal("There was an error converting ", line, " to its calibration value: ", err)
	}

	return calibrationVal
}

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
