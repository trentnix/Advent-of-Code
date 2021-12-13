package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var inputFile string
	inputFile = os.Args[1]

	fileContents, err := readFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	numIncreases, err := countIncreases(fileContents)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("increases: %d\n", numIncreases)

	numIncreases, err = countIncreasesSlidingWindow(fileContents)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("increases sliding window: %d", numIncreases)
}

func readFile(filename string) ([]string, error) {
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

func countIncreases(fileContents []string) (int, error) {
	numIncreases := 0
	for i, currentLine := range fileContents[1:] {
		previous, err := strconv.Atoi(fileContents[i])
		if err != nil {
			return 0, err
		}
		current, err := strconv.Atoi(currentLine)
		if err != nil {
			return 0, err
		}

		if previous < current {
			numIncreases++
		}
	}

	return numIncreases, nil
}

func countIncreasesSlidingWindow(fileContents []string) (int, error) {
	numIncreases := 0
	for i, currentLine := range fileContents[3:] {
		previous2, err := strconv.Atoi(fileContents[i])
		if err != nil {
			return 0, err
		}
		previous1, err := strconv.Atoi(fileContents[i+1])
		if err != nil {
			return 0, err
		}
		previous, err := strconv.Atoi(fileContents[i+2])
		if err != nil {
			return 0, err
		}
		current, err := strconv.Atoi(currentLine)
		if err != nil {
			return 0, err
		}

		if previous+previous1+previous2 < current+previous+previous1 {
			numIncreases++
		}
	}

	return numIncreases, nil
}
