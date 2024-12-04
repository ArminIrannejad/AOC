package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var grid []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	//fmt.Println(grid)


	word := "XMAS"
	rows := len(grid)
	cols := len(grid[0])

	directions := [][2]int{
		{-1, -1}, // Northwest
		{-1, 0},  // North
		{-1, 1},  // Northeast
		{0, -1},  // West
		{0, 1},   // East
		{1, -1},  // Southwest
		{1, 0},   // South
		{1, 1},   // Southeast
	}

	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range directions {
				dx := dir[0]
				dy := dir[1]
				x, y := i, j
				k := 0

				for k < len(word) {
					if x < 0 || x >= rows || y < 0 || y >= cols {
						break
					}
					if grid[x][y] != word[k] {
						break
					}
					x += dx
					y += dy
					k++
				}

				if k == len(word) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

