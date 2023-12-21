package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Day7/testInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []string{}

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	TypeOfHands(inputs)

}

type Hand struct {
	Cards string
	Rank  int
	Bid   int
}

func TypeOfHands(rows []string) map[string][]Hand {
	result := map[string][]Hand{}
	for _, row := range rows {
		split := strings.Split(row, " ")

		parseBid, _ := strconv.Atoi(split[1])
		hand := Hand{
			Cards: split[0],
			Rank:  0,
			Bid:   parseBid,
		}

		// Check for pairs
		currentCardIndex := 0
		matches := make(map[string]string)

		for currentCardIndex < len(hand.Cards) {
			currentCard := hand.Cards[currentCardIndex]
			remainingCards := hand.Cards[currentCardIndex+1:]
			for _, card := range remainingCards {
				if string(currentCard) == string(card) {
					if _, ok := matches[string(currentCard)]; !ok {
						matches[string(currentCard)] = string(currentCard)
						continue
					}
					matches[string(currentCard)] += string(currentCard)
					continue
				}
			}
			currentCardIndex++
		}

		if len(matches) == 2 {
			isFullHouse := false
			for a := range matches {
				if len(matches[a]) == 3 {
					isFullHouse = true
					break
				}
			}

			if isFullHouse {
				result["fullHouse"] = append(result["fullHouse"], Hand{
					Cards: split[0],
					Rank:  0,
					Bid:   parseBid,
				})
				continue
			}

			result["twoPair"] = append(result["twoPair"], Hand{
				Cards: split[0],
				Rank:  0,
				Bid:   parseBid,
			})
		} else if len(matches) == 1 {
			log.Printf("Matches: %v - %d", matches, len(matches))
			for a := range matches {
				log.Printf("Matches: %v - %d", matches[a], len(matches[a]))
				if len(matches[a]) == 5 {
					result["fiveOfAKind"] = append(result["fiveOfAKind"], Hand{
						Cards: split[0],
						Rank:  0,
						Bid:   parseBid,
					})
					continue
				}
				if len(matches[a]) == 4 {
					result["fourOfAKind"] = append(result["fourOfAKind"], Hand{
						Cards: split[0],
						Rank:  0,
						Bid:   parseBid,
					})
					continue
				} else if len(matches[a]) == 3 {
					log.Printf("Three of a kind: %v", matches[a])
					result["threeOfAKind"] = append(result["threeOfAKind"], Hand{
						Cards: split[0],
						Rank:  0,
						Bid:   parseBid,
					})
					continue
				} else if len(matches[a]) == 1 {
					result["pair"] = append(result["pair"], Hand{
						Cards: split[0],
						Rank:  0,
						Bid:   parseBid,
					})
					continue
				}
			}
		} else {
			result["highCard"] = append(result["highCard"], Hand{
				Cards: split[0],
				Rank:  0,
				Bid:   parseBid,
			})
		}
	}

	log.Printf("Result: %v", result)

	return result
}
