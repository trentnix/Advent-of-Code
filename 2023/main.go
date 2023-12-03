package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	choice := -1

	// check for a command-line argument with a preselection
	// - this will speed up testing and debugging a new day's solution
	argCount := len(os.Args)
	if argCount > 1 {
		selectionNum, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Invalid choice. Please try again.")
		} else {
			choice = selectionNum
		}
	} else {
		// there is no argument so show the full menu
		fmt.Print("\n")
		fmt.Println("Welcome to solutions for the Advent of Code 2023!")
	}

	// load the exercises that will be available in the menu
	exercises := initializeExercises()

	for {
		if choice < 0 {
			// no command-line menu choice was selected or the menu needs to be re-displayed
			selection := menu(exercises)
			selectionNum, err := strconv.Atoi(selection)
			if err != nil {
				fmt.Print("invalid choice:", err)
				continue
			} else {
				choice = selectionNum
			}
		}

		fmt.Print("\n")

		if 0 < choice && choice <= len(exercises) {
			// the choice is valid
			output := exercises[choice-1].Run()
			fmt.Print(output)
		} else if choice == 0 {
			fmt.Println("Exiting...")
			fmt.Print("\n")
			return
		} else {
			fmt.Println("Invalid choice. Please try again.")
		}

		// looping through, so any command-line choice should be rendered moot
		choice = -1
	}

}

type exercise struct {
	name   string
	input  string
	myFunc func(string, string) string
}

func (e *exercise) Run() string {
	return e.myFunc(e.name, e.input)
}

func menu(exercises []exercise) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n\n")
	fmt.Println("Pick an option below:")
	for index, ex := range exercises {
		fmt.Println((index + 1), ":", ex.name)
	}
	fmt.Println("0 : Exit")
	fmt.Print("\nChoose wisely: ")

	input, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(input)

	return choice
}

func initializeExercises() []exercise {
	var exercises = []exercise{
		{
			name:   "Day 1",
			input:  "Day 1/day.input",
			myFunc: day1,
		},
	}

	return exercises
}
