package main

import (
	"fmt"

	"github.com/benc-uk/aoc-2021/utils"
)

func main() {
	data := utils.LoadCSVAsIntArray("./data/input")

	utils.PrintAnswer(1, simulateFishCounts(data, 80))
	utils.PrintAnswer(2, simulateFishCounts(data, 256))
}

// This works by holding counts of each fish at each stage, rather than literal values
// Order is not important, so we can just keep track of the counts, and there are only 8 possible stages
// The index represents the stage, and the value represents the number of fish in that stage
func simulateFishCounts(data []int, days int) int {
	counts := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	// Initialize counts from raw data
	for s := 0; s < len(data); s++ {
		// See, I told you this was madness
		counts[data[s]]++
	}

	for d := 1; d <= days; d++ {
		// New counts array for this day
		newCounts := make([]int, len(counts))

		newCounts[0] = counts[1]
		newCounts[1] = counts[2]
		newCounts[2] = counts[3]
		newCounts[3] = counts[4]
		newCounts[4] = counts[5]
		newCounts[5] = counts[6]
		// This takes some wrapping your head around, but when a fish reaches 0 it resets to 6
		// AND fish at stage 7 move to stage 6
		newCounts[6] = counts[7] + counts[0]
		newCounts[7] = counts[8]
		// ALSO fish at zero add new fish at stage 8
		newCounts[8] = counts[0]

		counts = newCounts
	}

	total := 0
	for _, count := range counts {
		total += count
	}
	return total
}

// This way my first try, it works, but it's NOT the most efficient way to do it
// Part 2 is not possible with this method
func simulateFishStupid(dataAll []int, days int) int {
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
