package aoc2018

import (
	"math"
	"strconv"
	"strings"
)

type dayThreeGrid map[int]dayThreeGridLine
type dayThreeGridLine map[int]dayThreeLineClaims
type dayThreeLineClaims []dayThreeClaim

type dayThreeClaim struct {
	id int
	x  int
	y  int
	w  int
	h  int
}

func DayThreePartOne(input string) int {

	grid := dayThreeProduceGrid(input)

	numDoubles := 0

	for _, line := range grid {
		for _, cell := range line {
			if len(cell) > 1 {
				numDoubles++
			}
		}
	}

	return numDoubles
}

func DayThreePartTwo(input string) int {

	grid := dayThreeProduceGrid(input)

	for _, line := range grid {
		for _, cell := range line {
			if len(cell) == 1 && hasOverlap(grid, cell[0]) == false {
				return cell[0].id
			}
		}
	}

	return 0
}

func hasOverlap(g dayThreeGrid, c dayThreeClaim) bool {
	for i := c.x; i < c.x+c.w; i++ {
		for j := c.y; j < c.y+c.h; j++ {
			if len(g[i][j]) > 1 {
				return true
			}
		}
	}
	return false
}

// dayThreeProduceGrid turns a list of dayThreeLineClaims into a grid of [X][Y] = num dayThreeLineClaims
func dayThreeProduceGrid(input string) dayThreeGrid {
	//	Pass one: get the dimensions of the grid
	height, width := 0, 0

	for _, line := range strings.Split(input, "\n") {
		c := dayThreeParseLine(line)

		height = int(math.Max(float64(height), float64(c.y+c.h)))
		width = int(math.Max(float64(height), float64(c.x+c.w)))
	}

	//	Create the actual grid
	grid := make(dayThreeGrid, width)
	for i := 0; i <= width; i++ {
		grid[i] = make(dayThreeGridLine, height)
		for j := 0; j <= height; j++ {
			grid[i][j] = make(dayThreeLineClaims, 0)
		}
	}

	for _, line := range strings.Split(input, "\n") {
		c := dayThreeParseLine(line)

		for x := c.x; x < c.x+c.w; x++ {
			for y := c.y; y < c.y+c.h; y++ {
				grid[x][y] = append(grid[x][y], c)
			}
		}
	}

	//	Print the grid, for debugging
	//log.Printf("----------")
	//for i := 0; i < width; i++ {
	//	l := fmt.Sprintf("%d: ", i)
	//	for j := 0; j < height; j++ {
	//
	//		switch len(grid[i][j]) {
	//		case 0:
	//			l = l + `.`
	//			break
	//		case 1:
	//			l = l + strconv.Itoa(grid[i][j][0].id)
	//			break
	//		default:
	//			l = l + `x`
	//		}
	//	}
	//	log.Printf("%s \n", l)
	//}
	//log.Printf("----------")

	return grid
}

func dayThreeParseLine(line string) dayThreeClaim {

	c := dayThreeClaim{}

	parts := strings.Split(line, ` `)

	c.id, _ = strconv.Atoi(strings.Split(parts[0], `#`)[1])

	coord := strings.Split(parts[2], `,`)
	c.x, _ = strconv.Atoi(coord[0])
	c.y, _ = strconv.Atoi(strings.Split(coord[1], `:`)[0])

	dims := strings.Split(parts[3], `x`)
	c.w, _ = strconv.Atoi(dims[0])
	c.h, _ = strconv.Atoi(dims[1])

	return c
}
