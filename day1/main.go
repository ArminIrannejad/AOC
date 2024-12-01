package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
	"math"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("RIP", err)
		return
	}
	defer file.Close()

	var leftValues[]int
	var rightValues []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) !=2 {
			continue
		}

	
	left, err1 := strconv.Atoi(parts[0])
	right, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		fmt.Println( "RIP here:", line)
		continue
	}

	leftValues = append(leftValues, left)
	rightValues = append(rightValues, right)
}
	sort.Ints(leftValues)
	sort.Ints(rightValues)

	totDiff := 0
	for i := 0; i < len(leftValues); i++ {
		totDiff += int(math.Abs(float64(leftValues[i] - rightValues[i])))
	}

	fmt.Println("Sum of difference", totDiff)
}
