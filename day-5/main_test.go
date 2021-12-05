package main

import (
	"testing"

	"github.com/benc-uk/aoc-2021/utils"
)

var data []string

const part1Answer = 5
const part2Answer = 2000

var grid [][]int

func TestMain(t *testing.T) {
	data := utils.LoadFileAsStringArray("./data/test")

	grid = make([][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}

	populateGrid(data, grid)
	//printGrid(grid)

	//utils.PrintAnswer(1, part1(grid))
	// utils.PrintAnswer(2, part2(grid))
}

func TestPart1(t *testing.T) {
	result := part1(grid)

	if result != part1Answer {
		t.Errorf("Expected %d, got %d", part1Answer, result)
	}
}

// func TestPart2(t *testing.T) {
// 	result := part2(data)

// 	if result != part2Answer {
// 		t.Errorf("Expected %d, got %d", part2Answer, result)
// 	}
// }
