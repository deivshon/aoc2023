package day5

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"
)

//go:embed day5.txt
var input string
var sections []string = strings.Split(input, "\n\n")

const seedsPerGoroutine = 50_000_000

func SolveFirst() (string, error) {
	almanac, err := parseInput(sections)
	if err != nil {
		return "", fmt.Errorf("could not parse almanac")
	}

	var lowestLocation *int = nil
	for _, seed := range almanac.Seeds {
		mappedValue := seed
		for _, m := range almanac.Maps {
			mappedValue = getCorrespondingMapValue(mappedValue, m)
		}

		if lowestLocation == nil || *lowestLocation > mappedValue {
			lowestLocation = &mappedValue
		}
	}

	if lowestLocation == nil {
		return "", fmt.Errorf("no location values were found")
	}

	return fmt.Sprint(*lowestLocation), nil
}

func SolveSecond() (string, error) {
	almanac, err := parseInput(sections)
	if err != nil {
		return "", fmt.Errorf("could not parse almanac")
	}

	if len(almanac.Seeds)%2 != 0 {
		return "", fmt.Errorf("last range length missing in seeds line")
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	var lowestLocation *int = nil
	for i := 0; i < len(almanac.Seeds); i += 2 {
		seedsStart := almanac.Seeds[i]
		seedsLength := almanac.Seeds[i+1]

		for rangeStart := seedsStart; rangeStart < seedsStart+seedsLength; rangeStart += seedsPerGoroutine {
			var rangeLength int
			if rangeStart+seedsPerGoroutine < seedsStart+seedsLength {
				rangeLength = seedsPerGoroutine
			} else {
				rangeLength = (seedsStart + seedsLength) - rangeStart
			}

			wg.Add(1)
			go func(start int, length int) {
				defer wg.Done()

				for seed := start; seed < start+length; seed++ {
					mappedValue := seed
					for _, m := range almanac.Maps {
						mappedValue = getCorrespondingMapValue(mappedValue, m)
					}

					mu.Lock()
					if lowestLocation == nil || *lowestLocation > mappedValue {
						lowestLocation = &mappedValue
					}
					mu.Unlock()
				}
			}(rangeStart, rangeLength)
		}
	}
	wg.Wait()

	if lowestLocation == nil {
		return "", fmt.Errorf("no location values were found")
	}

	return fmt.Sprint(*lowestLocation), nil
}

func getCorrespondingMapValue(value int, relevantMap AlmanacMap) int {
	for _, m := range relevantMap {
		if value >= m.SourceStart && value < m.SourceStart+m.RangeLength {
			return value + (m.DestinationStart - m.SourceStart)
		}
	}

	return value
}
