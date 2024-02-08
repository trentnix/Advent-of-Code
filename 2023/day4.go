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

// day4() has you looking for the source of water and the elf asks you to help him
// figure out what he's won with his scratch cards
func day4(name string, inputFile string) string {
	fileContents, _, err := fileprocessing.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var cards []*card
	for _, s := range fileContents {
		c := new(card)
		_ = c.new(s)
		cards = append(cards, c)
	}

	sumWorth := day4part1(cards)
	numberOfCards := day4part2(cards)

	var output strings.Builder
	output.WriteString("***** DAY 4 *****\n")
	output.WriteString("Part 1:\n")
	output.WriteString("Scratchcards are worth: " + strconv.Itoa(sumWorth))
	output.WriteString("\n")
	output.WriteString("Part 2:\n")
	output.WriteString("Number of scratch cards with rules: " + strconv.Itoa(numberOfCards))

	return output.String()
}

type card struct {
	number         int
	winningNumbers []int
	cardNumbers    []int
	count          int
}

// new() parses the input string into a card struct
func (c *card) new(input string) (number int) {
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

	c.count = 1

	return c.number
}

// worth() tells you what the score of a given card is given the winning numbers
// that match the card numbers
func (c *card) worth() int {
	count := c.matches()
	if count == 0 {
		return 0
	}

	return int(math.Pow(2, float64(count-1)))
}

// matches() tells you how many winning numbers match the card numbers
func (c *card) matches() int {
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

	return count
}

// print() prints the value of a given card
func (c *card) print() {
	fmt.Println("Card: " + strconv.Itoa(c.number))
	fmt.Printf("Winning Numbers: %#v\n", c.winningNumbers)
	fmt.Printf("Card Numbers: %#v\n", c.cardNumbers)
	fmt.Printf("Number of Matches: %d\n", c.matches())
	fmt.Printf("Worth: %d\n", c.worth())
	fmt.Println()
}

// day4part1() sums the winning worth of all the cards
func day4part1(cards []*card) int {
	sum := 0

	for _, c := range cards {
		sum += c.worth()
	}

	return sum
}

// day4part2() determines the total number of cards given the rules outlined in part 2.
// The rules dictate that for each card, the number of winning numbers that match the card numbers
// results in having a copy of the next set of cards that equals the number of matches. So if
// card 1 has 3 matches, then you get a copy of card 2, 3, and 4. If card 2 has 3 matches, you get TWO
// copies of 3, 4, and 5 because after processing card 1, you now have two card 2s.
func day4part2(cards []*card) int {
	numberOfCards := len(cards)

	for i := 0; i < numberOfCards; i++ {
		matches := cards[i].matches()
		for j := i + 1; j <= i+matches; j++ {
			if j < numberOfCards {
				cards[j].count += cards[i].count
			}
		}
	}

	numberOfCards = 0
	for _, card := range cards {
		numberOfCards += card.count
	}

	return numberOfCards
}
