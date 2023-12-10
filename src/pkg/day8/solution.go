package day8

import (
	_ "embed"
	"fmt"
)

//go:embed day8.txt
var input string

var (
	startingWaypoint = [waypointLength]byte{'A', 'A', 'A'}
	endingWaypoint   = [waypointLength]byte{'Z', 'Z', 'Z'}
)

func SolveFirst() (string, error) {
	desertTravel, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("could not parse input: %v", err)
	}

	currentWaypointLookup, exists := desertTravel.Map[startingWaypoint]
	if !exists {
		return "", fmt.Errorf("desert map does not contain expected starting waypoint")
	}

	steps, err := getStepsNumber(desertTravel, currentWaypointLookup, func(waypoint [waypointLength]byte) bool {
		return waypoint == endingWaypoint
	})
	if err != nil {
		return "", fmt.Errorf("could not get steps: %v", err)
	}

	return fmt.Sprint(steps), nil
}

func SolveSecond() (string, error) {
	desertTravel, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("could not parse input: %v", err)
	}

	startingPoints := []DesertLookup{}
	for s, w := range desertTravel.Map {
		if s[len(s)-1] == 'A' {
			startingPoints = append(startingPoints, w)
		}
	}

	steps := make([]int, 0, len(startingPoints))
	for _, s := range startingPoints {
		currentSteps, err := getStepsNumber(desertTravel, s, func(waypoint [waypointLength]byte) bool {
			return waypoint[len(waypoint)-1] == 'Z'
		})
		if err != nil {
			return "", fmt.Errorf("could not determine step amount for a starting waypoint: %v", err)
		}

		steps = append(steps, currentSteps)
	}

	if len(steps) == 1 {
		return fmt.Sprint(steps), nil
	} else if len(steps) != 0 {
		result := leastCommonMultiple(steps[0], steps[1], steps[2:]...)
		return fmt.Sprint(result), nil
	}

	return "", fmt.Errorf("got no starting points")
}

func getStepsNumber(travel DesertTravel, startingWaypointLookup DesertLookup, isFinalWaypoint func(waypoint [waypointLength]byte) bool) (int, error) {
	steps := 0
	endReached := false
	currentWaypointLookup := startingWaypointLookup
	for !endReached {
		for _, d := range travel.Path {
			steps++

			var nextWaypoint [waypointLength]byte
			if d == DirectionLeft {
				nextWaypoint = currentWaypointLookup.Left
			} else if d == DirectionRight {
				nextWaypoint = currentWaypointLookup.Right
			} else {
				return 0, fmt.Errorf("got unknown direction `%v`", d)
			}

			if isFinalWaypoint(nextWaypoint) {
				endReached = true
				break
			}

			nextWaypointLookup, exists := travel.Map[nextWaypoint]
			if !exists {
				return 0, fmt.Errorf("got into waypoint `%v` which does not have map information", nextWaypoint)
			}

			currentWaypointLookup = nextWaypointLookup
		}
	}

	return steps, nil
}

func greatestCommonDivisor(a int, b int) int {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}

	return a
}

func leastCommonMultiple(a int, b int, integers ...int) int {
	result := a * b / greatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = leastCommonMultiple(result, integers[i])
	}

	return result
}
