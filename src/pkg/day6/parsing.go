package day6

import (
	"fmt"
	"main/src/pkg/utils"
	"regexp"
	"strconv"
	"strings"
)

const timesLineIdx = 0
const distancesLineIdx = 1

func parseInputFirst(input string) ([]Race, error) {
	lines := utils.RemoveEmptyStrings(strings.Split(input, "\n"))
	if len(lines) != 2 {
		return nil, fmt.Errorf("expected two lines on input")
	}

	numbersRegexp := regexp.MustCompile(`(\d+)`)

	timesMatches := numbersRegexp.FindAllString(lines[timesLineIdx], -1)
	distanceMatches := numbersRegexp.FindAllString(lines[distancesLineIdx], -1)

	if len(timesMatches) != len(distanceMatches) {
		return nil, fmt.Errorf("expected same amount of times and distances, got `%v` times and `%v` distances", len(timesMatches), len(distanceMatches))
	}

	racesAmount := len(timesMatches)
	races := make([]Race, 0, racesAmount)
	for i := 0; i < racesAmount; i++ {
		rawTime := timesMatches[i]
		rawDistance := distanceMatches[i]

		time, err := strconv.Atoi(rawTime)
		if err != nil {
			return nil, fmt.Errorf("could not parse time `%v`: %v", rawTime, err)
		}

		distance, err := strconv.Atoi(rawDistance)
		if err != nil {
			return nil, fmt.Errorf("could not parse distance `%v`: %v", rawDistance, err)
		}

		races = append(races, Race{
			Duration: time,
			Record:   distance,
		})
	}

	return races, nil
}

func parseInputSecond(input string) (Race, error) {
	lines := utils.RemoveEmptyStrings(strings.Split(input, "\n"))
	if len(lines) != 2 {
		return Race{}, fmt.Errorf("expected two lines on input")
	}

	numbersRegex := regexp.MustCompile(`(\d+)`)

	lines[timesLineIdx] = strings.Join(strings.Split(lines[timesLineIdx], " "), "")
	lines[distancesLineIdx] = strings.Join(strings.Split(lines[distancesLineIdx], " "), "")

	timeMatch := numbersRegex.FindAllString(lines[timesLineIdx], -1)
	distanceMatch := numbersRegex.FindAllString(lines[distancesLineIdx], -1)

	if len(timeMatch) != 1 || len(distanceMatch) != 1 {
		return Race{}, fmt.Errorf("expected same one time and distance, got `%v` times and `%v` distances", len(timeMatch), len(distanceMatch))
	}

	time, err := strconv.Atoi(timeMatch[0])
	if err != nil {
		return Race{}, fmt.Errorf("could not parse time `%v`: %v", timeMatch[0], err)
	}

	distance, err := strconv.Atoi(distanceMatch[0])
	if err != nil {
		return Race{}, fmt.Errorf("could not parse distance `%v`: %v", distanceMatch[0], err)
	}

	return Race{
		Duration: time,
		Record:   distance,
	}, nil
}
