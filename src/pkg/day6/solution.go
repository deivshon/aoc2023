package day6

import (
	_ "embed"
	"fmt"
)

//go:embed day6.txt
var input string

func SolveFirst() (string, error) {
	races, err := parseInputFirst(input)
	if err != nil {
		return "", fmt.Errorf("could not parse input: %v", err)
	}

	result := 1
	for _, r := range races {
		result *= getWaysAmount(r)
	}

	return fmt.Sprint(result), nil
}

func SolveSecond() (string, error) {
	race, err := parseInputSecond(input)
	if err != nil {
		return "", fmt.Errorf("could not parse input: %v", err)
	}

	ways := getWaysAmount(race)
	return fmt.Sprint(ways), nil
}

func getWaysAmount(race Race) int {
	bottomThreshold := 0
	topThreshold := race.Duration

	for distanceTraveled(race, bottomThreshold) <= race.Record {
		bottomThreshold++
	}

	for distanceTraveled(race, topThreshold) <= race.Record {
		topThreshold--
	}

	return topThreshold - bottomThreshold + 1
}

func distanceTraveled(race Race, buttonPressedTime int) int {
	return buttonPressedTime * (race.Duration - buttonPressedTime)
}
