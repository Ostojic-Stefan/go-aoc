package main

import (
	"fmt"
	"goaoc/utils"
	"math"
)

type boundedQueue struct {
	maxlen int
	queue  []int
}

func NewBoundedQueue(maxlen int) boundedQueue {
	return boundedQueue{
		maxlen: maxlen,
		queue:  make([]int, 0),
	}
}

func (bq *boundedQueue) Add(val int) {
	bq.queue = append(bq.queue, val)
	if len(bq.queue) > bq.maxlen {
		bq.queue = bq.queue[1:] // remove first element
	}
}

func populateQueue(data []int, bq *boundedQueue) {
	i := 0
	for len(bq.queue) < bq.maxlen {
		bq.Add(data[i])
		i++
	}
}

func twoSum(data []int, target int) bool {
	freq := make(map[int]bool)
	for _, v := range data {
		if _, exists := freq[v]; exists {
			return true
		}
		diff := target - v
		freq[diff] = true
	}
	return false
}

func part1() int {
	data := utils.ReadLinesAsInt("input.txt")
	bq := NewBoundedQueue(25)
	populateQueue(data, &bq)
	for i := 25; i < len(data); i++ {
		if !twoSum(bq.queue, data[i]) {
			fmt.Println(data[i])
			return data[i]
		}
		bq.Add(data[i])
	}
	return -1
}

func part2() {
	data := utils.ReadLinesAsInt("input.txt")
	targetSum := part1()
	left, right := 0, 0
	sum := 0
	for right < len(data) {
		if sum < targetSum {
			sum += data[right]
			right++
		} else if sum > targetSum {
			sum -= data[left]
			left++
		} else {
			break
		}
	}
	mx, mn := math.MinInt32, math.MaxInt32
	for i := left; i < right-1; i++ {
		mx = utils.IntMax(mx, data[i])
		mn = utils.IntMin(mn, data[i])
	}
	fmt.Println(mx + mn)
}

func main() {
	part1()
	part2()
}
