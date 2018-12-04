package aoc2018

import (
	"strings"
)

func DayTwoTaskOne(input string) int {

	ids := strings.Split(input, "\n")

	doubleCount := 0
	tripleCount := 0

	for _, id := range ids {
		hasDouble := false
		hasTriple := false

		for _, letter := range strings.Split(id, "") {
			count := strings.Count(id, letter)

			if count == 2 {
				hasDouble = true
			}

			if count == 3 {
				hasTriple = true
			}
		}

		if hasDouble {
			doubleCount++
		}

		if hasTriple {
			tripleCount++
		}
	}

	return doubleCount * tripleCount
}

func DayTwoTaskTwo(input string) string {

	ids := strings.Split(input, "\n")

	for _, a := range ids {
		for _, b := range ids {
			if a == b {
				continue
			}
			same := ""
			numDiffs := 0

			for pos, char := range a {
				if b[pos:pos+1] == string(char) {
					same += string(char)
				} else {
					numDiffs++
				}
			}

			if numDiffs == 1 {
				return same
			}

		}
	}
	return ""
}
