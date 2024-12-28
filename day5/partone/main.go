package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := []struct {
		X string
		Y string
	}{}
	updates := [][]string{}

	isRulesSection := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			isRulesSection = false
			continue
		}

		if isRulesSection {
			parts := strings.Split(line, "|")
			if len(parts) == 2 {
				x := strings.TrimSpace(parts[0])
				y := strings.TrimSpace(parts[1])
				rules = append(rules, struct {
					X string
					Y string
				}{X: x, Y: y})
			}
		} else {
			pages := strings.Split(line, ",")
			for i := range pages {
				pages[i] = strings.TrimSpace(pages[i])
			}
			updates = append(updates, pages)
		}
	}

	totalMiddleSum := 0

	for _, update := range updates {
		pagePosition := make(map[string]int)
		for idx, page := range update {
			pagePosition[page] = idx
		}

		isValid := true
		for _, rule := range rules {
			posX, existsX := pagePosition[rule.X]
			posY, existsY := pagePosition[rule.Y]
			if existsX && existsY {
				if posX > posY {
					isValid = false
					break
				}
			}
		}

		if isValid {
			n := len(update)
			middleIndex := n / 2
			if n%2 == 0 {
				middleIndex = (n / 2) - 1
			}
			middlePageStr := update[middleIndex]

			middlePage, err := strconv.Atoi(middlePageStr)
			if err != nil {
				fmt.Println("Invalid page number:", middlePageStr)
				continue
			}

			totalMiddleSum += middlePage
		}
	}

	fmt.Println(totalMiddleSum)
}

