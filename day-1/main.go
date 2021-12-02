package main

import (
	"math"

	"github.com/benc-uk/aoc-2021/utils"
)

func main() {
	data := utils.LoadFileAsIntArray("./input")

	utils.PrintAnswer(1, part1(data))
	utils.PrintAnswer(2, part2(data, 3))
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

func part2(data []int, windowSize int) int {
	prevWindow := math.MaxInt32
	count := 0
	for w := 0; w < len(data); w++ {
		windowSum := 0
		for i := 0; i < windowSize; i++ {
			if w+i >= len(data) {
				windowSum = -1
				break
			} else {
				windowSum += data[w+i]
			}
		}

		if windowSum > prevWindow {
			count++
		}
		prevWindow = windowSum
	}

	return count
}
