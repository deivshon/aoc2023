package day3

import (
	"fmt"
	"strconv"
	"unicode"
)

type VerticalDirection int

const (
	DirectionUp   VerticalDirection = -1
	DirectionDown VerticalDirection = 1
)

func getSingleVerticalDirectionAdjecents(grid []string, lineIdx int, starIdx int, direction VerticalDirection) ([]int, error) {
	notRightBorder := starIdx != len(grid[lineIdx])-1
	notLeftBorder := starIdx != 0
	analyzedLineIdx := lineIdx + int(direction)

	rawVerticalNums := []int{}
	verticalNums := []int{}
	errs := []error{}

	verticalCentreIsDigit := unicode.IsDigit(rune(grid[analyzedLineIdx][starIdx]))
	verticalRightIsDigit := notRightBorder && unicode.IsDigit(rune(grid[analyzedLineIdx][starIdx+1]))
	verticalLeftIsDigit := notLeftBorder && unicode.IsDigit(rune(grid[analyzedLineIdx][starIdx-1]))
	if verticalCentreIsDigit {
		verticalCentreNumStart, verticalCentreNumEnd := getStartEndDigitIdx(grid, analyzedLineIdx, starIdx)
		verticalCentreNum, err := strconv.Atoi(grid[analyzedLineIdx][verticalCentreNumStart : verticalCentreNumEnd+1])
		errs = append(errs, err)
		rawVerticalNums = append(rawVerticalNums, verticalCentreNum)
	}
	if verticalRightIsDigit {
		verticalRightNumStart, verticalRightNumEnd := getStartEndDigitIdx(grid, analyzedLineIdx, starIdx+1)
		verticalRightNum, err := strconv.Atoi(grid[analyzedLineIdx][verticalRightNumStart : verticalRightNumEnd+1])
		errs = append(errs, err)
		rawVerticalNums = append(rawVerticalNums, verticalRightNum)
	}
	if verticalLeftIsDigit {
		verticalLeftNumStart, verticalLeftNumEnd := getStartEndDigitIdx(grid, analyzedLineIdx, starIdx-1)
		verticalLeftNum, err := strconv.Atoi(grid[analyzedLineIdx][verticalLeftNumStart : verticalLeftNumEnd+1])
		errs = append(errs, err)
		rawVerticalNums = append(rawVerticalNums, verticalLeftNum)
	}

	if len(rawVerticalNums) == 3 || len(rawVerticalNums) == 1 {
		verticalNums = append(verticalNums, rawVerticalNums[0])
	} else if len(rawVerticalNums) == 2 {
		if verticalCentreIsDigit {
			verticalNums = append(verticalNums, rawVerticalNums[0])
		} else {
			verticalNums = append(verticalNums, rawVerticalNums...)
		}
	} else if len(rawVerticalNums) != 0 {
		return nil, fmt.Errorf("got length higher than three on parsed numbers array")
	}

	for _, err := range errs {
		if err != nil {
			return nil, fmt.Errorf("could not parse some number: %v", err)
		}
	}

	return verticalNums, nil
}

func getVerticalAdjacents(grid []string, lineIdx int, starIdx int) ([]int, error) {
	notUpperLine := lineIdx != 0
	notBottomLine := lineIdx != len(grid)-1

	allVerticalNums := []int{}
	if notUpperLine {
		upperNums, err := getSingleVerticalDirectionAdjecents(grid, lineIdx, starIdx, DirectionUp)
		if err != nil {
			return nil, fmt.Errorf("could not parse upper numbers: %v", err)
		}

		allVerticalNums = append(allVerticalNums, upperNums...)
	}
	if notBottomLine {
		bottomNums, err := getSingleVerticalDirectionAdjecents(grid, lineIdx, starIdx, DirectionDown)
		if err != nil {
			return nil, fmt.Errorf("could not parse bottom numbers: %v", err)
		}

		allVerticalNums = append(allVerticalNums, bottomNums...)
	}

	return allVerticalNums, nil
}

func getLateralAdjacents(grid []string, lineIdx int, starIdx int) ([]int, error) {
	notRightBorder := starIdx != len(grid[lineIdx])-1
	notLeftBorder := starIdx != 0

	lateralNums := []int{}
	if notRightBorder && unicode.IsDigit(rune(grid[lineIdx][starIdx+1])) {
		rightNumStart, rightNumEnd := getStartEndDigitIdx(grid, lineIdx, starIdx+1)
		rightNum, err := strconv.Atoi(grid[lineIdx][rightNumStart : rightNumEnd+1])
		if err != nil {
			return nil, fmt.Errorf("could not parse right number: %v", err)
		}

		lateralNums = append(lateralNums, rightNum)
	}
	if notLeftBorder && unicode.IsDigit(rune(grid[lineIdx][starIdx-1])) {
		leftNumStart, leftNumEnd := getStartEndDigitIdx(grid, lineIdx, starIdx-1)
		leftNum, err := strconv.Atoi(grid[lineIdx][leftNumStart : leftNumEnd+1])
		if err != nil {
			return nil, fmt.Errorf("could not parse left number: %v", err)
		}

		lateralNums = append(lateralNums, leftNum)
	}

	return lateralNums, nil
}
