package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []string{}

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	result := 0
	result2 := 0
	for _, v := range inputs {
		result += ParseRow(v)
		result2 += ParseRowPart2(v)
	}

	log.Printf("Day 2 Part 1 Result: %d", result)
	log.Printf("Day 2 Part 2 Result: %d", result2)
}

const intRed = 12
const intGreen = 13
const intBlue = 14

func ParseRow(row string) int {
	gameNumber := row[5:strings.Index(row, ":")]
	parsedGame, _ := strconv.Atoi(gameNumber)

	results := strings.Split(row[strings.Index(row, ":")+1:], ",")
	allPossible := true

	for _, v := range results {
		isPossible := true
		s := strings.Split(v, ";")
		for _, g := range s {
			g = strings.TrimSpace(g)
			split := strings.Split(g, " ")

			color := split[1]
			number, _ := strconv.Atoi(strings.TrimSpace(split[0]))

			if color == "red" {
				isPossible = intRed/number > 0
			} else if color == "green" {
				isPossible = intGreen/number > 0
			} else if color == "blue" {
				log.Printf("BLUE: %d", intBlue/number)
				isPossible = intBlue/number > 0
			}

			if !isPossible {
				allPossible = false
				break
			}
		}
	}

	if allPossible {
		return parsedGame
	}

	return 0
}

func ParseRowPart2(row string) int {
	results := strings.Split(row[strings.Index(row, ":")+1:], ",")
	red := 0
	green := 0
	blue := 0

	for _, v := range results {
		s := strings.Split(v, ";")
		for _, g := range s {
			g = strings.TrimSpace(g)
			split := strings.Split(g, " ")

			color := split[1]
			number, _ := strconv.Atoi(strings.TrimSpace(split[0]))

			if color == "red" {
				if red == 0 {
					red = number
				} else {
					possible := red/number > 0
					if !possible {
						red = number
					}
				}

			} else if color == "green" {
				if green == 0 {
					green = number
				} else {
					possible := green/number > 0
					if !possible {
						green = number
					}
				}
			} else if color == "blue" {
				if blue == 0 {
					blue = number
				} else {
					possible := blue/number > 0
					if !possible {
						blue = number
					}
				}
			}
		}
	}

	return red * green * blue
}
