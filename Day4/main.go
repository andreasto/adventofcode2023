package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./Day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []string{}

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	result := CheckRows(inputs)

	log.Printf("Day 4 Part 1 Result: %v", result)
	// log.Printf("Day 4 Part 2 Result: %d", result2)
}

func CheckRows(rows []string) int {
	result := 0
	for i := 0; i < len(rows); i++ {
		split := strings.Split(rows[i][strings.Index(rows[i], ": ")+1:], "|")
		winningNumbers := strings.Split(strings.TrimSpace(split[0]), " ")
		cardNumbers := strings.Split(strings.TrimSpace(split[1]), " ")

		rowResult := 0
		for j := 0; j < len(winningNumbers); j++ {
			if winningNumbers[j] == "" {
				continue
			}
			for k := 0; k < len(cardNumbers); k++ {
				if cardNumbers[k] == "" {
					continue
				}
				if winningNumbers[j] == cardNumbers[k] {
					if rowResult == 0 {
						rowResult++
					} else {
						rowResult = rowResult * 2
					}
				}
			}
		}
		result += rowResult
	}

	return result
}
