package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	x1, y1 int
	x2, y2 int
	sVal   string
}

func NewLine(data string) *Line {
	if len(data) <= 0 {
		// insufficient input
		return nil
	}

	parsedData := strings.Fields(data)
	if len(parsedData) != 3 {
		//invalid input
		return nil
	}

	var err error

	line := new(Line)
	line.x1, line.y1, err = parseCoordinates(parsedData[0])
	if err != nil {
		return nil
	}

	line.x2, line.y2, err = parseCoordinates(parsedData[2])
	if err != nil {
		return nil
	}

	return line
}

func (l *Line) Print(w io.Writer) {
	fmt.Fprintf(w, "(%d,%d) -> (%d, %d)\n", l.x1, l.y1, l.x2, l.y2)
}

func parseCoordinates(coordinates string) (x int, y int, err error) {
	data := strings.Split(coordinates, ",")
	if len(data) != 2 {
		//invalid input
		return -1, -1, fmt.Errorf("could not parse the provided coordinates: %s", coordinates)
	}

	x, err = strconv.Atoi(data[0])
	if err != nil {
		return -1, -1, err
	}

	y, err = strconv.Atoi(data[1])
	if err != nil {
		return -1, -1, err
	}

	return x, y, nil
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

	var lines []*Line

	// import Line data
	for i := 0; i < numLines; i++ {
		lines = append(lines, NewLine(fileContents[i]))
	}

	fmt.Printf("Number of Lines: %d\n", numLines)

	part1 := Day5PartOne(lines)
	fmt.Printf("Number of Overlaps (part 1): %d\n", part1)

	part2 := Day5PartTwo(lines)
	fmt.Printf("Number of Overlaps (part 2): %d\n", part2)
}

func getSmallerLarger(one int, two int) (smaller int, larger int) {
	if one < two {
		return one, two
	} else {
		return two, one
	}
}

func Day5PartOne(lines []*Line) int {
	grid := make(map[string]int)

	for _, line := range lines {
		if line.x1 == line.x2 {
			// horizontal line
			smaller, larger := getSmallerLarger(line.y1, line.y2)
			for i := smaller; i <= larger; i++ {
				mapIndex := strconv.Itoa(line.x1) + "," + strconv.Itoa(i)
				grid[mapIndex]++
			}

			continue
		}

		if line.y1 == line.y2 {
			// vertical line
			smaller, larger := getSmallerLarger(line.x1, line.x2)
			for i := smaller; i <= larger; i++ {
				mapIndex := strconv.Itoa(i) + "," + strconv.Itoa(line.y1)
				grid[mapIndex]++
			}

			continue
		}
	}

	countOverlaps := 0
	for _, val := range grid {
		if val > 1 {
			countOverlaps++
		}
	}

	return countOverlaps
}

func getAbsVal(input int) int {
	if input < 0 {
		return 0 - input
	}

	return input
}

func Day5PartTwo(lines []*Line) int {
	grid := make(map[string]int)

	for _, line := range lines {
		if line.x1 == line.x2 {
			// horizontal line
			smaller, larger := getSmallerLarger(line.y1, line.y2)
			for i := smaller; i <= larger; i++ {
				mapIndex := strconv.Itoa(line.x1) + "," + strconv.Itoa(i)
				grid[mapIndex]++
			}

			continue
		}

		if line.y1 == line.y2 {
			// vertical line
			smaller, larger := getSmallerLarger(line.x1, line.x2)
			for i := smaller; i <= larger; i++ {
				mapIndex := strconv.Itoa(i) + "," + strconv.Itoa(line.y1)
				grid[mapIndex]++
			}

			continue
		}

		if (getAbsVal(line.x1 - line.x2)) == (getAbsVal(line.y1 - line.y2)) {
			// diagonal line
			lineLength := getAbsVal(line.x1-line.x2) + 1

			xModifier, yModifier := 1, 1
			if line.x1 > line.x2 {
				xModifier = -1
			}

			if line.y1 > line.y2 {
				yModifier = -1
			}

			for i := 0; i < lineLength; i++ {
				x := line.x1 + int(xModifier*i)
				y := line.y1 + int(yModifier*i)
				mapIndex := strconv.Itoa(x) + "," + strconv.Itoa(y)
				grid[mapIndex]++
			}

			continue
		}
	}

	countOverlaps := 0
	for _, val := range grid {
		if val > 1 {
			countOverlaps++
		}
	}

	return countOverlaps
}
