package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
	result2 := CheckRowsPart2(inputs)

	log.Printf("Day 4 Part 1 Result: %v", result)
	log.Printf("Day 4 Part 2 Result: %d", result2)
}

type Card struct {
	CardNumber             int
	NumberOfWinningNumbers int
	NumberOfCopies         int
}

func CheckRowsPart2(rows []string) int {
	cards := []Card{}

	for i := 0; i < len(rows); i++ {
		cardNumber := rows[i][strings.Index(rows[i], "d")+1 : strings.Index(rows[i], ":")]

		cardNumberParsed, _ := strconv.Atoi(strings.TrimSpace(cardNumber))
		// log.Printf("CardParsed %d", cardNumberParsed)

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
					rowResult++
				}
			}
		}

		cards = append(cards, Card{
			CardNumber:             cardNumberParsed,
			NumberOfWinningNumbers: rowResult,
		})
	}

	var totalCards int
	for i := 0; i < len(cards); i++ {
		if cards[i].NumberOfWinningNumbers == 0 {
			continue
		}

		for j := 0; j < cards[i].NumberOfWinningNumbers; j++ {
			cards[i+1+j].NumberOfCopies++
		}

		for k := 0; k < cards[i].NumberOfCopies; k++ {
			for y := 0; y < cards[i].NumberOfWinningNumbers; y++ {
				cards[i+1+y].NumberOfCopies++
			}
		}
	}

	for x := 0; x < len(cards); x++ {
		totalCards = totalCards + cards[x].NumberOfCopies + 1
	}

	return totalCards
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
