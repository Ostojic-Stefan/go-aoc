package main

import (
	"fmt"
	"goaoc/utils"
	"strconv"
	"strings"
)

var instructions map[string]func(int, *state)

type state struct {
	lineNum     int
	prevLineNum int
	acc         int
	visited     map[int]bool
}

func NewState() *state {
	return &state{
		lineNum:     0,
		prevLineNum: 0,
		acc:         0,
		visited:     make(map[int]bool),
	}
}

func init() {
	instructions = make(map[string]func(int, *state))

	instructions["acc"] = func(i int, s *state) {
		s.acc += i
	}
	instructions["jmp"] = func(i int, s *state) {
		s.prevLineNum = s.lineNum
		s.lineNum += i
	}
	instructions["nop"] = func(i int, s *state) {}
}

func simulate(lines []string) (bool, int) {
	s := NewState()
	for s.lineNum < len(lines) {
		if _, ex := s.visited[s.lineNum]; ex {
			return false, s.acc
		}
		s.visited[s.lineNum] = true
		line := lines[s.lineNum]
		parts := strings.Split(line, " ")
		op := parts[0]
		val, _ := strconv.ParseInt(parts[1], 10, 32)
		fn := instructions[op]
		fn(int(val), s)
		if op != "jmp" {
			s.lineNum++
		}
	}
	return true, s.acc
}

func part1() {
	lines := utils.ReadLinesAsString("input.txt")
	_, res := simulate(lines)
	fmt.Println(res)
}

func part2() {
	lines := utils.ReadLinesAsString("input.txt")
	for i := 0; i < len(lines); i++ {
		test := make([]string, len(lines))
		copy(test, lines)
		parts := strings.Split(lines[i], " ")
		if parts[0] == "jmp" {
			test[i] = fmt.Sprintf("%s %s", "nop", parts[1])
		} else if parts[0] == "nop" {
			test[i] = fmt.Sprintf("%s %s", "jmp", parts[1])
		}
		success, res := simulate(test)
		if success {
			fmt.Println(res)
			break
		}
	}
}

func main() {
	part1()
	part2()
}
