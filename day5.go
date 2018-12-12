package aoc2018

import (
	"math"
	"strings"
)

func DayFivePartOne(input string) int {

	for i := 0; i < len(input)-1; i++ {
		a, b := input[i:i+1], input[i+1:i+2]

		if strings.ToLower(a) == strings.ToLower(b) && a != b {
			return DayFivePartOne(input[0:i] + input[i+2:len(input)])
		}
	}

	return len(input)
}

func DayFivePartTwo(input string) int {

	uniqueChars := make(map[string]bool, 0)
	chars := make([]string, 0)

	for i := 0; i < len(input); i++ {
		c := strings.ToLower(input[i : i+1])

		if _, ok := uniqueChars[c]; ok == false {
			uniqueChars[c] = true
			chars = append(chars, c)
		}
	}

	shortest := math.MaxInt32
	for _, c := range chars {
		r := DayFivePartOne(strings.Replace(strings.Replace(input, c, ``, -1), strings.ToUpper(c), ``, -1))

		if r < shortest {
			shortest = r
		}
	}

	return shortest
}
