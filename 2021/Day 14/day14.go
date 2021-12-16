package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"sciencerocketry.com/fileprocessing"
)

var characterCount = map[byte]uint64{}

type Pair struct {
	pair  string
	count uint64
}

func NewPair(p string, c uint64) *Pair {
	var pair = new(Pair)
	pair.pair = p
	pair.count = c
	return pair
}

func resetGlobals() {
	characterCount = map[byte]uint64{}
}

func main() {
	inputFile := os.Args[1]

	fileContents, err := fileprocessing.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	numLines := len(fileContents)
	if numLines <= 0 {
		// invalid input
		log.Fatal(fmt.Errorf("invalid input in %s", inputFile))
	}

	template := fileContents[0]
	rules := make(map[string]string)

	for _, line := range fileContents[2:] {
		rule := strings.Split(line, " -> ")
		rules[rule[0]] = rule[1]
	}

	fmt.Printf("Template:       %s\n", template)

	iterations := 10
	ApplyInsertionRules(template, rules, iterations)
	fmt.Printf("Part One - (Day 10) Most Common - Least Common = %d\n", SubtractLeastCommonFromMostCommon())

	resetGlobals()

	iterations = 40
	ApplyInsertionRules(template, rules, iterations)
	fmt.Printf("Part Two - (Day 40) Most Common - Least Common = %d\n", SubtractLeastCommonFromMostCommon())
}

func ApplyInsertionRules(template string, rules map[string]string, iterations int) {
	if len(template) == 0 {
		return
	}

	var templatePairs = map[string]*Pair{}
	var newPairs = map[string]*Pair{}

	// divide into pairs
	for i := 0; i < len(template)-1; i++ {
		characterCount[template[i]]++

		var thisPair = new(Pair)
		thisPair.pair = template[i : i+2]
		thisPair.count = 1

		if templatePairs[thisPair.pair] == nil {
			templatePairs[thisPair.pair] = thisPair
		} else {
			templatePairs[thisPair.pair].count++
		}
	}

	characterCount[template[len(template)-1]]++

	for iteration := 0; iteration < iterations; iteration++ {
		for key, pair := range templatePairs {
			ruleResult, ok := rules[key]
			if ok {
				leftPair := key[0:1] + ruleResult
				rightPair := ruleResult + key[1:2]

				if _, exists := newPairs[leftPair]; !exists {
					newPairs[leftPair] = NewPair(leftPair, pair.count)
					characterCount[ruleResult[0]] += pair.count
				} else {
					characterCount[ruleResult[0]] += pair.count
					newPairs[leftPair].count += pair.count
				}

				if _, exists := newPairs[rightPair]; !exists {
					newPairs[rightPair] = NewPair(rightPair, pair.count)
				} else {
					newPairs[rightPair].count += pair.count
				}
			} else {
				if _, exists := newPairs[pair.pair]; !exists {
					newPairs[pair.pair] = NewPair(pair.pair, pair.count)
				} else {
					newPairs[pair.pair].count += pair.count
				}
			}
		}

		for key := range templatePairs {
			delete(templatePairs, key)
		}

		templatePairs = newPairs
		newPairs = map[string]*Pair{}
	}
}

func SubtractLeastCommonFromMostCommon() uint64 {
	var mostCommonCharacter, leastCommonCharacter uint64

	for _, count := range characterCount {
		if leastCommonCharacter == 0 || count < leastCommonCharacter {
			leastCommonCharacter = count
		}

		if count > mostCommonCharacter {
			mostCommonCharacter = count
		}

	}

	return mostCommonCharacter - leastCommonCharacter
}
