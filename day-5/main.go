package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strconv"
	"strings"

	"github.com/benc-uk/aoc-2021/utils"
)

const size = 1000

func main() {
	data := utils.LoadFileAsStringArray("./data/input")

	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}

	populateGrid(data, grid)

	utils.PrintAnswer(1, part1(grid))
	drawGrid("part-1", grid)

	// utils.PrintAnswer(2, part2(grid))
}

func part1(grid [][]int) int {
	// count any values in grid greater than 1
	count := 0

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[x][y] > 1 {
				count++
			}
		}
	}

	return count
}

func part2(grid [][]int) int {
	return 1
}

func populateGrid(data []string, grid [][]int) {
	for _, line := range data {
		coords := strings.Split(line, " -> ")
		c1 := coords[0]
		c2 := coords[1]
		x1, _ := strconv.Atoi(strings.Split(c1, ",")[0])
		y1, _ := strconv.Atoi(strings.Split(c1, ",")[1])
		x2, _ := strconv.Atoi(strings.Split(c2, ",")[0])
		y2, _ := strconv.Atoi(strings.Split(c2, ",")[1])

		// Horizontal line
		if y1 == y2 {
			if x2 > x1 {
				for i := x1; i <= x2; i++ {
					grid[i][y1]++
				}
			} else {
				for i := x2; i <= x1; i++ {
					grid[i][y1]++
				}
			}
		}

		// Veritcal line
		if x1 == x2 {
			if y2 > y1 {
				for i := y1; i <= y2; i++ {
					grid[x1][i]++
				}
			} else {
				for i := y2; i <= y1; i++ {
					grid[x1][i]++
				}
			}
		}

	}
}

func drawGrid(name string, grid [][]int) {
	out, err := os.Create("./" + name + ".png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	imgRect := image.Rect(0, 0, size, size)
	img := image.NewRGBA(imgRect)
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[x][y] == 1 {
				img.Set(x, y, color.RGBA{0, 0, 0, 128})
			} else if grid[x][y] > 1 {
				img.Set(x, y, color.RGBA{255, 0, 0, 255})
				img.Set(x+1, y+1, color.RGBA{255, 0, 0, 255})
				img.Set(x+1, y, color.RGBA{255, 0, 0, 255})
				img.Set(x, y+1, color.RGBA{255, 0, 0, 255})
			}
		}
	}

	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
