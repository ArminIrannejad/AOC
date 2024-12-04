package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"regexp"
)
	
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	total := 0	

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			x,_ := strconv.Atoi(match[1])
			y,_ := strconv.Atoi(match[2])
			total += x * y
		}
	}
	fmt.Println(total)
}
