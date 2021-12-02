package main

import (
	"strconv"
	"strings"

	"github.com/benc-uk/aoc-2021/utils"
)

func main() {
	data := utils.LoadFileAsStringArray("./input")

	utils.PrintAnswer(1, part1(data))
	utils.PrintAnswer(2, part2(data))
}

func part1(data []string) int {
	depth := 0
	pos := 0

	for _, instructionLine := range data {
		instruction := strings.Split(instructionLine, " ")
		value, _ := strconv.Atoi(instruction[1])

		switch instruction[0] {
		case "forward":
			pos += value
		case "up":
			depth -= value
		case "down":
			depth += value
		}
	}
	return pos * depth
}

func part2(data []string) int {
	depth := 0
	pos := 0
	aim := 0

	for _, instructionLine := range data {
		instruction := strings.Split(instructionLine, " ")
		value, _ := strconv.Atoi(instruction[1])

		switch instruction[0] {
		case "forward":
			pos += value
			depth += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}
	return pos * depth
}
