package main

import (
	"testing"

	"github.com/benc-uk/aoc-2021/utils"
)

var data []int

func TestMain(t *testing.T) {
	data = utils.LoadCSVAsIntArray("./data/test")
}

const part1Answer = 5934

func TestPart1(t *testing.T) {
	result := part1(data, 80)

	if result != part1Answer {
		t.Errorf("Expected %d, got %d", part1Answer, result)
	}

}

const part2Answer = 26984457539

func TestPart2(t *testing.T) {
	result := part1(data, 256)
	if result != part2Answer {
		t.Errorf("Expected %d, got %d", part2Answer, result)
	}
}
