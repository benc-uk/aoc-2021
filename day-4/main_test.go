package main

import (
	"testing"

	"github.com/benc-uk/aoc-2021/utils"
)

var data []string
var numbers []int
var boards []bingoBoard

const part1Answer = 4512
const part2Answer = 1924

func TestMain(t *testing.T) {

}

func TestPart1(t *testing.T) {
	data = utils.LoadFileAsStringArray("./data/test")
	numbers, boards = parseBingo(data)

	result := part1(numbers, boards)

	if result != part1Answer {
		t.Errorf("Expected %d, got %d", part1Answer, result)
	}
}

func TestPart2(t *testing.T) {
	data = utils.LoadFileAsStringArray("./data/test")
	numbers, boards = parseBingo(data)

	result := part2(numbers, boards)

	if result != part2Answer {
		t.Errorf("Expected %d, got %d", part2Answer, result)
	}
}
