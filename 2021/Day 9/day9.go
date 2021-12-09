package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func ReadFile(filename string) ([]string, error) {
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

	fmt.Printf("num lines: %d\n", numLines)

	sumRiskLevels, productBasinSizes := getSumRiskLevelsAndProductLargestBasins(fileContents)
	fmt.Printf("Part One - Sum of Risk Levels: %d\n", sumRiskLevels)
	fmt.Printf("Part Two - Product of Basin Sizes: %d\n", productBasinSizes)
}

const riskLevel = 1

func getSumRiskLevelsAndProductLargestBasins(data []string) (sumLowPoints int, productLargestBasins int) {
	productLargestBasins = 1
	largestBasins := []int{1, 1, 1}

	for y, line := range data {
		for x, point := range line {
			if isLowest(point, x, y, data) {
				sumLowPoints += Runetoi(point) + riskLevel
				currentBasin := make(map[string]bool)
				sizeBasin := sizeBasin(x, y, x, y, data, currentBasin)
				fmt.Printf("point: %s, (%d, %d) - basin size: %d\n", string(point), x, y, sizeBasin)

				sort.Ints(largestBasins)
				if sizeBasin > largestBasins[0] {
					largestBasins[0] = sizeBasin
				}
			}
		}
	}

	for _, largeBasin := range largestBasins {
		productLargestBasins *= largeBasin
	}

	return sumLowPoints, productLargestBasins
}

func isLowest(val rune, x int, y int, data []string) bool {
	var xs, ys []int

	xs = append(xs, x)
	ys = append(ys, y)

	switch x {
	case 0:
		xs = append(xs, x+1)
	case len(data[0]) - 1:
		xs = append(xs, x-1)
	default:
		xs = append(xs, x-1)
		xs = append(xs, x+1)
	}

	switch y {
	case 0:
		ys = append(ys, y+1)
	case len(data) - 1:
		ys = append(ys, y-1)
	default:
		ys = append(ys, y-1)
		ys = append(ys, y+1)
	}

	currentValue := Runetoi(val)
	for _, xVal := range xs {
		for _, yVal := range ys {
			if currentValue >= Runetoi(rune(data[yVal][xVal])) {
				if !(x == xVal && y == yVal) && !(x != xVal && y != yVal) {
					return false
				}
			}
		}
	}

	fmt.Printf("Low point at (%d, %d) = %d\n", x, y, currentValue)
	return true
}

func sizeBasin(x int, y int, currentX int, currentY int, data []string, currentBasin map[string]bool) int {
	var xs, ys []int

	xs = append(xs, x)
	ys = append(ys, y)

	switch x {
	case 0:
		xs = append(xs, x+1)
	case len(data[0]) - 1:
		xs = append(xs, x-1)
	default:
		xs = append(xs, x-1)
		xs = append(xs, x+1)
	}

	switch y {
	case 0:
		ys = append(ys, y+1)
	case len(data) - 1:
		ys = append(ys, y-1)
	default:
		ys = append(ys, y-1)
		ys = append(ys, y+1)
	}

	size := 1

	coordinates := getCoordinateString(x, y)

	currentValue := Runetoi(rune(data[y][x]))
	if currentValue == 9 || currentBasin[coordinates] {
		return 0
	}

	currentBasin[coordinates] = true

	for _, xVal := range xs {
		for _, yVal := range ys {
			if !(x != xVal && y != yVal) {
				// not a diagonal
				currentBasin[coordinates] = true
				size += sizeBasin(xVal, yVal, x, y, data, currentBasin)
			}
		}
	}

	return size
}

func getCoordinateString(x int, y int) string {
	xString := strconv.Itoa(x)
	yString := strconv.Itoa(y)

	return xString + "," + yString
}

func Runetoi(val rune) int {
	iVal, err := strconv.Atoi(string(val))
	if err != nil {
		log.Fatal(fmt.Errorf("Error converting the rune value %s\n", string(val)))
	}

	return iVal
}
