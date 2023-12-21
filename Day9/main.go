package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func main() {
	file, err := os.Open("./Day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []string{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		inputs = append(inputs, scanner.Text())
	}

	log.Printf("Day 9, Part 1: %d", part1(inputs))
	log.Printf("Day 9, Part 2: %d", part2(inputs))
}

func part1(rows []string) int {
	result := 0
	for i := 0; i < len(rows); i++ {
		split, _ := sliceAtoi(strings.Split(rows[i], " "))
		nextRow := []int{}
		completeRows := [][]int{}

		nextRow = split

		completeRows = append(completeRows, nextRow)

		for !AllZeros(nextRow) {
			nextNextRow := []int{}
			for j := 0; j < len(nextRow); j++ {
				if j == len(nextRow)-1 {
					break
				} else {
					nextNextRow = append(nextNextRow, nextRow[j+1]-nextRow[j])
				}
			}
			completeRows = append(completeRows, nextNextRow)
			nextRow = nextNextRow
		}

		slices.Reverse(completeRows)
		numberToAdd := 0

		for i, row := range completeRows {
			if i == len(completeRows)-1 {
				result += row[len(row)-1] + numberToAdd
				numberToAdd = 0
			} else {
				numberToAdd = row[len(row)-1] + numberToAdd
			}
		}
	}

	return result
}

func part2(rows []string) int {
	result := 0
	for i := 0; i < len(rows); i++ {
		split, _ := sliceAtoi(strings.Split(rows[i], " "))
		nextRow := []int{}
		completeRows := [][]int{}

		nextRow = split

		completeRows = append(completeRows, nextRow)

		for !AllZeros(nextRow) {
			nextNextRow := []int{}
			for j := 0; j < len(nextRow); j++ {
				if j == len(nextRow)-1 {
					break
				} else {
					nextNextRow = append(nextNextRow, nextRow[j+1]-nextRow[j])
				}
			}
			completeRows = append(completeRows, nextNextRow)
			nextRow = nextNextRow
		}

		slices.Reverse(completeRows)
		numberToAdd := 0
		for i, row := range completeRows {
			newRow := lo.Reverse(row)
			if i == len(completeRows)-1 {
				result += newRow[len(newRow)-1] - numberToAdd
				numberToAdd = 0
			} else {
				numberToAdd = newRow[len(newRow)-1] - numberToAdd
			}
		}
	}

	return result
}

func AllZeros(row []int) bool {
	for i := 1; i < len(row); i++ {
		if row[0] != 0 || row[i] != row[0] {
			return false
		}
	}
	return true
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
