package main

import (
	"fmt"
	"goaoc/utils"
)

func part1() {
	lines := utils.ReadLinesAsString("input.txt")
	set := make(map[rune]bool)
	sum := 0
	for _, line := range lines {
		if line == "" {
			sum += len(set)
			set = make(map[rune]bool)
		}
		for _, r := range line {
			set[r] = true
		}
	}
	sum += len(set)
	fmt.Println(sum)
}

func sumAllPeople(freq map[rune]int, nPeople int) int {
	sum := 0
	for _, v := range freq {
		if v == nPeople {
			sum++
		}
	}
	return sum
}

func part2() {
	lines := utils.ReadLinesAsString("input.txt")
	freq := make(map[rune]int)
	nPeople := 0
	sum := 0
	for _, line := range lines {
		if line == "" {
			sum += sumAllPeople(freq, nPeople)
			nPeople = 0
			freq = make(map[rune]int)
			continue
		}
		for _, r := range line {
			freq[r]++
		}
		nPeople++
	}
	sum += sumAllPeople(freq, nPeople)
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
