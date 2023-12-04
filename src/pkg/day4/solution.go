package day4

import (
	_ "embed"
	"fmt"
	"main/src/pkg/utils"
	"math"
	"strings"
)

//go:embed day4.txt
var input string
var lines []string = utils.RemoveEmptyStrings(strings.Split(input, "\n"))

func SolveFirst() (string, error) {
	pileValue := 0
	for _, l := range lines {
		card, err := parseCard(l)
		if err != nil {
			return "", fmt.Errorf("could not parse card: %v", err)
		}

		pileValue += card.ComputeValue()
	}

	return fmt.Sprint(pileValue), nil
}

func SolveSecond() (string, error) {
	copies := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		card, err := parseCard(lines[i])
		if err != nil {
			return "", fmt.Errorf("could not parse card: %v", err)
		}

		matchingNumbers := 0
		value := card.ComputeValue()
		if value != 0 {
			matchingNumbers = int(math.Floor(math.Log2(float64(card.ComputeValue())))) + 1
		}

		for j := i + 1; j < len(lines) && j <= i+matchingNumbers; j++ {
			copies[j] += copies[i] + 1
		}
	}

	copiesAmount := 0
	for _, amount := range copies {
		copiesAmount += amount + 1
	}

	return fmt.Sprint(copiesAmount), nil
}
