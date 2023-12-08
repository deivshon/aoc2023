package day5

import (
	"fmt"
	"regexp"
	"strconv"
)

const seedsListSection = 0
const mapsSectionsStart = 1

func parseInput(inputSections []string) (Almanac, error) {
	if len(inputSections) != 8 {
		return Almanac{}, fmt.Errorf("expected 8 input sections in almanac parsing")
	}

	seedsRegex := regexp.MustCompile(`(\d+)`)
	almanacMapRegex := regexp.MustCompile(`(\d+) (\d+) (\d+)`)

	rawSeeds := seedsRegex.FindAllString(inputSections[seedsListSection], -1)
	seeds := make([]int, 0, len(rawSeeds))
	for _, s := range rawSeeds {
		parsedSeed, err := strconv.Atoi(s)
		if err != nil {
			return Almanac{}, fmt.Errorf("could not parse seed: `%v`: %v", s, err)
		}

		seeds = append(seeds, parsedSeed)
	}

	almanacMaps := make([]AlmanacMap, 0, len(inputSections)-1)
	for _, mapSection := range inputSections[mapsSectionsStart:] {
		matches := almanacMapRegex.FindAllStringSubmatch(mapSection, -1)

		currentAlmanacMap := make(AlmanacMap, 0, len(matches))
		for _, match := range matches {
			if len(match) != 4 {
				return Almanac{}, fmt.Errorf("expected three matches on all almanac lines but got: `%v`", match)
			}

			triplet := make([]int, 0, 3)
			for _, rawNumber := range match[mapsSectionsStart:] {
				num, err := strconv.Atoi(rawNumber)
				if err != nil {
					return Almanac{}, fmt.Errorf("could not convert number `%v`: %v", rawNumber, err)
				}

				triplet = append(triplet, num)
			}

			currentAlmanacMap = append(currentAlmanacMap, AlmanacLine{
				SourceStart:      triplet[1],
				DestinationStart: triplet[0],
				RangeLength:      triplet[2],
			})
		}

		almanacMaps = append(almanacMaps, currentAlmanacMap)
	}

	return Almanac{
		Seeds: seeds,
		Maps:  almanacMaps,
	}, nil
}
