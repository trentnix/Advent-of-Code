package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputFile string
	inputFile = os.Args[1]

	fileContents, err := ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	position, depth, navigateError := navigate(fileContents)
	if navigateError != nil {
		log.Fatal(navigateError)
	}

	fmt.Printf("Part One - horizontal position: %d, depth: %d, multiplied %d\n", position, depth, position*depth)

	position, depth, navigateError = navigateWithAim(fileContents)
	if navigateError != nil {
		log.Fatal(navigateError)
	}

	fmt.Printf("Part Two - horizontal position: %d, depth: %d, multiplied %d\n", position, depth, position*depth)
}

func navigate(fileContents []string) (position int, depth int, err error) {
	position, depth = 0, 0
	for i, line := range fileContents[0:] {
		instruction := strings.Split(line, " ")
		if len(instruction) <= 1 {
			return -1, -1, fmt.Errorf("navigate - could not parse the specified instruction at line %d", i)
		}

		magnitude, err := strconv.Atoi(instruction[1])
		if err != nil {
			return -1, -1, err
		}

		switch instruction[0] {
		case "forward":
			position += magnitude
		case "up":
			depth -= magnitude
		case "down":
			depth += magnitude
		}
	}

	return
}

func navigateWithAim(fileContents []string) (position int, depth int, err error) {
	position, depth = 0, 0
	aim := 0

	for i, line := range fileContents[0:] {
		instruction := strings.Split(line, " ")
		if len(instruction) <= 1 {
			return -1, -1, fmt.Errorf("navigateWithAim - could not parse the specified instruction at line %d", i)
		}

		magnitude, err := strconv.Atoi(instruction[1])
		if err != nil {
			return -1, -1, err
		}

		switch instruction[0] {
		case "forward":
			position += magnitude
			depth += aim * magnitude
		case "up":
			aim -= magnitude
		case "down":
			aim += magnitude
		}
	}

	return
}
