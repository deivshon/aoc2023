package day2

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed day2.txt
var input string
var lines []string = strings.Split(input, "\n")

const partOneRedLimit = 12
const partOneGreenLimit = 13
const partOneBlueLimit = 14

var regexGameId = regexp.MustCompile(`Game (\d+):`)
var regexColors = regexp.MustCompile(`(\d+ \w+|;)`)

func parsePair(pair string) (int, Color, error) {
	splitted := strings.Split(pair, " ")
	if len(splitted) != 2 {
		return 0, "", fmt.Errorf("incorrect split")
	}

	rawAmount, colorType := splitted[0], splitted[1]
	amount, err := strconv.Atoi(rawAmount)
	if err != nil {
		return 0, "", fmt.Errorf("could not convert amount in pair")
	}

	var color Color
	switch colorType {
	case string(ColorRed):
		color = ColorRed
	case string(ColorGreen):
		color = ColorGreen
	case string(ColorBlue):
		color = ColorBlue
	default:
		return 0, "", fmt.Errorf("unknown color in pair")
	}

	return amount, color, nil
}

func parseGameId(line string) (int, error) {
	var gameId int
	gameIdString := regexGameId.FindStringSubmatch(line)
	if len(gameIdString) != 2 {
		return 0, fmt.Errorf("regex failed")
	}
	gameId, err := strconv.Atoi(gameIdString[1])
	if err != nil {
		return 0, fmt.Errorf("conversion failed")
	}

	return gameId, nil
}

func SolveFirst() (string, error) {
	sum := 0
	for _, l := range lines {
		gameId, err := parseGameId(l)
		if err != nil {
			return "", fmt.Errorf("could not determine game id for line `%v`", l)
		}

		redAmount := 0
		greenAmount := 0
		blueAmount := 0

		roundStrings := append(regexColors.FindAllString(l, -1), ";")
		gameOverflowed := false
		for _, match := range roundStrings {
			if match == ";" {
				if redAmount > partOneRedLimit || greenAmount > partOneGreenLimit || blueAmount > partOneBlueLimit {
					gameOverflowed = true
					break
				}
				redAmount, greenAmount, blueAmount = 0, 0, 0
				continue
			}

			amount, color, err := parsePair(match)
			if err != nil {
				return "", fmt.Errorf("could not parse pair `%v`: %v", match, err)
			}
			switch color {
			case ColorRed:
				redAmount += amount
			case ColorGreen:
				greenAmount += amount
			case ColorBlue:
				blueAmount += amount
			}
		}

		if !gameOverflowed {
			sum += gameId
		}
	}

	return fmt.Sprint(sum), nil
}

func SolveSecond() (string, error) {
	sum := 0
	for _, l := range lines {
		redGameMax := 0
		greenGameMax := 0
		blueGameMax := 0

		roundStrings := append(regexColors.FindAllString(l, -1), ";")
		for _, match := range roundStrings {
			if match == ";" {
				continue
			}

			amount, color, err := parsePair(match)
			if err != nil {
				return "", fmt.Errorf("could not parse pair `%v`: %v", match, err)
			}

			switch color {
			case ColorRed:
				if amount > redGameMax {
					redGameMax = amount
				}
			case ColorGreen:
				if amount > greenGameMax {
					greenGameMax = amount
				}
			case ColorBlue:
				if amount > blueGameMax {
					blueGameMax = amount
				}
			}
		}

		sum += redGameMax * greenGameMax * blueGameMax
	}

	return fmt.Sprint(sum), nil
}
