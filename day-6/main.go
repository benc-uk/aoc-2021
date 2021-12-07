package main

import (
	"fmt"

	"github.com/benc-uk/aoc-2021/utils"
)

func main() {
	data := utils.LoadCSVAsIntArray("./data/input")

	utils.PrintAnswer(1, part1(data, 80))
	utils.PrintAnswer(2, part1(data, 256))
}

func part1(dataAll []int, days int) int {
	count := 0
	for s := 0; s < len(dataAll); s++ {
		data := make([]int, 1)
		data[0] = dataAll[s]
		for day := 1; day <= days; day++ {
			newFish := 0
			for f := 0; f < len(data); f++ {
				data[f]--
				if data[f] < 0 {
					data[f] = 6
					newFish++
				}
			}
			for f := 0; f < newFish; f++ {
				data = append(data, 8)
			}

			fmt.Printf("Day %d, slice %d: %d fish\n", day, s, len(data))
			utils.Debug(fmt.Sprintf("Day %d: %v\n", day, data))
		}
		count += len(data)
	}

	return count
}
