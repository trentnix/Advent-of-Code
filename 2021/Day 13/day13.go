package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"sciencerocketry.com/fileprocessing"
)

type Fold struct {
	axis string
	line int
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

	maxX, maxY := 0, 0
	for _, line := range fileContents {
		// get the grid length and width
		if len(line) == 0 {
			// the coordinates are complete
			break
		}

		data := strings.Split(line, ",")
		x, _ := strconv.Atoi(data[0])
		y, _ := strconv.Atoi(data[1])

		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}
	}

	var folds []Fold

	// create a starting grid
	coordinates := make([][]bool, maxY+1)
	for i := 0; i < len(coordinates); i++ {
		coordinates[i] = make([]bool, maxX+1)
	}

	isFoldsInput := false
	for _, line := range fileContents {
		if len(line) == 0 {
			isFoldsInput = true
			continue
		}

		if !isFoldsInput {
			// the current line is an x,y coordinate
			data := strings.Split(line, ",")
			x, _ := strconv.Atoi(data[0])
			y, _ := strconv.Atoi(data[1])

			coordinates[y][x] = true
		} else {
			// the current line is a fold instruction
			data := strings.Split(line[11:], "=")
			foldLine, _ := strconv.Atoi(data[1])

			f := new(Fold)
			f.axis = data[0]
			f.line = foldLine
			folds = append(folds, *f)
		}
	}

	partOne := FoldGrid(coordinates, folds[0].axis, folds[0].line)
	fmt.Printf("After first fold - %d dots\n\n", CountDots(partOne))

	partTwo := coordinates
	for _, f := range folds {
		partTwo = FoldGrid(partTwo, f.axis, f.line)
	}

	fmt.Printf("The activation code is:\n")
	PrintGrid(os.Stdout, partTwo)
}

func FoldGrid(g [][]bool, axis string, line int) [][]bool {
	gridSizeX, gridSizeY := len(g[0]), len(g)

	xSize, ySize := 0, 0
	xStart, yStart := 0, 0

	if axis == "x" {
		// horizontal fold at line
		left := line
		right := gridSizeX - 1 - line
		xSize = left
		if left > right {
			xStart = left - right
		}

		if left <= right {
			xStart = 0
		}

		ySize = gridSizeY
	} else {
		// vertical fold at line
		upper := line
		lower := gridSizeY - 1 - line
		ySize = upper

		if upper > lower {
			yStart = upper - lower
		}

		if upper <= lower {
			yStart = 0
		}

		xSize = gridSizeX
	}

	// make a new destination grid with the size computed above
	newGrid := make([][]bool, ySize)
	for i := 0; i < len(newGrid); i++ {
		newGrid[i] = make([]bool, xSize)
	}

	iterations := 0
	if axis == "x" {
		// fill in the early part of the new grid
		for x := 0; x < xStart; x++ {
			for y := 0; y < ySize; y++ {
				newGrid[y][x] = g[y][x]
			}
		}

		// fill in the overlap from the fold
		for x := xStart; x < xSize; x++ {
			left := line - iterations - 1
			right := line + iterations + 1
			for y := 0; y < ySize; y++ {
				newGrid[y][xSize-iterations-1] = g[y][left] || g[y][right]
			}

			iterations++
		}
	} else {
		// axis == "y"
		// fill in the early part of the new grid
		for y := 0; y < yStart; y++ {
			for x := 0; x < xSize; x++ {
				newGrid[y][x] = g[y][x]
			}
		}

		// fill in the overlap from the fold
		for y := yStart; y < ySize; y++ {
			upper := line - iterations - 1
			lower := line + iterations + 1

			for x := 0; x < xSize; x++ {
				newGrid[ySize-iterations-1][x] = g[upper][x] || g[lower][x]
			}

			iterations++
		}
	}

	return newGrid
}

func PrintGrid(w io.Writer, g [][]bool) {
	verticalSize := len(g)
	for y := 0; y < verticalSize; y++ {
		horizontalSize := len(g[y])
		for x := 0; x < horizontalSize; x++ {
			if g[y][x] {
				fmt.Fprintf(w, "#")
			} else {
				fmt.Fprintf(w, " ")
			}
		}

		fmt.Fprintf(w, "\n")
	}

	fmt.Fprintf(w, "\n")
}

func CountDots(g [][]bool) int {
	xSize := len(g[0])
	ySize := len(g)

	numDots := 0
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			if g[y][x] {
				numDots++
			}
		}
	}

	return numDots
}
