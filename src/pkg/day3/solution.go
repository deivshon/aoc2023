package day3

import (
	_ "embed"
	"fmt"
	"main/src/pkg/utils"
	"strconv"
	"strings"
)

//go:embed day3.txt
var input string
var grid []string = utils.RemoveEmptyStrings(strings.Split(input, "\n"))

func SolveFirst() (string, error) {
	sum := 0
	for lineIdx := 0; lineIdx < len(grid); lineIdx++ {
		for digitIdx := 0; digitIdx < len(grid[lineIdx]); digitIdx++ {
			currentChar := grid[lineIdx][digitIdx]
			if currentChar == '.' || isSymbol(currentChar) {
				continue
			}

			startDigitIdx := digitIdx
			endDigitIdx := getEndDigitIndex(grid, lineIdx, startDigitIdx)

			isPartNumber := false
			for idx := startDigitIdx; idx <= endDigitIdx; idx++ {
				if isCloseToSymbol(grid, lineIdx, idx) {
					isPartNumber = true
					break
				}
			}

			if isPartNumber {
				rawNumber := grid[lineIdx][startDigitIdx : endDigitIdx+1]
				number, err := strconv.Atoi(rawNumber)
				if err != nil {
					return "", fmt.Errorf("could not parse raw number `%v` of line `%v`", rawNumber, grid[lineIdx])
				}

				sum += number
			}

			digitIdx = endDigitIdx
		}
	}

	return fmt.Sprint(sum), nil
}

func SolveSecond() (string, error) {
	sum := 0
	for lineIdx := 0; lineIdx < len(grid); lineIdx++ {
		for digitIdx := 0; digitIdx < len(grid[lineIdx]); digitIdx++ {
			currentChar := grid[lineIdx][digitIdx]
			if currentChar != '*' {
				continue
			}

			lateralNums, err := getLateralAdjacents(grid, lineIdx, digitIdx)
			if err != nil {
				return "", fmt.Errorf("could not get lateral adjacent numbers: %v", err)
			}

			verticalNums, err := getVerticalAdjacents(grid, lineIdx, digitIdx)
			if err != nil {
				return "", fmt.Errorf("could not get vertical adjacent numbers: %v", err)
			}

			adjacents := append(lateralNums, verticalNums...)
			if len(adjacents) == 2 {
				sum += adjacents[0] * adjacents[1]
			}
		}
	}

	return fmt.Sprint(sum), nil
}
