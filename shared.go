package aoc2018

import (
	"io/ioutil"
)

func getInput(day string, name string) string {
	b, err := ioutil.ReadFile("inputs/" + day + "/" + name + ".txt")

	if err != nil {
		panic(err)
	}

	return string(b)
}
