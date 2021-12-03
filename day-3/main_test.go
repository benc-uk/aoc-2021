package main

import (
	"testing"

	"github.com/benc-uk/aoc-2021/utils"
)

var data []string

func TestMain(t *testing.T) {
	data = utils.LoadFileAsStringArray("./data/test")
}

func TestPart1(t *testing.T) {
	res := part1(data)
	if res != 198 {
		t.Errorf("Expected 198, got %d", res)
	}
}

func TestPart2(t *testing.T) {
	res := part2(data)
	if res != 230 {
		t.Errorf("Expected 230, got %d", res)
	}
}
