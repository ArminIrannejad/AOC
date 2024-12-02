package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeReports := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		levelsStr := strings.Fields(line)
		levels := make([]int, len(levelsStr))
		for i, s := range levelsStr {
			n, err := strconv.Atoi(s)
			if err != nil {
				return
			}
			levels[i] = n
		}

		if isSafe(levels) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true 
	}

	direction := 0 
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff == 0 {
			return false 
		}

		absDiff := math.Abs(float64(diff))
		if absDiff < 1 || absDiff > 3 {
			return false 
		}

		currentDirection := 1
		if diff < 0 {
			currentDirection = -1
		}

		if direction == 0 {
			direction = currentDirection
		} else if direction != currentDirection {
			return false 
		}
	}

	return true
}

