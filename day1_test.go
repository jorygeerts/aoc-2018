package aoc2018

import "testing"

func TestDayOneTaskOne(t *testing.T) {
	r := DayOneTaskOne()

	if r != "lala" {
		t.Error("Failed")
	}
}
