package main

import (
	"errors"
	"strconv"
	"testing"
)

func TestPossibleGames(t *testing.T) {
	inputs := make([]string, 0)
	inputs = append(inputs, "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	inputs = append(inputs, "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")
	inputs = append(inputs, "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
	inputs = append(inputs, "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red")
	inputs = append(inputs, "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")

	games := make([]*game, 0)
	for _, input := range inputs {
		games = append(games, newGame(input))
	}

	expectedSum := 8
	actualSum := day2part1(games)
	if expectedSum != actualSum {
		t.Errorf("expected: " + strconv.Itoa(expectedSum) + ", actual: " + strconv.Itoa(actualSum))
	}
}

func TestPowerSum(t *testing.T) {
	inputs := make([]string, 0)
	inputs = append(inputs, "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	inputs = append(inputs, "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")
	inputs = append(inputs, "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
	inputs = append(inputs, "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red")
	inputs = append(inputs, "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")

	games := make([]*game, 0)
	for _, input := range inputs {
		games = append(games, newGame(input))
	}

	expectedPowerSum := 2286
	actualPowerSum := day2part2(games)
	if expectedPowerSum != actualPowerSum {
		t.Errorf("expected: " + strconv.Itoa(expectedPowerSum) + ", actual: " + strconv.Itoa(actualPowerSum))
	}
}

func TestNewGame(t *testing.T) {
	inputs := make([]string, 0)
	inputs = append(inputs, "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	inputs = append(inputs, "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")
	inputs = append(inputs, "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
	inputs = append(inputs, "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red")
	inputs = append(inputs, "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")

	results := make([]game, 0)
	results = append(results, game{
		sets: []*set{
			{red: 4, green: 0, blue: 3},
			{red: 1, green: 2, blue: 6},
			{red: 0, green: 2, blue: 0},
		},
		id: 1},
	)

	for i, result := range results {
		err := validateGame(*newGame(inputs[i]), result)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

}

func validateGame(g1 game, g2 game) error {
	if g1.id != g2.id {
		return errors.New("game1.id != game2.id")
	}

	if len(g1.sets) != len(g2.sets) {
		return errors.New("game1.sets count is different from game2.sets count")
	}

	for i, set := range g1.sets {
		err := validateSet(set, g2.sets[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func validateSet(s1 *set, s2 *set) error {
	if s1.blue != s2.blue {
		return errors.New("blue = " + strconv.Itoa(s1.blue) + ", expected: " + strconv.Itoa(s2.blue))
	}

	if s1.red != s2.red {
		return errors.New("red = " + strconv.Itoa(s1.red) + ", expected: " + strconv.Itoa(s2.red))
	}

	if s1.green != s2.green {
		return errors.New("green = " + strconv.Itoa(s1.green) + ", expected: " + strconv.Itoa(s2.green))
	}

	return nil
}
