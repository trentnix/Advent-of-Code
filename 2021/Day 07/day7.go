package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

	numLines := len(fileContents)
	if numLines <= 0 {
		// invalid input
		log.Fatal(fmt.Errorf("invalid input in %s\n", inputFile))
	}

	crabPositionInputs := strings.Split(fileContents[0], ",")
	var crabPositions = []int{}

	numInputs := len(crabPositionInputs)

	for i := 0; i < numInputs; i++ {
		value, err := strconv.Atoi(crabPositionInputs[i])
		if err != nil {
			log.Fatal(fmt.Errorf("There was an error converting a position input to a numeric value at index %d\n", i))
		}

		crabPositions = append(crabPositions, value)
	}

	fuelPartOne := minimizeFuelExpense(crabPositions, true)
	fmt.Printf("Part One - fuel expense: %d\n", fuelPartOne)

	fuelPartTwo := minimizeFuelExpense(crabPositions, false)
	fmt.Printf("Part Two - fuel expense: %d\n", fuelPartTwo)
}

func minimizeFuelExpense(positions []int, constantBurn bool) int {
	currentPosition := calculateMedian(positions)

	multiplier := 0

	minExpense := calculateDistanceSum(positions, currentPosition, constantBurn)
	minExpensePlus := calculateDistanceSum(positions, currentPosition+1, constantBurn)

	if minExpensePlus <= minExpense {
		multiplier = 1
		minExpense = minExpensePlus
	}

	if multiplier != 1 {
		minExpenseMinus := calculateDistanceSum(positions, currentPosition-1, constantBurn)
		if minExpenseMinus <= minExpense {
			multiplier = -1
			minExpense = minExpenseMinus
		}
	}

	if multiplier == 0 {
		// minExpense is the least expensive position
		return minExpense
	}

	i := 2
	for {
		// multiplier determines whether we checking smaller or larger positions from the original
		// position
		index := currentPosition + (multiplier * i)
		expense := calculateDistanceSum(positions, index, constantBurn)

		if expense > minExpense {
			break
		}

		minExpense = expense

		i++
	}

	return minExpense
}

func calculateDistanceSum(nums []int, position int, constantBurn bool) int {
	sum := 0
	for _, value := range nums {

		if constantBurn {
			// fuel burn is constant
			sum += calculateAbsValue(value - position)
		} else {
			// fuel burn isn't constant
			diff := calculateAbsValue(value - position)
			for i := 0; i < diff; i++ {
				sum += i + 1
			}
		}
	}

	return sum
}

func calculateMedian(nums []int) int {
	sort.Ints(nums)
	middleNum := len(nums) / 2
	if (middleNum % 2) == 1 {
		// nums length is odd - return the middle number
		return nums[middleNum]
	}

	// nums length is even - return the average of the middle two
	return (nums[(middleNum)-1] + nums[middleNum]) / 2
}

func calculateAbsValue(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}
