package day8

import (
	"fmt"
	"regexp"
	"strings"
)

const pathSection = 0
const waypointsSection = 1

const sourceWaypointIdx = 1
const leftWaypointIdx = 2
const rightWaypointIdx = 3

func parseInput(input string) (DesertTravel, error) {
	waypointsRegex := regexp.MustCompile(fmt.Sprintf(`(\p{Lu}{%[1]v}) = \((\p{Lu}{%[1]v}), (\p{Lu}{%[1]v})\)`, waypointLength))

	inputSections := strings.Split(input, "\n\n")
	if len(inputSections) != 2 {
		return DesertTravel{}, fmt.Errorf("got wrong number of input sections: `%v`, expected 2", len(inputSections))
	}

	travelPath := strings.TrimSpace(inputSections[pathSection])

	waypointMatches := waypointsRegex.FindAllStringSubmatch(inputSections[waypointsSection], -1)
	desertMap := make(map[[waypointLength]byte]DesertLookup, len(waypointMatches))
	for _, w := range waypointMatches {
		if len(w) != 4 {
			return DesertTravel{}, fmt.Errorf("got wrong number of capture groups in waypoints line: `%v`, expected %v", len(w)-1, waypointLength)
		}

		sourceWaypoint, err := parseWaypointCode(w[sourceWaypointIdx])
		if err != nil {
			return DesertTravel{}, fmt.Errorf("could not parse source waypoint code `%v`: %v", w[sourceWaypointIdx], err)
		}

		leftWaypoint, err := parseWaypointCode(w[leftWaypointIdx])
		if err != nil {
			return DesertTravel{}, fmt.Errorf("could not parse left waypoint code `%v`: %v", w[leftWaypointIdx], err)
		}

		rightWaypoint, err := parseWaypointCode(w[rightWaypointIdx])
		if err != nil {
			return DesertTravel{}, fmt.Errorf("could not parse right waypoint code `%v`: %v", w[rightWaypointIdx], err)
		}

		_, exists := desertMap[sourceWaypoint]
		if exists {
			return DesertTravel{}, fmt.Errorf("got duplicate source waypoint `%v`", sourceWaypoint)
		}

		desertMap[sourceWaypoint] = DesertLookup{
			Left:  leftWaypoint,
			Right: rightWaypoint,
		}

	}
	return DesertTravel{
		Path: travelPath,
		Map:  desertMap,
	}, nil
}

func parseWaypointCode(code string) ([waypointLength]byte, error) {
	parsedWaypoint := [waypointLength]byte{0, 0, 0}
	if len(code) != waypointLength {
		return parsedWaypoint, fmt.Errorf("can't parse waypoint code of length: `%v`, expected %v", len(code), waypointLength)
	}

	for i := 0; i < waypointLength; i++ {
		parsedWaypoint[i] = code[i]
	}

	return parsedWaypoint, nil
}
