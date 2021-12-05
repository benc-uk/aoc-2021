package main

import (
	"testing"

	"github.com/benc-uk/aoc-2021/utils"
)

var data []string

const part1Answer = 1000
const part2Answer = 2000

func TestMain(t *testing.T) {
	data = utils.LoadFileAsStringArray("./data/test")
}

func TestPart1(t *testing.T) {
	result := part1(data)
	if result != part1Answer {
		t.Errorf("Expected %d, got %d", part1Answer, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(data)

	if result != part2Answer {
		t.Errorf("Expected %d, got %d", part2Answer, result)
	}
}
