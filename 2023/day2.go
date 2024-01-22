// this file implements the 2023 Advent of Code Day 2 assignment.
// See the readme.md for details on this assignment or visit the Advent
// of Code website: https://adventofcode.com/2023/day/2

package main

import (
	"fileprocessing"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type set struct {
	red   int
	green int
	blue  int
}

type game struct {
	sets []*set
	id   int
}

func (g game) print() {
	fmt.Printf("game.id: %d\n", g.id)
	for i, s := range g.sets {
		fmt.Printf("Set %d - Red: %d, Green: %d, Blue: %d\n", i+1, s.red, s.green, s.blue)
	}
}

// day2() wraps the two parts of the solution and returns a string (output) with the
// results of the two solves.
func day2(name string, inputFile string) string {
	fileContents, _, err := fileprocessing.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	games := make([]*game, len(fileContents))
	for i, input := range fileContents {
		games[i] = newGame(input)
	}

	sumOfPossibleGames := day2part1(games)
	sumOfPowers := day2part2(games)

	var output strings.Builder
	output.WriteString("part 1 - sum of possible games: " + strconv.Itoa(sumOfPossibleGames))
	output.WriteString("\n")
	output.WriteString("part 2 - sum of powers: " + strconv.Itoa(sumOfPowers))
	output.WriteString("\n")

	return output.String()
}

// day2part1() handles the first part of the day's challenges by determining the sum of the
// game numbers of the games that are possible based on the Elf's question on what
// games are possible if the bag only contains 12 red cubes, 13 green cubes, and 14 blue cubes
func day2part1(games []*game) int {
	redLimit := 12
	greenLimit := 13
	blueLimit := 14

	sum := 0

	for _, g := range games {
		possible := true
		for _, s := range g.sets {
			if s.red > redLimit ||
				s.green > greenLimit ||
				s.blue > blueLimit {
				// a set was pulled that has a color that exceeds the limits, making this game
				// not possible
				possible = false
			}
		}

		if possible {
			sum += g.id
		}
	}

	return sum
}

// day2part2() handles the first part of the day's challenges by determining the sum of the
// game numbers of the games that are possible based on the Elf's question on what
// games are possible if the bag only contains 12 red cubes, 13 green cubes, and 14 blue cubes
func day2part2(games []*game) int {
	sum := 0

	for _, g := range games {
		redValue := 0
		greenValue := 0
		blueValue := 0

		for _, s := range g.sets {
			if s.red > redValue {
				redValue = s.red
			}

			if s.blue > blueValue {
				blueValue = s.blue
			}

			if s.green > greenValue {
				greenValue = s.green
			}
		}

		sum += redValue * blueValue * greenValue
	}

	return sum
}

// newGame() takes the string input and parses it into a 'game' structure
func newGame(input string) *game {
	g := new(game)

	// get the game number - this may be unnecessary since they seem sequential but who knows what
	// part 2 of day 2 holds
	index := strings.Index(input, ": ")
	if index < 0 {
		log.Fatal("When parsing '" + input + "' the index for a ':' value could not be found")
	}

	id, err := strconv.Atoi(input[5:index])
	if err != nil {
		log.Fatal(err)
	}

	g.id = id

	sets := strings.Split(input[index+1:], ";")

	numSets := len(sets)
	if numSets > 0 {
		g.sets = make([]*set, numSets)
	}

	for i, thisSet := range sets {
		// clean up the input string by removing leading and trailing spaces
		thisSet = strings.Trim(thisSet, " ")

		// initialize the set
		g.sets[i] = new(set)
		g.sets[i].blue = 0
		g.sets[i].green = 0
		g.sets[i].red = 0

		values := strings.Split(thisSet, ", ")
		for _, value := range values {
			redVal := getColorValue(value, "red")
			if redVal > 0 {
				g.sets[i].red = redVal
			}

			blueVal := getColorValue(value, "blue")
			if blueVal > 0 {
				g.sets[i].blue = blueVal
			}

			greenVal := getColorValue(value, "green")
			if greenVal > 0 {
				g.sets[i].green = greenVal
			}
		}
	}

	return g
}

// getColorValue() takes a given color and, if it exists in the input string 'value', it
// returns the numeric value of the number preceding the color
//
//	e.g. 'value' of '23 red' will return the integer 23 if the 'color' parameter is 'red',
//	     but will return 0 if the 'color' parameter is 'green'
func getColorValue(value string, color string) int {
	colorIndex := strings.Index(value, color)
	if colorIndex > 0 {
		val, err := strconv.Atoi(value[:colorIndex-1])
		if err != nil {
			log.Fatal(err)
		}

		return val
	}

	return -1
}
