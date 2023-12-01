package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []string{}

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	var result = 0

	for _, v := range inputs {
		r := Part1(v)

		result += r
	}

	log.Printf("Day 1 Part 1 Result: %d", result)

	result = 0
	for _, v := range inputs {

		r := Part2(v)
		result += r
	}
	log.Printf("Day 1 Part 2 Result: %d", result)
}

func Part1(row string) int {
	result := [2]int{-1, -1}

	split := strings.Split(row, "")

	for _, v := range split {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			continue
		}

		if result[0] == -1 {
			result[0] = parsed
			continue
		}

		result[1] = parsed
	}

	if len(result) == 2 {
		var second = 0
		if result[1] == -1 {
			second = result[0]
		} else {
			second = result[1]
		}

		combined := strconv.Itoa(result[0]) + strconv.Itoa(second)
		parsed, _ := strconv.Atoi(combined)

		return parsed
	}

	return 0
}

type StringValue struct {
	stringValue string
}

func (s StringValue) GetValue() int {
	if strings.Contains(s.stringValue, "one") {
		return 1
	} else if strings.Contains(s.stringValue, "two") {
		return 2
	} else if strings.Contains(s.stringValue, "three") {
		return 3
	} else if strings.Contains(s.stringValue, "four") {
		return 4
	} else if strings.Contains(s.stringValue, "five") {
		return 5
	} else if strings.Contains(s.stringValue, "six") {
		return 6
	} else if strings.Contains(s.stringValue, "seven") {
		return 7
	} else if strings.Contains(s.stringValue, "eight") {
		return 8
	} else if strings.Contains(s.stringValue, "nine") {
		return 9
	}

	return -1
}

func Part2(row string) int {
	result := [2]string{"", ""}

	split := strings.Split(row, "")

	a := StringValue{}
	for _, v := range split {
		_, err := strconv.Atoi(string(v))

		if err != nil {
			a.stringValue += v
			if a.GetValue() != -1 {
				result[0] = strconv.Itoa(a.GetValue())
				break
			}
			continue
		}

		if result[0] == "" {
			result[0] = string(v)
			break
		}
	}

	a = StringValue{}
	reverted := StringValue{}
	for i := len(split) - 1; i >= 0; i-- {
		v := split[i]
		_, err := strconv.Atoi(string(v))

		if err != nil {
			a.stringValue += v
			reverted.stringValue = reverseString(a.stringValue)
			if reverted.GetValue() != -1 {
				result[1] = strconv.Itoa(reverted.GetValue())
				break
			}
			continue
		}

		if result[1] == "" {
			result[1] = string(v)
			break
		}
	}

	if result[1] == "" {
		result[1] = result[0]
	}

	res := result[0] + result[1]
	parsed, err := strconv.Atoi(res)

	if err != nil {
		return 0
	}

	return parsed

}

func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}
