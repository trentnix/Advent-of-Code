package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const numPatterns = 10
const numOutput = 4

type Entry struct {
	patterns                                                              []string
	output                                                                []string
	top, middle, bottom, upper_left, upper_right, lower_left, lower_right string
}

func NewEntry(input string) *Entry {
	if len(input) <= 0 {
		// insufficient input
		return nil
	}

	entry := new(Entry)

	parts := strings.Split(input, "|")
	entry.patterns = strings.Split(strings.TrimSpace(parts[0]), " ")
	entry.output = strings.Split(strings.TrimSpace(parts[1]), " ")
	entry.processPatterns()

	return entry
}

func (e *Entry) processPatterns() {
	var one, four, seven, eight string
	var len6, len5 []string

	for _, pattern := range e.patterns {
		lenPattern := len(pattern)
		switch lenPattern {
		case 2:
			// the digit is a 1
			one = pattern
		case 3:
			// the digit is a 7
			seven = pattern
		case 4:
			// the digit is a 4
			four = pattern
		case 7:
			// the digit is an 8
			eight = pattern
		case 6:
			len6 = append(len6, pattern)
		case 5:
			len5 = append(len5, pattern)
		}
	}

	for _, pattern := range len6 {
		difference := subtractString(one, pattern)
		if len(difference) == 1 {
			// this is a six and the difference is the upper-right segment
			e.upper_right = difference
			e.lower_right = subtractString(one, e.upper_right)
			continue
		}

		difference = subtractString(four, pattern)
		if len(difference) == 1 {
			// this is a zero and the difference is the middle segment
			e.middle = difference
			continue
		}

		// this is a 9 and the difference is the lower-left segment
		difference = subtractString(eight, pattern)
		e.lower_left = difference
	}

	e.top = subtractString(seven, one)
	e.upper_left = subtractString(four, one+e.middle)
	e.bottom = subtractString(eight, e.top+e.middle+e.upper_left+e.upper_right+e.lower_left+e.lower_right)

	return
}

func subtractString(a, b string) (diff string) {
	for _, a_char := range a {
		if !strings.Contains(b, string(a_char)) {
			diff += string(a_char)
		}
	}
	return
}

func (e *Entry) evaluate(input string) int {
	switch len(input) {
	case 2:
		// the digit is a 1
		return 1
	case 3:
		// the digit is a 7
		return 7
	case 4:
		// the digit is a 4
		return 4
	case 7:
		// the digit is an 8
		return 8
	case 5:
		if strings.Contains(input, e.lower_left) {
			// the digit is a 2
			return 2
		}

		if strings.Contains(input, e.upper_left) {
			// the digit is a 5
			return 5
		}

		return 3
	case 6:
		if subtractString("abcdefg", input) == e.middle {
			return 0
		}
		if subtractString("abcdefg", input) == e.upper_right {
			return 6
		}

		return 9
	}

	return -1
}

func (e *Entry) OutputValue() string {
	var outputVal string
	for _, digit := range e.output {
		outputVal += strconv.Itoa(e.evaluate(digit))
	}
	return outputVal
}

func (e *Entry) Print(w io.Writer) {
	fmt.Fprintf(w, "patterns: %v - output: %v\n", e.patterns, e.output)
	fmt.Fprintf(w, " %s%s%s%s \n", e.top, e.top, e.top, e.top)
	fmt.Fprintf(w, "%s    %s\n", e.upper_left, e.upper_right)
	fmt.Fprintf(w, "%s    %s\n", e.upper_left, e.upper_right)
	fmt.Fprintf(w, " %s%s%s%s \n", e.middle, e.middle, e.middle, e.middle)
	fmt.Fprintf(w, "%s    %s\n", e.lower_left, e.lower_right)
	fmt.Fprintf(w, "%s    %s\n", e.lower_left, e.lower_right)
	fmt.Fprintf(w, " %s%s%s%s \n", e.bottom, e.bottom, e.bottom, e.bottom)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "output value: %s\n", e.OutputValue())
	fmt.Fprintf(w, "\n")
}

func main() {
	var inputFile string
	inputFile = os.Args[1]

	fileContents, err := ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	numLines := len(fileContents)
	if numLines <= 0 {
		// invalid input
		log.Fatal(fmt.Errorf("invalid input in %s\n", inputFile))
	}

	fmt.Printf("num lines: %d\n", numLines)

	var entries []*Entry

	// import entry data
	for i := 0; i < numLines; i++ {
		entries = append(entries, NewEntry(fileContents[i]))
	}

	numUniqueSegments := countUniqueSegments(entries)
	fmt.Printf("Part One - digits 1, 4, 7, or 8 appear %d times\n", numUniqueSegments)

	sumOutputs := sumOutputValues(entries)
	fmt.Printf("Part Two - sum of the output values: %d\n", sumOutputs)
}

func countUniqueSegments(entries []*Entry) int {
	numUniqueSegments := 0
	numEntries := len(entries)

	for i := 0; i < numEntries; i++ {
		entry := entries[i]
		for j := 0; j < numOutput; j++ {
			if checkDigitUnique(entry.output[j]) {
				numUniqueSegments++
			}
		}
	}

	return numUniqueSegments
}

func checkDigitUnique(digit string) bool {
	switch len(digit) {
	case 2:
		// the digit is a 1
		fallthrough
	case 3:
		// the digit is a 7
		fallthrough
	case 4:
		// the digit is a 4
		fallthrough
	case 7:
		// the digit is an 8
		return true
	}

	return false
}

func sumOutputValues(entries []*Entry) int {
	sumOutput := 0

	numEntries := len(entries)
	for i := 0; i < numEntries; i++ {
		entry := entries[i]
		output, err := strconv.Atoi(entry.OutputValue())
		if err != nil {
			log.Fatal("Couldn't convert the output value to a numeric value")
		}

		sumOutput += output
	}

	return sumOutput
}

func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
