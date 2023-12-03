package main

import "testing"

func TestGetCalibrationValue(t *testing.T) {
	var tests = []struct {
		first  rune
		last   rune
		result int
	}{
		{'1', '2', 12},
		{'4', '5', 45},
		{'5', '4', 54},
		{'5', '5', 55},
	}

	for _, test := range tests {
		value := getCalibrationValue(test.first, test.last)
		if value != test.result {
			t.Errorf("getCalibrationValue():\nwant %v\ngot %v\n", test.result, value)
		}
	}
}

func TestFindFirstDigit(t *testing.T) {
	var tests = []struct {
		line   string
		result rune
	}{
		{"1abc2", '1'},
		{"pqr3stu8vwx", '3'},
		{"a1b2c3d4e5f", '1'},
		{"treb7uchet", '7'},
	}

	for _, test := range tests {
		value := findFirstDigit(test.line)
		if value != test.result {
			t.Errorf("findFirstDigit():\nwant %v\ngot %v\n", test.result, value)
		}
	}
}

func TestFindLastDigit(t *testing.T) {
	var tests = []struct {
		line   string
		result rune
	}{
		{"1abc2", '2'},
		{"pqr3stu8vwx", '8'},
		{"a1b2c3d4e5f", '5'},
		{"treb7uchet", '7'},
	}

	for _, test := range tests {
		value := findLastDigit(test.line)
		if value != test.result {
			t.Errorf("findLastDigit():\nwant %v\ngot %v\n", test.result, value)
		}
	}
}

func TestFindFirstDigitWithSubstitution(t *testing.T) {
	var tests = []struct {
		line   string
		result rune
	}{
		{"two1nine", '2'},
		{"eightwothree", '8'},
		{"abcone2threexyz", '1'},
		{"xtwone3four", '2'},
		{"4nineeightseven2", '4'},
		{"zoneight234", '1'},
		{"7pqrstsixteen", '7'},
	}

	for _, test := range tests {
		value := findFirstDigitWithSubstitution(test.line)
		if value != test.result {
			t.Errorf("findfirstDigitWithSubstitution():\nwant %v\ngot %v\n", test.result, value)
		}
	}
}

func TestFindLastDigitWithSubstitution(t *testing.T) {
	var tests = []struct {
		line   string
		result rune
	}{
		{"two1nine", '9'},
		{"eightwothree", '3'},
		{"abcone2threexyz", '3'},
		{"xtwone3four", '4'},
		{"4nineeightseven2", '2'},
		{"zoneight234", '4'},
		{"7pqrstsixteen", '6'},
	}

	for _, test := range tests {
		value := findLastDigitWithSubstitution(test.line)
		if value != test.result {
			t.Errorf("findfirstDigitWithSubstitution():\nwant %v\ngot %v\n", test.result, value)
		}
	}
}

func TestDayOnePartOne(t *testing.T) {
	var tests = []struct {
		input  []string
		result int
	}{
		{[]string{
			"1abc2",
			"pqr3stu8vwx",
			"a1b2c3d4e5f",
			"treb7uchet"}, 142},
	}

	for _, test := range tests {
		value := 0

		for _, line := range test.input {
			value += getCalibrationValue(
				findFirstDigit(line),
				findLastDigit(line))
		}

		if value != test.result {
			t.Errorf("Day One - Part 1 Test:\nwant %v\ngot %v\n", test.result, value)
		}
	}
}

func TestDayOnePartTwo(t *testing.T) {
	var tests = []struct {
		input  []string
		result int
	}{
		{[]string{
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen"}, 281},
	}

	for _, test := range tests {
		value := 0

		for _, line := range test.input {
			value += getCalibrationValue(
				findFirstDigitWithSubstitution(line),
				findLastDigitWithSubstitution(line))
		}

		if value != test.result {
			t.Errorf("Day One - Part 2 Test:\nwant %v\ngot %v\n", test.result, value)
		}
	}
}
