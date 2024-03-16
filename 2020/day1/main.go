package main

import (
	"fmt"
	"goaoc/utils"
	"slices"
)

func part1() {
	data := utils.ReadLinesAsInt("input.txt")
	mapping := map[int]int{}

	target := 2020
	for i, v := range data {
		if _, ex := mapping[v]; ex {
			fmt.Println("Result:", v*data[mapping[v]])
			break
		}
		complement := target - v
		mapping[complement] = i
	}
}

func part2() {
	target := 2020
	data := utils.ReadLinesAsInt("input.txt")
	slices.Sort(data)
	for i, v := range data {
		left, right := i+1, len(data)-1
		for left < right {
			sum := v + data[left] + data[right]
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				fmt.Println("Result:", v*data[left]*data[right])
				return
			}
		}
	}
}

func main() {
	part1()
	part2()
}
