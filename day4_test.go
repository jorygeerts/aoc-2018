package aoc2018

import "testing"

func TestDayFourPartOne(t *testing.T) {
	cases := map[string]int{
		"sample": 240,
		"real":   138280,
	}

	for testCase, expected := range cases {
		r := DayFourPartOne(getInput("4", testCase))

		if r != expected {
			t.Errorf("Failed for %s: expected %d but got %d\n", testCase, expected, r)
		}
	}
}

func TestDayFourPartTwo(t *testing.T) {
	cases := map[string]int{
		"sample": 4455,
		"real":   89347,
	}

	for testCase, expected := range cases {
		r := DayFourPartTwo(getInput("4", testCase))

		if r != expected {
			t.Errorf("Failed for %s: expected %d but got %d\n", testCase, expected, r)
		}
	}
}
