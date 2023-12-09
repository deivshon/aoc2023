package day7

import (
	_ "embed"
	"fmt"
	"slices"
)

//go:embed day7.txt
var input string

func SolveFirst() (string, error) {
	result, err := solve(PartOne)
	return fmt.Sprint(result), err
}

func SolveSecond() (string, error) {
	result, err := solve(PartTwo)
	return fmt.Sprint(result), err
}

func solve(part SolutionPart) (string, error) {
	hands, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("could not parse input: %v", err)
	}

	slices.SortFunc(hands, func(a Hand, b Hand) int {
		result, err := a.WinsAgainst(b, part)
		if err != nil {
			panic(fmt.Sprintf("could not determine victory for hand pair: %v", err))
		}

		return result
	})

	totalWinnings := 0
	for i, h := range hands {
		totalWinnings += (i + 1) * h.Bid
	}

	return fmt.Sprint(totalWinnings), nil
}
