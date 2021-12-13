package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	var inputFile string
	inputFile = os.Args[1]

	fileContents, err := ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	gamma, epsilon, err := getGammaEpsilonValues(fileContents)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 - gamma * epsilon = %d\n", gamma*epsilon)

	co2, err := getPreferredValues(fileContents, false)
	if err != nil {
		log.Fatal(err)
	}

	oxygen, err := getPreferredValues(fileContents, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 2 - oxygen (%d) * co2 (%d) = %d\n", oxygen, co2, oxygen*co2)
}

func countZeroesAndOnes(data []string, dataLength int) (zero []int, one []int) {
	zero, one = make([]int, dataLength), make([]int, dataLength)

	// iterate over each digit in each line to determinte the most common value
	for _, line := range data {
		for i := 0; i < dataLength; i++ {
			switch line[i] {
			case '0':
				zero[i] += 1
			case '1':
				one[i] += 1
			}
		}
	}

	return
}

func binaryToInteger(value string) int {
	valueLength := len(value)
	integer := 0
	for i := 0; i < valueLength; i++ {
		if value[valueLength-i-1] == '1' {
			integer += int(math.Pow(2, float64(i)))
		}
	}

	return integer
}

func getGammaEpsilonValues(data []string) (int, int, error) {
	if len(data) <= 0 {
		return 0, 0, errors.New("getGammaRate - no data to parse")
	}

	gamma, epsilon := 0, 0
	dataLength := len(data[0])
	var gammaBinary string
	var epsilonBinary string

	zero, one := countZeroesAndOnes(data, dataLength)

	for i := 0; i < dataLength; i++ {
		switch {
		case zero[i] > one[i]:
			gammaBinary += "0"
			epsilonBinary += "1"
		case zero[i] < one[i]:
			gammaBinary += "1"
			epsilonBinary += "0"
		case zero[i] == one[i]:
			return 0, 0, fmt.Errorf("index %d has an equivalent number of 0s and 1s\n", i)
		}
	}

	gamma = binaryToInteger(gammaBinary)
	epsilon = binaryToInteger(epsilonBinary)

	fmt.Printf("gamma: %s = %d\nepsilon: %s = %d\n", gammaBinary, gamma, epsilonBinary, epsilon)

	return gamma, epsilon, nil
}

func getPreferredValues(data []string, mostCommon bool) (int, error) {
	if len(data) <= 0 {
		return 0, errors.New("getPreferredValues - no data to parse")
	}

	dataLength := len(data[0])

	list := make([]string, len(data))
	copy(list, data)

	for i := 0; i < dataLength; i++ {
		j := 0
		zero, one := countZeroesAndOnes(list, dataLength)

		for {
			switch {
			case zero[i] == one[i]:
				// '0's and '1's are equally common - keep the 1s if most common is preferred, 0s if least common is preferred
				if mostCommon {
					if list[j][i] == '0' {
						list = RemoveIndex(list, j)
					} else {
						j++
					}
				} else {
					if list[j][i] == '1' {
						list = RemoveIndex(list, j)
					} else {
						j++
					}
				}
			case zero[i] > one[i]:
				// '0' is the most common value at i
				if mostCommon {
					if list[j][i] == '1' {
						list = RemoveIndex(list, j)
					} else {
						j++
					}
				} else {
					if list[j][i] != '1' {
						list = RemoveIndex(list, j)
					} else {
						j++
					}
				}
			case one[i] > zero[i]:
				// '1' is the most common value at i
				if mostCommon {
					if list[j][i] == '0' {
						list = RemoveIndex(list, j)
					} else {
						j++
					}
				} else {
					if list[j][i] != '0' {
						list = RemoveIndex(list, j)
					} else {
						j++
					}
				}
			}

			if j >= len(list) {
				break
			}
		}

		if len(list) <= 1 {
			break
		}
	}

	return binaryToInteger(list[0]), nil
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
