package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
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

	rightCount := make(map[int]int)
	for _, value := range rightValues { 
		rightCount[value]++
	}

	similarityScore := 0
	for _, leftValue := range leftValues { 
		similarityScore += leftValue * rightCount[leftValue]
	}

	fmt.Println("Sim Scor:", similarityScore)
}
