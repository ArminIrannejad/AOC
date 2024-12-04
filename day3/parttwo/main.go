package main

import (
    "fmt"
	"os"
    "regexp"
    "strconv"
)

func main() {
    data, err := os.ReadFile("input.txt")
    if err != nil {
        return
    }
    input := string(data)

    re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
    matches := re.FindAllStringSubmatch(input, -1)

    total := 0
    enabled := true

    for _, match := range matches {
        if match[0] == "do()" {
            enabled = true
        } else if match[0] == "don't()" {
            enabled = false
        } else if match[1] != "" && match[2] != "" {
            if enabled {
                x, err1 := strconv.Atoi(match[1])
                y, err2 := strconv.Atoi(match[2])
                if err1 != nil || err2 != nil {
                    fmt.Println("Error here mby??", err1, err2)
                    continue
                }

                total += x * y
            }
        }
    }

    fmt.Println(total)
}

