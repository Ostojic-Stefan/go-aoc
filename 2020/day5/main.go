package main

import (
	"fmt"
	"goaoc/utils"
	"math"
)

func getSeatRowCol(line string) (int, int) {
	l, r := 0, 127
	for i := 0; i < 7; i++ {
		mid := (r + l) / 2
		if line[i] == 'F' {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	c, d := 0, 7
	for i := 7; i < 10; i++ {
		mid := (c + d) / 2
		if line[i] == 'R' {
			c = mid + 1
		} else {
			d = mid - 1
		}
	}
	return l, c
}

func part1() {
	lines := utils.ReadLinesAsString("input.txt")
	maxId := math.MinInt32
	for _, line := range lines {
		row, col := getSeatRowCol(line)
		maxId = utils.IntMax(maxId, row*8+col)
	}
	fmt.Println(maxId)
}

func part2() {
	seats := [128 * 8]bool{}
	lines := utils.ReadLinesAsString("input.txt")
	for _, line := range lines {
		row, col := getSeatRowCol(line)
		seats[row*8+col] = true
	}
	i := 0
	for i < len(seats) && !seats[i] {
		i++
	}
	for i < len(seats) && seats[i] {
		i++
	}
	row := i / 8
	col := i % 8
	fmt.Println(row*8 + col)
}

func main() {
	part1()
	part2()
}
