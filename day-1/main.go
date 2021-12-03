package main

import (
	"math"

	"github.com/benc-uk/aoc-2021/utils"
)

func main() {
	data := utils.LoadFileAsIntArray("./input")

	utils.PrintAnswer(1, part1(data))
	utils.PrintAnswer(2, part2(data))
}

func part1(data []int) int {
	count := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			count++
		}
	}

	return count
}

func part2(data []int) int {
	prevWindow := math.MaxInt32
	count := 0
	for i := 0; i <= len(data)-3; i++ {
		windowSum := data[i] + data[i+1] + data[i+2]
		if windowSum > prevWindow {
			count++
		}
		prevWindow = windowSum
	}

	return count
}
