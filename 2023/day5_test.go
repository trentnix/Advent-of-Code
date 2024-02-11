package main

import (
	"strings"
	"testing"
)

func TestLowestLocationNumber(t *testing.T) {
	inputText := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

	input := strings.Split(inputText, "\n")
	a := new(almanac)
	a.new(input)

	// for _, seed := range a.seeds {
	// 	val, err := a.seedTo(seed, "location")
	// 	if err != nil {
	// 		log.Fatal("seedTo error")
	// 	}
	// 	fmt.Printf("seedTo(): %d, 'location': %d\n", seed, val)
	// }

	result := day5part1(a)
	expectedResult := "35"
	if result != expectedResult {
		t.Errorf("expected: %s, got: %s", expectedResult, result)
	}
}

func TestLowestLocationNumberWithPart2Rules(t *testing.T) {
	inputText := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

	input := strings.Split(inputText, "\n")
	a := new(almanac)
	a.new(input)

	// for _, seed := range a.seeds {
	// 	val, err := a.seedTo(seed, "location")
	// 	if err != nil {
	// 		log.Fatal("seedTo error")
	// 	}
	// 	fmt.Printf("seedTo(): %d, 'location': %d\n", seed, val)
	// }

	result := day5part2(a)
	expectedResult := "46"
	if result != expectedResult {
		t.Errorf("expected: %s, got: %s", expectedResult, result)
	}
}
