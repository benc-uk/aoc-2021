package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func LoadFileAsIntArray(filename string) []int {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	var data []int
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}

		data = append(data, val)
	}

	return data
}
