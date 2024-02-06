// this file implements the 2023 Advent of Code Day 2 assignment.
// See the readme.md for details on this assignment or visit the Advent
// of Code website: https://adventofcode.com/2023/day/2

package main

import (
	"fileprocessing"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// day3()
func day4(name string, inputFile string) string {
	fileContents, _, err := fileprocessing.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var cards []card
	for _, s := range fileContents {
		cards = append(cards, newCard(s))
	}

	sumWorth := day4part1(cards)

	var output strings.Builder
	output.WriteString("Part 1:\n")
	output.WriteString("Scratchcards are worth: " + strconv.Itoa(sumWorth))

	return output.String()
}

type card struct {
	number         int
	winningNumbers []int
	cardNumbers    []int
}

func (c card) worth() int {
	elementsMap := make(map[int]bool)

	for _, elem := range c.winningNumbers {
		elementsMap[elem] = true
	}

	count := 0
	for _, elem := range c.cardNumbers {
		if _, found := elementsMap[elem]; found {
			count++
		}
	}

	if count == 0 {
		return 0
	}

	return int(math.Pow(2, float64(count-1)))
}

func (c card) print() {
	fmt.Println("Card: " + strconv.Itoa(c.number))
	fmt.Printf("Winning Numbers: %#v\n", c.winningNumbers)
	fmt.Printf("Card Numbers: %#v\n", c.cardNumbers)
	fmt.Println()
}

// day4part1()
func day4part1(cards []card) int {
	sum := 0

	for _, c := range cards {
		sum += c.worth()
	}

	return sum
}

func newCard(input string) card {
	var c card

	parts := strings.Split(input, "|")
	cardNumbers := strings.Split(strings.TrimSpace(parts[1]), " ")
	for _, cardNumber := range cardNumbers {
		if len(cardNumber) > 0 {
			number, err := strconv.Atoi(cardNumber)
			if err != nil {
				log.Fatal(err)
			}

			c.cardNumbers = append(c.cardNumbers, number)
		}
	}

	firstPart := strings.Split(parts[0], ":")
	winningNumbers := strings.Split(strings.TrimSpace(firstPart[1]), " ")
	for _, winningNumber := range winningNumbers {
		if len(winningNumber) > 0 {
			number, err := strconv.Atoi(winningNumber)
			if err != nil {
				log.Fatal(err)
			}

			c.winningNumbers = append(c.winningNumbers, number)
		}
	}

	initialPart := strings.Split(firstPart[0], " ")
	numStrings := len(initialPart)
	if numStrings > 0 {
		cardNumber, err := strconv.Atoi(initialPart[numStrings-1])
		if err != nil {
			log.Fatal(err)
		}

		c.number = cardNumber
	}

	return c
}
