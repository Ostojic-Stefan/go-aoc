package main

import (
	"fmt"
	"goaoc/utils"
	"strconv"
	"strings"
)

var input map[string][]pair

type pair struct {
	name    string
	numBags int
}

func parseBullshit(str string) pair {
	firstSpace := strings.Index(str, " ")
	lastSpace := strings.LastIndex(str, " ")
	nBags, _ := strconv.Atoi(str[:firstSpace])
	name := str[firstSpace+1 : lastSpace]
	return pair{
		name:    name,
		numBags: nBags,
	}
}

func parseInput() {
	input = make(map[string][]pair)
	lines := utils.ReadLinesAsString("input.txt")
	for _, line := range lines {
		parts1 := strings.Split(line, "contain")
		source := strings.TrimSpace(strings.Split(parts1[0], "bags")[0])
		parts2 := strings.Split(parts1[1], "  ")
		dst := strings.Split(parts2[0], ", ")
		for _, v := range dst {
			if strings.TrimSpace(v) == "no other bags." {
				input[source] = make([]pair, 0)
				continue
			}
			input[source] = append(input[source], parseBullshit(strings.TrimSpace(v)))
		}
	}
}

func init() {
	parseInput()
}

func part1() {
	res := 0
	for k, v := range input {
		if k == "shiny gold" {
			continue
		}
		stack := make([]pair, 0)
		stack = append(stack, v...)

		for len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			if top.name == "shiny gold" {
				res++
				stack = nil
				break
			}
			stack = append(stack, input[top.name]...)
		}
	}

	fmt.Println(res)
}

func dfs(name string) int {
	if len(input[name]) == 0 {
		return 0
	}
	sum := 0
	for _, v := range input[name] {
		sum += v.numBags + v.numBags*dfs(v.name)
	}
	return sum
}

func part2() {
	fmt.Println(dfs("shiny gold"))
}

func main() {
	part1()
	part2()
}
