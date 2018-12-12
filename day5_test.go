package aoc2018

import "testing"

func TestDayFivePartOne(t *testing.T) {
	cases := map[string]int{
		"sample": 10,
		"real":   9390,
	}

	for testCase, expected := range cases {
		r := DayFivePartOne(getInput("5", testCase))

		if r != expected {
			t.Errorf("Failed for %s: expected %d but got %d\n", testCase, expected, r)
		}
	}
}

func TestDayFivePartTwo(t *testing.T) {
	cases := map[string]int{
		"sample": 4,
		"real":   5898,
	}

	for testCase, expected := range cases {
		r := DayFivePartTwo(getInput("5", testCase))

		if r != expected {
			t.Errorf("Failed for %s: expected %d but got %d\n", testCase, expected, r)
		}
	}
}
