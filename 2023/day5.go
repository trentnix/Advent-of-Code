// this file implements the 2023 Advent of Code Day 5 assignment.
// See the readme.md for details on this assignment or visit the Advent
// of Code website: https://adventofcode.com/2023/day/5

package main

import (
	"errors"
	"fileprocessing"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

// src_dest_map is a structure that tracks the individual source-to-destination mappings
type src_dest_map struct {
	destination_start int
	source_start      int
	range_length      int
}

// almanac_map is a structure that tracks source-to-destination mapping and individual location mappings
type almanac_map struct {
	source      string
	destination string

	mappings []src_dest_map
}

// almanac is a structure that identifies the seeds and maps of a given input file
type almanac struct {
	seeds []int
	maps  []*almanac_map
}

// almanac_map.new() parses the input of a map section into an almanac_map structure
func (m *almanac_map) new(input []string) {
	if len(input) <= 0 {
		log.Fatal("no almanac_map input to process")
	}

	// parsing a string like "fertilizer-to-water map:" to extract "fertilizer" and "map"
	header := strings.Split(input[0], " ")
	mapping := strings.Split(header[0], "-")

	m.source = mapping[0]
	m.destination = mapping[2]

	for _, s := range input[1:] {
		var numMap src_dest_map
		mappingVals := strings.Split(s, " ")
		numMap.destination_start, _ = strconv.Atoi(mappingVals[0])
		numMap.source_start, _ = strconv.Atoi(mappingVals[1])
		numMap.range_length, _ = strconv.Atoi(mappingVals[2])
		m.mappings = append(m.mappings, numMap)
	}
}

// almanac_map.print() prints a given almanac_map structure
func (m *almanac_map) print() {
	fmt.Printf("Map - Source: %s, Destination: %s\n", m.source, m.destination)
	for _, mapping := range m.mappings {
		fmt.Printf("source start: %d, destination start: %d, range length: %d\n", mapping.source_start, mapping.destination_start, mapping.range_length)
	}
	fmt.Println()
}

// almanac_map.navigate() navigates a given set of mappings to find the destination
// value given the specified source value
func (m *almanac_map) navigate(source int) int {
	// sort the mappings by source from least to greatest so you know they are in order
	sort.Slice(m.mappings, func(i, j int) bool {
		return m.mappings[i].source_start < m.mappings[j].source_start
	})

	numMappings := len(m.mappings)
	if numMappings > 0 {
		if source < m.mappings[0].source_start {
			return source
		}

		if source >= m.mappings[numMappings-1].source_start+m.mappings[numMappings-1].range_length {
			return source
		}
	}

	for _, m := range m.mappings {
		if source >= m.source_start && source < m.source_start+m.range_length {
			return m.destination_start + (source - m.source_start)
		}
	}

	return source
}

// almanac.new() parses the input file into an almanac structure
func (a *almanac) new(input []string) {
	if len(input) <= 0 {
		log.Fatal("no almanac input to process")
	}

	seedsInput := strings.Split(input[0], " ")
	for i := 1; i < len(seedsInput); i++ {
		seed, err := strconv.Atoi(seedsInput[i])
		if err != nil {
			log.Fatal("There was an error processing the seeds input")
		}
		a.seeds = append(a.seeds, seed)
	}

	startIndex, endIndex, lengthInput := 2, 0, len(input)

	for currentIndex := 2; currentIndex < lengthInput; currentIndex++ {
		if len(input[currentIndex]) <= 0 {
			endIndex = currentIndex

			m := new(almanac_map)
			m.new(input[startIndex:endIndex])
			a.maps = append(a.maps, m)

			startIndex = currentIndex + 1
		}
	}

	if startIndex < lengthInput {
		m := new(almanac_map)
		m.new(input[startIndex:])
		a.maps = append(a.maps, m)
	}
}

// almanac.print() prints a given almanac structure
func (a *almanac) print() {
	fmt.Println("Seeds: ")
	for _, seed := range a.seeds {
		fmt.Printf("%d ", seed)
	}
	fmt.Println()

	for _, m := range a.maps {
		m.print()
	}
}

// almanac.seedTo() traverses the x-to-y maps, given the specified 'seed' value and returns the result
// when you've arrived at the specified 'to' value
func (a *almanac) seedTo(seed int, to string) (int, error) {
	result := seed
	//fmt.Printf("seed %d, ", result)

	for _, m := range a.maps {
		result = m.navigate(result)
		if m.destination == to {
			//fmt.Printf("%s %d\n", m.destination, result)
			return result, nil
		}

		//fmt.Printf("%s %d, ", m.destination, result)
	}

	return 0, errors.New("The destination specified (" + to + ") was not found.")
}

// day5() has you helping Island Island with their food production problem described
// in the assignment
func day5(name string, inputFile string) string {
	fileContents, _, err := fileprocessing.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	a := new(almanac)
	a.new(fileContents)

	var output strings.Builder
	output.WriteString("***** DAY 5 *****\n")
	output.WriteString("Part 1:\n")
	output.WriteString("Lowest Location: " + day5part1(a))
	output.WriteString("\n")
	output.WriteString("Part 2:\n")
	output.WriteString("Lowest Location: " + day5part2(a))

	return output.String()
}

// day5part1() has you traversing the maps on the almanac to find the lowest "location" number for
// the given seed values
func day5part1(a *almanac) string {
	lowestLocation := math.MaxInt64

	for _, seed := range a.seeds {
		val, err := a.seedTo(seed, "location")
		if err != nil {
			log.Fatal("There was an error using seedTo()")
		}

		if val < lowestLocation {
			lowestLocation = val
		}
	}

	return strconv.Itoa(lowestLocation)
}

// day5part2() has you traversing the maps on the almanac to find the lowest "location" number for
// the given seed values but the seed values are a massive range
//
// note - current implementation is brute force but there are more efficient ways to solve this
func day5part2(a *almanac) string {
	lowestLocation := math.MaxInt64

	numberOfSeedPairs := len(a.seeds) / 2

	fmt.Printf("%d of %d pairs processed", 0, numberOfSeedPairs)

	for i := 0; i < numberOfSeedPairs; i++ {
		start := a.seeds[i*2]
		end := a.seeds[i*2] + a.seeds[i*2+1]
		for seed := start; seed < end; seed++ {
			val, err := a.seedTo(seed, "location")
			if err != nil {
				log.Fatal("There was an error using seedTo()")
			}

			if val < lowestLocation {
				lowestLocation = val
			}
		}

		fmt.Print("\r\033[2K")
		// Print the updated status message
		fmt.Printf("%d of %d pairs processed", i, numberOfSeedPairs)
	}

	return strconv.Itoa(lowestLocation)
}
