package day1

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed day1.txt
var input string
var lines []string = strings.Split(input, "\n")
var zero int = 0

func SolveFirst() (string, error) {
	sum := 0
	var firstDigit *int
	var lastDigit *int

	for _, l := range lines {
		firstDigit = nil
		lastDigit = nil

		for _, c := range l {
			asciiValue := int(c)
			if asciiValue >= 48 && asciiValue <= 57 {
				digitValue := asciiValue - 48

				if firstDigit == nil {
					firstDigit = &digitValue
				}
				lastDigit = &digitValue
			}
		}

		if firstDigit == nil {
			firstDigit = &zero
		}
		if lastDigit == nil {
			lastDigit = &zero
		}

		sum += *firstDigit*10 + *lastDigit
	}

	return fmt.Sprint(sum), nil
}

var digitLookup map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func SolveSecond() (string, error) {
	sum := 0
	regexTwoNos := regexp.MustCompile(`^.*?(\d|one|two|three|four|five|six|seven|eight|nine).*(\d|one|two|three|four|five|six|seven|eight|nine).*$`)
	regexOneNo := regexp.MustCompile(`^.*(\d|one|two|three|four|five|six|seven|eight|nine).*$`)

	for _, l := range lines {
		var firstDigit *string
		var secondDigit *string

		doubleMatch := regexTwoNos.FindStringSubmatch(l)
		if len(doubleMatch) == 3 {
			firstDigit = &doubleMatch[1]
			secondDigit = &doubleMatch[2]
		} else if singleMatch := regexOneNo.FindStringSubmatch(l); len(singleMatch) == 2 {
			firstDigit = &singleMatch[1]
			secondDigit = &singleMatch[1]
		} else {
			return "", fmt.Errorf("line `%v` does not contain numbers", l)
		}

		firstDigitInt, firstExists := digitLookup[*firstDigit]
		secondDigitInt, secondExists := digitLookup[*secondDigit]
		if !firstExists || !secondExists {
			return "", fmt.Errorf("either `%v` or `%v` does not match any known digit", *firstDigit, *secondDigit)
		}

		sum += firstDigitInt*10 + secondDigitInt
	}

	return fmt.Sprint(sum), nil
}
