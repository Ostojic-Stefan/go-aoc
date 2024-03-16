package main

import (
	"fmt"
	"goaoc/utils"
)

func part1() {
	lines := utils.ReadLinesAsString("input.txt")
	xPos := 0
	numTrees := 0
	for _, line := range lines {
		if line[xPos%len(line)] == '#' {
			numTrees++
		}
		xPos += 3
	}
	fmt.Println(numTrees)
}

type slope struct {
	x        int
	y        int
	dx       int
	dy       int
	numTrees int
}

func (s *slope) advance() {
	s.x += s.dx
	s.y += s.dy
}

func part2() {
	lines := utils.ReadLinesAsString("input.txt")
	slopes := []slope{{0, 0, 1, 1, 0}, {0, 0, 3, 1, 0}, {0, 0, 5, 1, 0}, {0, 0, 7, 1, 0}, {0, 0, 1, 2, 0}}
	nTrees := []int{}
	for _, slope := range slopes {
		for slope.y < len(lines) {
			if lines[slope.y][slope.x%len(lines[slope.y])] == '#' {
				slope.numTrees++
			}
			slope.advance()
		}
		nTrees = append(nTrees, slope.numTrees)
	}

	res := 1
	for _, n := range nTrees {
		res *= n
	}
	fmt.Println(res)
}

func main() {
	part1()
	part2()
}
