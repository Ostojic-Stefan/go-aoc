package main

import (
	"fmt"
	"goaoc/utils"
	"log"
	"math"
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

func part2() {
	s := state{10, 1, DIR_EAST}
	shipx, shipy := 0, 0
	instructions := map[string]func(int){}
	instructions["N"] = func(i int) {
		s.yPos += i
	}
	instructions["S"] = func(i int) {
		s.yPos -= i
	}
	instructions["E"] = func(i int) {
		s.xPos += i
	}
	instructions["W"] = func(i int) {
		s.xPos -= i
	}
	instructions["L"] = func(i int) {
		if i == 90 {
			s.xPos, s.yPos = -s.yPos, s.xPos
		} else if i == 180 {
			s.xPos, s.yPos = -s.xPos, -s.yPos
		} else {
			s.xPos, s.yPos = s.yPos, -s.xPos
		}
	}
	instructions["R"] = func(i int) {
		if i == 90 {
			s.xPos, s.yPos = s.yPos, -s.xPos
		} else if i == 180 {
			s.xPos, s.yPos = -s.xPos, -s.yPos
		} else {
			s.xPos, s.yPos = -s.yPos, s.xPos
		}
	}
	instructions["F"] = func(i int) {
		shipx += i * s.xPos
		shipy += i * s.yPos
	}
	lines := utils.ReadLinesAsString("input.txt")
	for _, line := range lines {
		num, _ := strconv.Atoi(line[1:])
		op := line[0:1]
		instructions[op](num)
	}
	fmt.Println(math.Abs(float64(shipx)) + math.Abs(float64(shipy)))
}

func main() {
	// part1()
	part2()
}
