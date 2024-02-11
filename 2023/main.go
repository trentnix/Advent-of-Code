// package main implements (some...odds are I won't do them all) solutions to
// the exercises in the 2023 Advent of Code!
//
// The Advent of Code gives me a good excuse to shake off at least a little of
// the rust (little 'r') and atrophy that sets in when you don't code as much
// as you'd like.
//
// And I know, I tend to over-comment. But if you noticed, then you're reading my
// code and that's a good reason why comments might be useful! It's also a good
// refresh to me as to what I was thinking in a given moment (and to make sure I
// don't forget how Go Doc comments work).
//
// Usage:
//
//	2023 [input]
//
// 'input' represents a selection you'd like to run and can be omitted. If it
// is omitted, a menu is displayed and user input is requested to choose an
// exercise to run.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// main() is where the action starts (and, unless something goes badly, ends).
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
			// 'selection' captures the user's selection for processing
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

		if 0 < choice && choice <= len(exercises) { // the choice is valid
			output := exercises[choice-1].run()
			fmt.Print(output)
		} else if choice == 0 {
			fmt.Println("Exiting...")
			fmt.Print("\n")
			return
		} else {
			fmt.Println("Invalid choice. Please try again.")
		}

		// We are looping through, so any command-line choice should be rendered moot.
		// Resetting choice does just that.
		choice = -1
	}

}

// exercise is a structure for storing a given day's name, input data, and processing
// function. An array of these will be used to provide the user a menu as well as handle
// the processing required to solve the exercise.
type exercise struct {
	name   string
	input  string
	myFunc func(string, string) string
}

// run() executes the solve for a given exercise, passing both the name and input
// that was associated with the exercise
func (e *exercise) run() string {
	return e.myFunc(e.name, e.input)
}

// initializeExercises() builds the exercises array
func initializeExercises() []exercise {
	var exercises = []exercise{
		{
			name:   "Day 1",
			input:  "Day 1/day.input",
			myFunc: day1,
		},
		{
			name:   "Day 2",
			input:  "Day 2/day.input",
			myFunc: day2,
		},
		{
			name:   "Day 3",
			input:  "Day 3/day.input",
			myFunc: day3,
		},
		{
			name:   "Day 4",
			input:  "Day 4/day.input",
			myFunc: day4,
		},
		{
			name:   "Day 5",
			input:  "Day 5/day.input",
			myFunc: day5,
		},
	}

	return exercises
}

// menu() takes an array of exercises, builds a command-line menu to present to the
// user, and returns (as a string value) the selection made by the user
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
