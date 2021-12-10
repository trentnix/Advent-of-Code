package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"sciencerocketry.com/stack"
)

const openers = "{([<"
const closers = "})]>"

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
		log.Fatal(fmt.Errorf("invalid input in %s", inputFile))
	}

	var delimiters []*Delimiter

	// set up delimiters
	for i := 0; i < len(openers); i++ {
		delimiters = append(delimiters, NewDelimiter(rune(openers[i]), rune(closers[i])))
	}

	partOne := totalSyntaxErrorScore(fileContents, delimiters)
	fmt.Printf("Part One - total syntax error score: %d\n", partOne)

	partTwo := getMiddleCompletionStringScore(fileContents, delimiters)
	fmt.Printf("Part Two - middle completion string score: %d\n", partTwo)
}

type Delimiter struct {
	openingCharacter rune
	closingCharacter rune
}

func NewDelimiter(open rune, close rune) *Delimiter {
	delimiter := new(Delimiter)
	delimiter.openingCharacter = open
	delimiter.closingCharacter = close
	return delimiter
}

func totalSyntaxErrorScore(data []string, delimiters []*Delimiter) (totalScore int) {
	for _, line := range data {
		s := stack.NewStack()

		lineScore := 0
		for _, character := range line {
			if s.IsEmpty() || isOpener(character) {
				s.Push(getDelimiter(character, delimiters))
				continue
			}

			currentTop, ok := s.Top().(*Delimiter)
			if !ok {
				log.Fatal(fmt.Errorf("there was an error getting the top"))
			}

			if character == currentTop.closingCharacter {
				// this is closing a valid chunk - we need to pop the stack
				s.Pop()
			} else {
				// this line is corrupted
				lineScore += getSyntaxErrorScore(character)
				break
			}
		}

		totalScore += lineScore
	}

	return
}

func getSyntaxErrorScore(character rune) int {
	switch character {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}

	return 0
}

func getMiddleCompletionStringScore(data []string, delimiters []*Delimiter) (middleScore int) {
	var scores []int

	for _, line := range data {
		s := stack.NewStack()

		for _, character := range line {
			if s.IsEmpty() || isOpener(character) {
				s.Push(getDelimiter(character, delimiters))
				continue
			}

			currentTop, ok := s.Top().(*Delimiter)
			if !ok {
				log.Fatal(fmt.Errorf("there was an error getting the top"))
			}

			if character == currentTop.closingCharacter {
				// this is closing a valid chunk - we need to pop the stack
				s.Pop()
			} else {
				// this line is corrupted
				s.Clear()
				break
			}
		}

		lineScore := 0
		for !s.IsEmpty() {
			delimiter, ok := s.Pop().(*Delimiter)
			if !ok {
				log.Fatal(fmt.Errorf("there was an error popping the stack"))
			}

			lineScore = lineScore*5 + getCompletionScore(delimiter.closingCharacter)
		}

		if lineScore > 0 {
			scores = append(scores, lineScore)
		}
	}

	sort.Ints(scores)
	if len(scores) > 0 {
		return scores[len(scores)/2]
	}

	return 0
}

func getCompletionScore(character rune) int {
	switch character {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	}

	return 0
}

func getDelimiter(open rune, delimiters []*Delimiter) *Delimiter {
	for _, delimiter := range delimiters {
		if delimiter.openingCharacter == open {
			return delimiter
		}
	}

	return nil
}

func isOpener(character rune) bool {
	return strings.Contains(openers, string(character))
}
