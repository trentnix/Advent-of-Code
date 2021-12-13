package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const lanternFishCycle = 6
const firstCycleAdds = 2

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

	fish := strings.Split(fileContents[0], ",")
	var fishInts = []int{}
	var fishIntsPartTwo = []int{}

	for _, fishString := range fish {
		fishInt, error := strconv.Atoi(fishString)
		if error != nil {
			log.Fatal(fmt.Errorf("could not convert input value to string"))
		}

		fishInts = append(fishInts, fishInt)
	}

	numDaysPartOne := 80
	numFish := calculateNumberOfFish(fishInts, numDaysPartOne)
	fmt.Printf("%d days : %d fish\n", numDaysPartOne, numFish)

	for _, fishString := range fish {
		fishInt, error := strconv.Atoi(fishString)
		if error != nil {
			log.Fatal(fmt.Errorf("could not convert input value to string"))
		}

		fishIntsPartTwo = append(fishIntsPartTwo, fishInt)
	}

	numDaysPartTwo := 256
	numFish = calculateNumberOfFishFast(fishIntsPartTwo, numDaysPartTwo)
	fmt.Printf("%d days : %d fish\n", numDaysPartTwo, numFish)
}

func calculateNumberOfFish(fish []int, days int) int {
	// this algorithm brute forces the problem by maintaining an array of fish
	// that will grow exponentially large
	for i := 0; i < days; i++ {
		numFish := len(fish)
		for j := 0; j < numFish; j++ {
			fish[j]--
			if fish[j] < 0 {
				fish[j] = lanternFishCycle
				var newFish int
				newFish = lanternFishCycle + firstCycleAdds
				fish = append(fish, newFish)
			}
		}
	}

	return len(fish)
}

func calculateNumberOfFishFast(fish []int, numDays int) int {
	// this algorithm just keeps an array of the days and each day is assigned a
	// value representing the number of fish that will reproduce on a given day
	days := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, aFish := range fish {
		days[aFish]++
	}

	for i := 0; i < numDays; i++ {
		numberReproducing := days[0]
		for j := 1; j < len(days); j++ {
			days[j-1] = days[j]
		}

		days[lanternFishCycle] += numberReproducing
		days[lanternFishCycle+firstCycleAdds] = numberReproducing
	}

	numFish := 0
	for _, cycle := range days {
		numFish += cycle
	}

	return numFish
}
