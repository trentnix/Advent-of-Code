package main

// see the readme.md for details on this assignment

import (
	"fmt"
	"log"
	"os"

	"sciencerocketry.com/fileprocessing"
	"sciencerocketry.com/graph"
)

// main() receives a single argument - the name of the data file containing the 2-dimensional map of the cavern
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

	g := parseData(fileContents, 1)

	// If you want to print the grid, uncomment these lines.
	// fmt.Printf("Original Graph:\n")
	// g.PrintGraph(len(fileContents), len(fileContents[0]), int(fileContents[0][0]-'0'))
	// fmt.Printf("\n")

	// calculate the lowest risk from the uppermost left position to the lowest right position
	lowestTotalRisk := calculateLowestRisk(g)

	multiplier := 5
	gLarge := parseData(fileContents, multiplier)

	// calculate the lowest risk from the uppermost left position to the lowest right position
	// of a 5x5 grid of the input map
	lowestTotalRiskLargerGrid := calculateLowestRisk(gLarge)

	// If you want to print the grid, uncomment these lines.
	// fmt.Printf("Original Graph:\n")
	// gLarge.PrintGraph(len(fileContents)*multiplier, len(fileContents[0])*multiplier, int(fileContents[0][0]-'0'))
	// fmt.Printf("\n")

	fmt.Printf("Lowest total risk (1x1) = %d\n", lowestTotalRisk)
	fmt.Printf("Lowest total risk (5x5) = %d\n", lowestTotalRiskLargerGrid)
}

func calculateLowestRisk(g *graph.Graph) int {
	start := 0
	end := g.NumNodes - 1

	return g.ShortestPathFromOrigin(start, end)
}

func parseData(fileContents []string, multiplier int) *graph.Graph {
	// build a 5x5 graph of the input data set and increment each value by one per row and column iteration
	// the resulting risk value should wrap around at 9 and start at 1
	numInputLines := len(fileContents)
	lineInputLength := 0
	if numInputLines > 0 {
		lineInputLength = len(fileContents[0])
	}

	numLines := multiplier * numInputLines
	lineLength := multiplier * lineInputLength

	numElements := lineLength * numLines
	g := graph.NewGraph(numElements)

	for lineIndex := 0; lineIndex < numLines; lineIndex++ {
		for charIndex := 0; charIndex < lineLength; charIndex++ {
			thisNode := lineIndex*lineLength + charIndex

			if charIndex > 0 {
				// this isn't the first character, add an edge to the previous item
				toValue := int(fileContents[lineIndex%numInputLines][(charIndex-1)%lineInputLength] - '0')
				toValue = toValue + calculateAdditive(thisNode-1, numInputLines, lineInputLength, multiplier)
				if toValue > 9 {
					toValue = (toValue % 9)
				}

				g.AddEdge(thisNode, thisNode-1, int(toValue))
			}

			if lineIndex > 0 {
				// this isn't the first line, add an edge to the item above
				toValue := int(fileContents[(lineIndex-1)%numInputLines][charIndex%lineInputLength] - '0')
				toValue = toValue + calculateAdditive(thisNode-lineLength, numInputLines, lineInputLength, multiplier)
				if toValue > 9 {
					toValue = (toValue % 9)
				}
				g.AddEdge(thisNode, thisNode-lineLength, int(toValue))
			}

			if charIndex+1 < lineLength {
				// this isn't the last character, add an edge to the next item
				toValue := int(fileContents[lineIndex%numInputLines][(charIndex+1)%lineInputLength] - '0')
				toValue = toValue + calculateAdditive(thisNode+1, numInputLines, lineInputLength, multiplier)
				if toValue > 9 {
					toValue = (toValue % 9)
				}

				g.AddEdge(thisNode, thisNode+1, toValue)
			}

			if lineIndex+1 < numLines {
				// this isn't the last line, add an edge to the item below
				toValue := int(fileContents[(lineIndex+1)%numInputLines][charIndex%lineInputLength] - '0')
				toValue = toValue + calculateAdditive(thisNode+lineLength, numInputLines, lineInputLength, multiplier)
				if toValue > 9 {
					toValue = (toValue % 9)
				}

				g.AddEdge(thisNode, thisNode+lineLength, int(toValue))
			}
		}
	}

	return g
}

func calculateAdditive(index int, numRows int, numColumns int, multiplier int) int {
	// When duplicating the grid however many times are specified in the multiplier, the
	// weights should be adjusted adding 1 for each new grid on the x axis and 1 for each
	// new grid on the y axis. See the readme.md file for details.
	additive := 0
	// y axis
	additive += index / (numRows * numColumns * multiplier)
	// x axis
	additive += (index % (numColumns * multiplier)) / numColumns

	return additive
}
