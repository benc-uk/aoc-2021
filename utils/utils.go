package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var DebugEnabled = false

func PrintAnswer(part int, answer int) {
	fmt.Printf(" ✨ Part %d answer: \033[35m%v\033[0m \n", part, answer)
}

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

func LoadCSVAsIntArray(filename string) []int {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	strValues := strings.Split(string(data), ",")
	var intValues []int
	for _, val := range strValues {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalln(err)
		}
		intValues = append(intValues, intVal)
	}

	return intValues
}

func LoadFileAsStringArray(filename string) []string {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	var data []string
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if err != nil {
			log.Fatalln(err)
		}

		data = append(data, scanner.Text())
	}

	return data
}

func Debug(s string) {
	if DebugEnabled {
		fmt.Printf("\033[35m%s\n\033[0m", s)
	}
}

func PrintArray(a []int) {
	fmt.Println(strings.Join(strings.Fields(fmt.Sprint(a)), ","))
}
