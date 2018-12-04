package aoc2018

import "testing"

func TestDayTwoPartOne(t *testing.T) {
	cases := map[string]int{
		"sample_1": 12,
		"real":     6200,
	}

	for testCase, expected := range cases {
		r := DayTwoTaskOne(getInput("2", testCase))

		if r != expected {
			t.Errorf("Failed for %s: expected %d but got %d\n", testCase, expected, r)
		}
	}
}

func TestDayTwoPartTwo(t *testing.T) {
	cases := map[string]string{
		"sample_2": "fgij",
		"real":     "xpysnnkqrbuhefmcajodplyzw",
	}

	for testCase, expected := range cases {
		r := DayTwoTaskTwo(getInput("2", testCase))

		if r != expected {
			t.Errorf("Failed for %s: expected %s but got %s\n", testCase, expected, r)
		}
	}
}
