package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []string{}

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	log.Printf("Day 6 Part1 result: %d", Part1(inputs))
	log.Printf("Day 6 Part2 result: %d", Part2(inputs))
}

type Race struct {
	Time     int
	Distance int
}

func Part2(rows []string) int {
	races := []Race{}
	for i := 0; i < len(rows); i++ {
		space := regexp.MustCompile(`\s+`)
		s := space.ReplaceAllString(rows[i], "")
		split := strings.Split(s, ":")
		log.Printf("Parsed: %s", split)
		for j := 0; j < len(split); j++ {
			parsed, err := strconv.Atoi(strings.TrimSpace(split[j]))
			if err != nil {
				// Row description, skipping index
				continue
			}

			if i == 0 {
				// add only for first row so map is correct
				races = append(races, Race{})
			}

			if strings.Contains(split[0], "Time") {
				races[j-1].Time = parsed
			} else if strings.Contains(split[0], "Distance") {
				races[j-1].Distance = parsed
			}
		}
	}

	return ParseRaces(races)
}

func Part1(rows []string) int {
	races := []Race{}
	for i := 0; i < len(rows); i++ {
		space := regexp.MustCompile(`\s+`)
		s := space.ReplaceAllString(rows[i], " ")

		split := strings.Split(s, " ")
		for j := 0; j < len(split); j++ {
			parsed, err := strconv.Atoi(strings.TrimSpace(split[j]))
			if err != nil {
				// Row description, skipping index
				continue
			}

			if i == 0 {
				// add only for first row so map is correct
				races = append(races, Race{})
			}

			if strings.Contains(split[0], "Time:") {
				races[j-1].Time = parsed
			} else if strings.Contains(split[0], "Distance:") {
				races[j-1].Distance = parsed
			}
		}
	}

	return ParseRaces(races)

}

func ParseRaces(races []Race) int {
	totalResult := []int{}
	for _, r := range races {
		winningRounds := 0

		for i := 0; i < r.Time+1; i++ {
			boatWillTravel := r.Time - i
			boatSpeed := i * 1

			result := boatWillTravel * boatSpeed
			if result > r.Distance {
				winningRounds++
			}

			if i == r.Time {
				totalResult = append(totalResult, winningRounds)
			}
		}
	}

	returnValue := 1
	for _, v := range totalResult {
		returnValue *= v
	}

	return returnValue
}
