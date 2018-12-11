package aoc2018

import "testing"

func TestDayThreePartOne(t *testing.T) {
	cases := map[string]int{
		"sample_1": 4,
		//"sample_2": 4,
		//"real": 96569,
	}

	for testCase, expected := range cases {
		r := DayThreePartOne(getInput("3", testCase))

		if r != expected {
			t.Errorf("Failed for %s: expected %d but got %d\n", testCase, expected, r)
		}
	}
}

func TestDayThreePartTwo(t *testing.T) {
	cases := map[string]int{
		"sample_1": 3,
		"real":     1023,
	}

	for testCase, expected := range cases {
		r := DayThreePartTwo(getInput("3", testCase))

		if r != expected {
			t.Errorf("Failed for %s: expected %d but got %d\n", testCase, expected, r)
		}
	}
}
