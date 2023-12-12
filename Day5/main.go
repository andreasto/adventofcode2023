package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func main() {
	file, err := os.Open("./Day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []string{}

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	var locations []int
	seeds, almanacs := ParseMap(inputs)
	for _, v := range seeds {
		locations = append(locations, GetLocation(v, almanacs))
	}

	log.Printf("Day 5 Part 1 Result: %v", slices.Min(locations))

	// Part 2
	locations = []int{}

	chunks, _ := SplitSliceInChunks(seeds, 2)
	channel := make(chan int, len(seeds))
	var wg sync.WaitGroup
	for _, c := range chunks {
		wg.Add(1)
		go EvaluateChunk(c[0], c[1], almanacs, channel, &wg)
	}
	wg.Wait()
	close(channel)

	for i := range channel {
		locations = append(locations, i)
	}

	log.Printf("Day 5 Part 2 Result: %v", slices.Min(locations))
}

func EvaluateChunk(start, length int, almanac Almanac, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var locations []int
	for i := start; i < start+length; i++ {
		if location := GetLocation(i, almanac); location > 0 {
			locations = append(locations, GetLocation(i, almanac))
		}
	}
	ch <- slices.Min(locations)
}

func SplitSliceInChunks(a []int, chuckSize int) ([][]int, error) {
	if chuckSize < 1 {
		return nil, errors.New("chuckSize must be greater that zero")
	}
	chunks := make([][]int, 0, (len(a)+chuckSize-1)/chuckSize)

	for chuckSize < len(a) {
		a, chunks = a[chuckSize:], append(chunks, a[0:chuckSize:chuckSize])
	}
	chunks = append(chunks, a)
	return chunks, nil
}

func GetNext(value int, almanac Almanac, destination string) int {
	if v, ok := almanac[destination]; ok {
		for _, v := range v {
			if v.source <= value && value <= v.source+v.rangeLenght {
				return v.destination + (value - v.source)
			}
		}
	}
	return value
}

func GetLocation(seed int, almanac Almanac) int {
	if soil := GetNext(seed, almanac, "seed-to-soil"); soil >= 0 {
		if fertilizer := GetNext(soil, almanac, "soil-to-fertilizer"); fertilizer >= 0 {
			if water := GetNext(fertilizer, almanac, "fertilizer-to-water"); water >= 0 {
				if light := GetNext(water, almanac, "water-to-light"); light >= 0 {
					if temperature := GetNext(light, almanac, "light-to-temperature"); temperature >= 0 {
						if humidity := GetNext(temperature, almanac, "temperature-to-humidity"); humidity >= 0 {
							if location := GetNext(humidity, almanac, "humidity-to-location"); location >= 0 {
								return location
							}
						}
					}
				}
			}
		}
	}
	return -1
}

type mapInfo struct {
	destination int
	source      int
	rangeLenght int
}

type Almanac map[string][]mapInfo

func ParseMap(rows []string) (seeds []int, almanac Almanac) {
	almanac = Almanac{}
	currentMapName := ""
	for i := 0; i < len(rows); i++ {
		// Get the seeds on first line
		if i == 0 {
			cleaned := strings.TrimSpace(rows[i][strings.Index(rows[i], ":")+1:])
			split := strings.Split(cleaned, " ")

			for _, v := range split {
				seed, _ := strconv.Atoi(v)
				seeds = append(seeds, seed)
			}
			continue
		}

		// Skip empty rows
		if rows[i] == "" {
			continue
		}

		// Get a new map
		if strings.Contains(rows[i], "map:") {
			currentMapName = rows[i][0 : strings.Index(rows[i], ":")-4]
			continue
		}

		// handle numbers'
		destRangeStart, _ := strconv.Atoi(strings.Split(rows[i], " ")[0])
		sourceRangestart, _ := strconv.Atoi(strings.Split(rows[i], " ")[1])
		rangeLength, _ := strconv.Atoi(strings.Split(rows[i], " ")[2])

		if currentMapName == "seed-to-soil" {
			almanac[currentMapName] = append(almanac[currentMapName], mapInfo{destRangeStart, sourceRangestart, rangeLength})
		} else if currentMapName == "soil-to-fertilizer" {
			almanac[currentMapName] = append(almanac[currentMapName], mapInfo{destRangeStart, sourceRangestart, rangeLength})
		} else if currentMapName == "fertilizer-to-water" {
			almanac[currentMapName] = append(almanac[currentMapName], mapInfo{destRangeStart, sourceRangestart, rangeLength})
		} else if currentMapName == "water-to-light" {
			almanac[currentMapName] = append(almanac[currentMapName], mapInfo{destRangeStart, sourceRangestart, rangeLength})
		} else if currentMapName == "light-to-temperature" {
			almanac[currentMapName] = append(almanac[currentMapName], mapInfo{destRangeStart, sourceRangestart, rangeLength})
		} else if currentMapName == "temperature-to-humidity" {
			almanac[currentMapName] = append(almanac[currentMapName], mapInfo{destRangeStart, sourceRangestart, rangeLength})
		} else {
			almanac[currentMapName] = append(almanac[currentMapName], mapInfo{destRangeStart, sourceRangestart, rangeLength})
		}
	}

	return seeds, almanac
}
