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


	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if grid[i][j] != 'A' {
				continue
			}

			diag1 := false
			diag2 := false
			
			if i > 0 && j > 0 && i+1 < rows && j+1 < cols {
				if (grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S') ||
					(grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M') {
					diag1 = true
				}
			}
			

			if i > 0 && j > 0 && i+1 < rows && j+1 < cols {
				if (grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S') ||
					(grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M') {
					diag2 = true
				}
			}

			if diag1 && diag2 {
				count++
			}
		}
	}
	fmt.Println(count)
}
