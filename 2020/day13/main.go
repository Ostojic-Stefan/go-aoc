package main

import (
	"fmt"
	"goaoc/utils"
	"math"
	"strconv"
	"strings"
)

func part1() {
	lines := utils.ReadLinesAsString("input.txt")
	base, _ := strconv.Atoi(lines[0])
	times := lines[1]

	earliest := math.MaxInt32
	busID := math.MaxInt32

	for _, time := range strings.Split(times, ",") {
		if time == "x" {
			continue
		}
		num, _ := strconv.Atoi(time)
		t := (base/num)*num + num
		if t < earliest {
			busID = num
			earliest = t
		}
	}
	fmt.Println(busID * (earliest - base))
}

func main() {
	part1()
}
