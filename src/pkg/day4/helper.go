package day4

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseNumbersString(numbers string) ([]int, error) {
	spaceSplit := strings.Split(numbers, " ")
	parsedNumbers := make([]int, 0, len(spaceSplit))
	for _, n := range spaceSplit {
		if n == "" {
			continue
		}

		parsedN, err := strconv.Atoi(n)
		if err != nil {
			return nil, fmt.Errorf("atoi errored on: `%v`", n)
		}

		parsedNumbers = append(parsedNumbers, parsedN)
	}

	return parsedNumbers, nil
}

func parseCard(card string) (Card, error) {
	cardIdSplit := strings.Split(card, ": ")
	if len(cardIdSplit) != 2 {
		return Card{}, fmt.Errorf("could not split card id and numbers")
	}

	rawNumbers := cardIdSplit[1]
	numbersSplit := strings.Split(rawNumbers, " | ")
	if len(numbersSplit) != 2 {
		return Card{}, fmt.Errorf("could not split winning and owned numbers")
	}

	winningNumbers, err := parseNumbersString(numbersSplit[0])
	if err != nil {
		return Card{}, fmt.Errorf("could not parse winning numbers: %v", err)
	}
	ownedNumbers, err := parseNumbersString(numbersSplit[1])
	if err != nil {
		return Card{}, fmt.Errorf("could not parse owned numbers: %v", err)
	}

	return Card{
		Winning: winningNumbers,
		Owned:   ownedNumbers,
	}, nil

}

func (c *Card) ComputeValue() int {
	value := 0
	for _, n := range c.Owned {
		if slices.Contains(c.Winning, n) {
			if value == 0 {
				value = 1
			} else {
				value *= 2
			}
		}
	}

	return value
}
