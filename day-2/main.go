package main

import (
	"fmt"

	"github.com/benc-uk/aoc-2021/utils"
)

func main() {
	data := utils.LoadFileAsIntArray("./data/input")

	result := part1(data)
	fmt.Println("### PART 1: ", result)

	result = part2(data)
	fmt.Println("### PART 2: ", result)
}

func part1(data []int) int {
	return 1
}

func part2(data []int) int {
	return 1
}
