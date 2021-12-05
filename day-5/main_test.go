package main

import (
	"testing"

	"github.com/benc-uk/aoc-2021/utils"
)

var data []string

const part1Answer = 5
const part2Answer = 12

func TestMain(t *testing.T) {
	data = utils.LoadFileAsStringArray("./data/test")
	size = 10
}

func TestPart1(t *testing.T) {
	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}

	populateGrid(data, grid, false)
	result := countVents(grid)

	if result != part1Answer {
		t.Errorf("Expected %d, got %d", part1Answer, result)
	}
}

func TestPart2(t *testing.T) {
	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}

	populateGrid(data, grid, true)
	result := countVents(grid)

	if result != part2Answer {
		t.Errorf("Expected %d, got %d", part2Answer, result)
	}
}
