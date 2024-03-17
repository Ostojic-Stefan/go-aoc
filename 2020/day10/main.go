package main

import (
	"fmt"
	"goaoc/utils"
	"slices"
)

func part1() {
	data := utils.ReadLinesAsInt("input.txt")
	slices.Sort(data)
	diff1 := 0
	diff3 := 0
	prev := 0
	for i := 0; i < len(data); i++ {
		diff := data[i] - prev
		if diff == 1 {
			diff1++
		} else if diff == 3 {
			diff3++
		}
		prev = data[i]
	}
	fmt.Println(diff1 * (diff3 + 1))
}

func main() {
	part1()
}
