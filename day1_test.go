package aoc2018

import "testing"

func TestDayOneTaskOne(t *testing.T) {
	cases := map[string]int{
		"sample_1": 3,
		"sample_2": 0,
		"sample_3": -6,
		"real":     522,
	}

	for testCase, expected := range cases {
		r := DayOneTaskOne(getInput("1", testCase))

		if r != expected {
			t.Errorf("Failed for %s: expected %d but got %d\n", testCase, expected, r)
		}
	}
}

func TestDayOneTaskTwo(t *testing.T) {
	cases := map[string]int{
		"sample_4": 0,
		"sample_5": 10,
		"sample_6": 5,
		"sample_7": 14,
		"real":     73364,
	}

	for testCase, expected := range cases {
		r := DayOneTaskTwo(getInput("1", testCase))

		if r != expected {
			t.Errorf("Failed for %s: expected %d but got %d\n", testCase, expected, r)
		}
	}
}
