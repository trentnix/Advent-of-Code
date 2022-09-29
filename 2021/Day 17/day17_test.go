package main

import "testing"

// TestParseTargetArea() tests the input parser that describes the target area.
func TestParseTargetArea(t *testing.T) {
	var tests = []struct {
		input  string
		target TargetArea
	}{
		{"target area: x=20..30, y=-10..-5", TargetArea{Xmin: 20, Xmax: 30, Ymin: -10, Ymax: -5}},
		{"target area: x=117..164, y=-140..-89", TargetArea{Xmin: 117, Xmax: 164, Ymin: -140, Ymax: -89}},
	}

	for _, test := range tests {
		parsedTargetArea := ParseInput(test.input)

		if parsedTargetArea.Xmin != test.target.Xmin ||
			parsedTargetArea.Xmax != test.target.Xmax ||
			parsedTargetArea.Ymin != test.target.Ymin ||
			parsedTargetArea.Ymax != test.target.Ymax {
			t.Errorf("parseInput(): %s\nwant %v\ngot  %v\n", test.input, test.target, parsedTargetArea)
		}
	}
}

// TestCalculatePath() validates whether or not specified initial velocities result in
// a trajectory that passes through the target area.
func TestCalculatePath(t *testing.T) {
	var tests = []struct {
		target    TargetArea
		isHit     bool
		xVelocity int
		yVelocity int
	}{
		{TargetArea{Xmin: 20, Xmax: 30, Ymin: -10, Ymax: -5}, true, 7, 2},
		{TargetArea{Xmin: 20, Xmax: 30, Ymin: -10, Ymax: -5}, true, 6, 3},
		{TargetArea{Xmin: 20, Xmax: 30, Ymin: -10, Ymax: -5}, true, 9, 0},
		{TargetArea{Xmin: 20, Xmax: 30, Ymin: -10, Ymax: -5}, false, 17, -4},
		{TargetArea{Xmin: 20, Xmax: 30, Ymin: -10, Ymax: -5}, true, 6, 9},
		{TargetArea{Xmin: 117, Xmax: 164, Ymin: -140, Ymax: -89}, true, 15, 139},
	}

	for _, test := range tests {
		isHit, path := CalculatePath(test.xVelocity, test.yVelocity, test.target)

		if test.isHit != isHit {
			t.Errorf("CalculatePath():\nwant %t\ngot  %t at path: %v\n", test.isHit, isHit, path)
		}
	}
}

func TestFindDistinctInitialVelocityValues(t *testing.T) {
	var tests = []struct {
		target   TargetArea
		numPaths int
	}{
		{TargetArea{Xmin: 20, Xmax: 30, Ymin: -10, Ymax: -5}, 112},
		{TargetArea{Xmin: 117, Xmax: 164, Ymin: -140, Ymax: -89}, 4110},
	}

	for _, test := range tests {
		launchers := FindAllIntersectingPaths(test.target)
		numPaths := len(launchers)

		if numPaths != test.numPaths {
			t.Errorf("FindAllIntersectingPaths():\nwant %d\ngot  %d\n", test.numPaths, numPaths)
		}
	}
}
