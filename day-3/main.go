package main

import (
	"fmt"
	"strconv"

	"github.com/benc-uk/aoc-2021/utils"
)

func main() {
	data := utils.LoadFileAsStringArray("./data/input")

	utils.PrintAnswer(1, part1(data))
	utils.PrintAnswer(2, part2(data))
}

func part1(data []string) int {
	bitcount := len(data[0])

	// First count the '1' bits in each line reading
	gamma := 0
	for bitIndex := 0; bitIndex < bitcount; bitIndex++ {
		countOfOnes := 0
		for _, reading := range data {
			if reading[bitIndex] == '1' {
				countOfOnes++
			}

			if countOfOnes > len(data)/2 {
				// Note we need to set the bits from the other end of the number
				// Endianness is a bitch. Taken from https://stackoverflow.com/a/23192263/1343261
				gamma |= 1 << (bitcount - bitIndex - 1)
			}
		}
	}

	// Flip all the bits to get epsilon
	// Taken from https://stackoverflow.com/a/66396727/1343261
	epsilon := gamma ^ ((1 << bitcount) - 1)
	return int(gamma) * int(epsilon)
}

func part2(data []string) int {
	oxygenString := lifeSupportRating(data, 0, true)
	co2String := lifeSupportRating(data, 0, false)

	utils.Debug(fmt.Sprintf("Oxygen: %s", oxygenString))
	utils.Debug(fmt.Sprintf("C02:    %s", co2String))

	oxygen, _ := strconv.ParseInt(oxygenString, 2, 32)
	co2, _ := strconv.ParseInt(co2String, 2, 32)
	return int(oxygen) * int(co2)
}

// Recursive function to find the oxygen or CO2 rating
func lifeSupportRating(data []string, bitIndex int, countMost bool) string {
	if len(data) == 1 {
		return data[0]
	}

	countOfOnes := 0
	var bitFilter byte = '0'
	for _, dataLine := range data {
		if dataLine[bitIndex] == '1' {
			countOfOnes++
		}
		// I can't understand why I need to use len-1, but I have to
		if countOfOnes > (len(data)-1)/2 {
			bitFilter = '1'
			break
		}
	}

	if !countMost {
		if bitFilter == '1' {
			bitFilter = '0'
		} else {
			bitFilter = '1'
		}
	}

	var filtered []string
	for _, dataLine := range data {
		if dataLine[bitIndex] == bitFilter {
			filtered = append(filtered, dataLine)
		}
	}

	// Recursion baby!
	return lifeSupportRating(filtered, bitIndex+1, countMost)
}
