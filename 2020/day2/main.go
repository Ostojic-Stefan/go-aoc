package main

import (
	"fmt"
	"goaoc/utils"
	"log"
	"strconv"
	"strings"
)

func parseLine(line string) (int, int, rune, string) {
	parts := strings.Split(line, " ")
	if len(parts) < 3 {
		log.Println("Parts has length less than 3", line)
	}
	bounds := strings.Split(parts[0], "-")
	if len(bounds) < 2 {
		log.Println("Bounds has length less than 2")
	}
	low, _ := strconv.Atoi(bounds[0])
	high, _ := strconv.Atoi(bounds[1])
	str := strings.ReplaceAll(parts[1], ":", "")
	password := parts[2]

	return low, high, rune(str[0]), password
}

func part1() {
	lines := utils.ReadLinesAsString("input.txt")
	res := 0
	for _, line := range lines {
		low, high, char, str := parseLine(line)
		cnt := 0
		for _, r := range str {
			if r == char {
				cnt++
				if cnt > high {
					break
				}
			}
		}
		if cnt >= low && cnt <= high {
			res++
		}
	}
	fmt.Println(res)
}

func part2() {
	lines := utils.ReadLinesAsString("input.txt")
	res := 0
	for _, line := range lines {
		first, second, char, password := parseLine(line)
		if (rune(password[first-1]) == char) != (rune(password[second-1]) == char) {
			res++
		}
	}
	fmt.Println(res)
}

func main() {
	part1()
	part2()
}
