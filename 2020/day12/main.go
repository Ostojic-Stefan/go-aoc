package main

import (
	"fmt"
	"goaoc/utils"
	"log"
	"strconv"
)

type state struct {
	xPos int
	yPos int
	deg  int
}

const (
	DIR_EAST  = 0
	DIR_NORTH = 90
	DIR_WEST  = 180
	DIR_SOUTH = 270
)

func part1() {
	s := state{0, 0, DIR_EAST}
	instructions := map[string]func(int){}
	instructions["N"] = func(i int) {
		s.yPos -= i
	}
	instructions["S"] = func(i int) {
		s.yPos += i
	}
	instructions["E"] = func(i int) {
		s.xPos += i
	}
	instructions["W"] = func(i int) {
		s.xPos -= i
	}
	instructions["L"] = func(i int) {
		s.deg = (s.deg + i) % 360
	}
	instructions["R"] = func(i int) {
		s.deg -= i
		if s.deg < 0 {
			s.deg = s.deg + 360
		}
	}
	instructions["F"] = func(i int) {
		switch s.deg {
		case DIR_EAST:
			s.xPos += i
		case DIR_NORTH:
			s.yPos -= i
		case DIR_WEST:
			s.xPos -= i
		case DIR_SOUTH:
			s.yPos += i
		default:
			log.Fatalf("Invalid direction orientation %d", s.deg)
		}
	}

	lines := utils.ReadLinesAsString("input.txt")
	for _, line := range lines {
		num, _ := strconv.Atoi(line[1:])
		op := line[0:1]
		instructions[op](num)
	}
	fmt.Println(s)
	fmt.Println(s.xPos + s.yPos)
}

func main() {
	part1()
}
