package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const gridSize = 10
const day1steps = 100

type Octopus struct {
	energy     int
	x, y       int
	flashCount int
	flashed    bool
}

func NewOctopus(x int, y int, energy int) *Octopus {
	o := new(Octopus)
	o.x = x
	o.y = y
	o.energy = energy
	o.flashCount = 0
	o.flashed = false
	return o
}

func (o *Octopus) reset() {
	if o.flashed {
		o.energy = 0
		o.flashed = false
	}
}

func (o *Octopus) getAdjacent() (xs []int, ys []int) {
	if o.x == 0 {
		xs = append(xs, o.x)
		xs = append(xs, o.x+1)
	}

	if o.x == gridSize-1 {
		xs = append(xs, o.x)
		xs = append(xs, o.x-1)
	}

	if len(xs) == 0 {
		xs = append(xs, o.x-1)
		xs = append(xs, o.x)
		xs = append(xs, o.x+1)
	}

	if o.y == 0 {
		ys = append(ys, o.y)
		ys = append(ys, o.y+1)
	}

	if o.y == gridSize-1 {
		ys = append(ys, o.y)
		ys = append(ys, o.y-1)
	}

	if len(ys) == 0 {
		ys = append(ys, o.y-1)
		ys = append(ys, o.y)
		ys = append(ys, o.y+1)
	}

	return
}

func (o *Octopus) addEnergy(xOriginal int, yOriginal int, octopi [][]*Octopus) {
	o.energy++

	if o.energy == 10 {
		o.flashCount++
		o.flashed = true
		xs, ys := o.getAdjacent()
		for _, y := range ys {
			for _, x := range xs {
				if !(xOriginal == x && yOriginal == y) {
					octopi[y][x].addEnergy(xOriginal, yOriginal, octopi)
				}
			}
		}
	}
}

func (o *Octopus) Print(w io.Writer) {
	if o.flashed {
		fmt.Fprintf(w, "0")
	} else {
		fmt.Fprintf(w, "%d", o.energy)
	}
}

func PrintOctopi(w io.Writer, o [][]*Octopus) {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			o[i][j].Print(w)
		}
		fmt.Fprintf(w, "\n")
	}
}

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
		log.Fatal(fmt.Errorf("invalid input in %s", inputFile))
	}

	octopi := make([][]*Octopus, gridSize)
	for i := range octopi {
		octopi[i] = make([]*Octopus, gridSize)
	}

	for y, line := range fileContents {
		for x, character := range line {
			energyValue, err := strconv.Atoi(string(character))
			if err != nil {
				log.Fatal(fmt.Errorf("invalid conversion to int at %d,%d", x, y))
			}

			octopi[y][x] = NewOctopus(x, y, energyValue)
		}
	}

	totalFlashes, simultaneous := ProcessSteps(octopi)
	fmt.Printf("Day One - Total Flashes: %d\n", totalFlashes)
	fmt.Printf("Day Two - All Flashed on Step %d\n", simultaneous)
}

func ProcessSteps(octopi [][]*Octopus) (totalFlashes int, simultaneous int) {
	step := 1
	for {
		stepFlashed := 0

		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				octopi[i][j].addEnergy(j, i, octopi)
			}
		}

		isSimultaneous := true
		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				if octopi[i][j].flashed {
					stepFlashed++
				} else {
					isSimultaneous = false
				}

				octopi[i][j].reset()
			}
		}

		if isSimultaneous && simultaneous == 0 {
			simultaneous = step
			break
		}

		if step <= day1steps {
			totalFlashes += stepFlashed
		}

		step++
	}

	return
}
