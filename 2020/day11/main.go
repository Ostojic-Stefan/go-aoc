package main

import (
	"fmt"
	"goaoc/utils"
)

var (
	mat  [][]byte
	dirs []utils.Pair
)

func parseInput() {
	lines := utils.ReadLinesAsString("input.txt")
	mat = make([][]byte, len(lines))
	for i, str := range lines {
		mat[i] = make([]byte, len(str))
		for j, r := range str {
			mat[i][j] = byte(r)
		}
	}
}

func init() {
	dirs = []utils.Pair{
		{-1, -1}, {0, -1},
		{1, -1}, {-1, 0},
		{-1, 1}, {1, 0},
		{0, 1}, {1, 1},
	}

}

func simulate(nAdjacent int, cntAdjFunc func(mat [][]byte, x, y int) int) {
	parseInput()
	numOccupied := 0
	for {
		copy := utils.MatCopyByte(mat)
		for i := 0; i < len(mat); i++ {
			for j := 0; j < len(mat[0]); j++ {
				cell := mat[i][j]
				if cell == '.' {
					continue
				}
				numAdj := cntAdjFunc(copy, i, j)
				if cell == 'L' {
					if numAdj == 0 {
						mat[i][j] = '#'
						numOccupied++
					}
				} else if cell == '#' {
					if numAdj >= nAdjacent {
						mat[i][j] = 'L'
						numOccupied--
					}
				}
			}
		}
		if utils.EqMatByte(mat, copy) {
			break
		}
	}
	fmt.Println(numOccupied)
}

func part1() {
	cntAdj := func(mat [][]byte, x, y int) int {
		cnt := 0
		for _, d := range dirs {
			adj := utils.Pair{X: x, Y: y}
			adj = adj.Add(d)
			if !utils.CheckBoundsByte(mat, adj.X, adj.Y) {
				continue
			}
			if mat[adj.X][adj.Y] == '#' {
				cnt++
			}
		}
		return cnt
	}

	simulate(4, cntAdj)
}

func part2() {
	cntAdj := func(mat [][]byte, x, y int) int {
		cnt := 0
		for _, d := range dirs {
			adj := utils.Pair{X: x, Y: y}
			adj = adj.Add(d)
			for utils.CheckBoundsByte(mat, adj.X, adj.Y) {
				cell := mat[adj.X][adj.Y]
				if cell != '.' {
					if cell == '#' {
						cnt++
					}
					break
				}
				adj = adj.Add(d)
			}
		}
		return cnt
	}
	simulate(5, cntAdj)
}

func main() {
	part1()
	part2()
}
