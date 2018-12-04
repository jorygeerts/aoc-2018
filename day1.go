package aoc2018

import (
	"strconv"
	"strings"
)

func DayOneTaskOne(changes string) int {
	steps := strings.Split(changes, ", ")
	result := 0

	for _, step := range steps {
		stepChange, err := strconv.Atoi(step)

		if err != nil {
			panic(err)
		}

		result += stepChange
	}

	return result
}

func DayOneTaskTwo(changes string) int {
	steps := strings.Split(changes, ", ")
	freqs := make(map[int]bool, 0)

	result := 0
	freqs[result] = true
	for {
		for _, step := range steps {
			stepChange, err := strconv.Atoi(step)

			if err != nil {
				panic(err)
			}

			result += stepChange

			//log.Printf("Step %s brings us to %d\n", step, result)

			if _, ok := freqs[result]; ok {
				return result
			}

			freqs[result] = true
		}
	}
}
