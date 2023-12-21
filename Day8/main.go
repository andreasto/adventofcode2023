package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./Day8/input.txt")
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

	m := NewDesertMap(inputs)
	// log.Printf("Day 8, Part 1 Steps: %d", m.GetDestinations())
	log.Printf("Day 8, Part 2 Steps: %d", m.GetDestinationsPart2())
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

type DesertMap struct {
	Instructions []string
	Destinations map[string]string
	FinalDest    string
}

func NewDesertMap(rows []string) *DesertMap {
	d := &DesertMap{
		FinalDest: "ZZZ",
	}

	m := make(map[string]string)

	replacer := strings.NewReplacer("(", "", ")", "", ")", "", ",", "")
	for i := 0; i < len(rows); i++ {
		if i == 0 {
			d.Instructions = strings.Split(rows[i], "")
		} else {
			m[rows[i][0:3]] = replacer.Replace(rows[i][6:])
		}
	}

	d.Destinations = m
	return d
}

func (d *DesertMap) GetDestinations() int {
	currentDest := "AAA"
	jumps := 0

	for {
		for i := 0; i < len(d.Instructions); i++ {
			if d.Instructions[i] == "L" {
				currentDest = strings.Split(d.Destinations[currentDest], " ")[0]
			} else {
				currentDest = strings.Split(d.Destinations[currentDest], " ")[1]
			}
			jumps++
		}

		if currentDest == d.FinalDest {
			break
		}
	}

	return jumps
}

func (d *DesertMap) GetDestinationsPart2() int64 {
	positions := []string{}
	for a := range d.Destinations {
		if strings.HasSuffix(a, "A") {
			positions = append(positions, a)
		}
	}

	cycles := [][]int{}

	for i := 0; i < len(positions); i++ {
		cycle := []int{}
		current_steps := d.Instructions
		step_count := 0
		first_z := ""

		for {
			for step_count == 0 || !strings.HasSuffix(positions[i], "Z") {
				step_count += 1

				if current_steps[0] == "L" {
					positions[i] = strings.Split(d.Destinations[positions[i]], " ")[0]
				} else {

					positions[i] = strings.Split(d.Destinations[positions[i]], " ")[1]
				}
				current_steps = append(current_steps[1:], current_steps[0])
			}

			cycle = append(cycle, step_count)
			if first_z == "" {
				first_z = positions[i]
				step_count = 0
			} else if positions[i] == first_z {
				break
			}
		}

		cycles = append(cycles, cycle)
	}

	numbers := []int{}
	for i := 0; i < len(cycles); i++ {
		numbers = append(numbers, cycles[i][0])
	}

	lcm := numbers[0]

	for _, v := range numbers {
		lcm = lcm * v / gcd(lcm, v)
	}

	return int64(lcm)
}
