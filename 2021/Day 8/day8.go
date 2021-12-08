package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const numPatterns = 10
const numOput = 4

type Entry struct {
	patterns []string
	output   []string
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

	return entry
}

func (e *Entry) Print(w io.Writer) {
	fmt.Fprintf(w, "patterns: %v - output: %v\n", e.patterns, e.output)
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
		entries[i].Print(os.Stdout)
	}
}
