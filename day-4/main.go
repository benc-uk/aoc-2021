package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/benc-uk/aoc-2021/utils"
)

type bingoCell struct {
	number int
	marked bool
}

type bingoBoard struct {
	cells *[5][5]bingoCell
	won   bool
}

func main() {
	dataLines := utils.LoadFileAsStringArray("./data/input")
	numbers, boards := parseBingo(dataLines)

	utils.PrintAnswer(1, part1(numbers, boards))
	utils.PrintAnswer(2, part2(numbers, boards))
}

func part1(numbers []int, boards []bingoBoard) int {
	utils.DebugEnabled = false

	for _, num := range numbers {
		for boardNum, board := range boards {
			for y := 0; y < 5; y++ {
				for x := 0; x < 5; x++ {
					if board.cells[x][y].number == num {
						utils.Debug(fmt.Sprintf("Marking %02d at %d,%d on %d", num, x, y, boardNum))
						board.cells[x][y].marked = true

						if checkForWin(board, x, y) {
							fmt.Printf("\nFirst board to win was %d, with number: %d\n", boardNum, num)
							printBoard(board)
							return scoreBoard(board) * num
						}
					}
				}
			}
		}
	}

	return 0
}

func part2(numbers []int, boards []bingoBoard) int {
	winCount := 0

	for _, num := range numbers {
		for boardNum := 0; boardNum < len(boards); boardNum++ {
			// Two VERY important things; get a pointer to the board for mutation
			// - and also skip boards that already have won
			board := &boards[boardNum]
			if board.won {
				continue
			}

			for y := 0; y < 5; y++ {
				for x := 0; x < 5; x++ {
					if board.cells[x][y].number == num {
						utils.Debug(fmt.Sprintf("Marking %02d at %d,%d on %d", num, x, y, boardNum))

						board.cells[x][y].marked = true

						if checkForWin(*board, x, y) {
							// This is why we need a pointer to the board
							board.won = true
							winCount++
						}

						if winCount == len(boards) {
							fmt.Printf("\nFinal board to win was %d, with number: %d\n", boardNum, num)
							printBoard(*board)
							return scoreBoard(*board) * num
						}
					}
				}
			}
		}
	}

	return 0
}

func checkForWin(board bingoBoard, x int, y int) bool {
	total := 0
	for i := 0; i < 5; i++ {
		if board.cells[x][i].marked {
			total++
		}
	}
	if total >= 5 {
		return total >= 5
	}
	total = 0
	for i := 0; i < 5; i++ {
		if board.cells[i][y].marked {
			total++
		}
	}
	return total >= 5
}

func printBoard(board bingoBoard) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if board.cells[x][y].marked {
				fmt.Printf("\033[33m[%02d] ", board.cells[x][y].number)
			} else {
				fmt.Printf("\033[34m %02d  ", board.cells[x][y].number)
			}
		}
		fmt.Printf("\033[0m\n")
	}
}

func scoreBoard(board bingoBoard) int {
	score := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !board.cells[x][y].marked {
				score += board.cells[x][y].number
			}
		}
	}
	return score
}

// Parse data as bingo game
// I really don't like this but it works
func parseBingo(dataLines []string) ([]int, []bingoBoard) {
	numbersStrArray := strings.Split(dataLines[0], ",")
	numbers := make([]int, len(numbersStrArray))
	for i, s := range numbersStrArray {
		numbers[i], _ = strconv.Atoi(s)
	}

	boards := make([]bingoBoard, 0)
	boardRow := 0
	for i := 1; i < len(dataLines); i++ {
		line := dataLines[i]
		if line == "" {
			board := bingoBoard{}
			board.cells = &[5][5]bingoCell{}
			boards = append(boards, board)
			boardRow = 0

			continue
		}

		board := boards[len(boards)-1]
		for col := 0; col < 5; col++ {
			num, _ := strconv.Atoi(strings.TrimSpace(line[col*3 : 2+col*3]))
			board.cells[col][boardRow] = bingoCell{num, false}
		}

		boardRow++
	}

	return numbers, boards
}
