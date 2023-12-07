package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Day3/testInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []string{}

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	numbers := []int{}
	numbers = append(numbers, IsAdjacentToASymbol(inputs)...)

	result := 0
	for _, v := range numbers {
		result += v
	}

	log.Printf("Day 2 Part 1 Result: %v", result)
	// log.Printf("Day 2 Part 2 Result: %d", result2)
}

func CheckRow(row string) bool {
	crr := strings.Split(row, "")
	hit := false
	for _, v := range crr {
		_, pErr := strconv.Atoi(v)
		if v == "." || pErr == nil {
			// Skip because it's either a dot or a number
			continue
		}
		hit = true
		break
	}

	return hit
}

type Gear struct {
	row         string
	gearIndex   int
	PartNumber1 int
	PartNumber2 int
}

func CheckRowForGear(row string) *Gear {
	crr := strings.Split(row, "")
	for i, v := range crr {
		if v == "*" {
			return &Gear{
				row:       row,
				gearIndex: i,
			}
		}
	}

	return &Gear{}
}

func IsAdjacentToASymbol(rows []string) []int {
	partNumbers := []int{}
	currentPartNumber := -1

	for i := 0; i < len(rows); i++ {
		cr := strings.Split(rows[i], "")
		for rowIndex, c := range cr {
			parsed, parseErr := strconv.Atoi(c)
			if currentPartNumber > 0 && parseErr != nil || (rowIndex == len(cr)-1 && currentPartNumber != -1 && parseErr == nil) {
				if rowIndex == len(cr)-1 && parseErr == nil {
					// if we are at the end of the row and the last character is a number
					// add it and check it !
					currentPartNumber = currentPartNumber*10 + parsed
				}

				start, end := getStartAndEndIndex(rows[i], strconv.Itoa(currentPartNumber), rowIndex)
				hit := CheckRow(rows[i][start:end])

				if hit {
					partNumbers = append(partNumbers, currentPartNumber)
					currentPartNumber = -1
					continue
				}

				// Check row before if exists
				if i > 0 {
					hit := CheckRow(rows[i-1][start:end])

					if hit {
						partNumbers = append(partNumbers, currentPartNumber)
						currentPartNumber = -1
						continue
					}
				}

				// check row after if exists
				if i < len(rows)-1 {
					hit := CheckRow(rows[i+1][start:end])

					if hit {
						partNumbers = append(partNumbers, currentPartNumber)
						currentPartNumber = -1
						continue
					}
				}

				currentPartNumber = -1
				continue
			}

			if currentPartNumber == -1 {
				currentPartNumber = parsed
				continue
			}

			currentPartNumber = currentPartNumber*10 + parsed
		}
	}

	return partNumbers
}

func getStartAndEndIndex(haystack string, number string, index int) (int, int) {
	startIndex := index - len(number)
	endIndex := index

	if startIndex > 0 {
		startIndex = startIndex - 1
	}

	if endIndex < len(haystack) {
		endIndex = endIndex + 1
	}

	return startIndex, endIndex
}
