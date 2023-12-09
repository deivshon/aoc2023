package day7

import (
	"fmt"
	"regexp"
	"strconv"
)

const cardsIdx = 1
const bidIdx = 2

func parseInput(input string) ([]Hand, error) {
	handsRegex := regexp.MustCompile(`([\p{Lu}|\d]+) (\d+)`)

	handMatches := handsRegex.FindAllStringSubmatch(input, -1)
	hands := make([]Hand, 0, len(handMatches))
	for _, match := range handMatches {
		if len(match) != 3 {
			return nil, fmt.Errorf("got unexpected match length on match %v", match)
		}

		bid, err := strconv.Atoi(match[bidIdx])
		if err != nil {
			return nil, fmt.Errorf("could not convert bid `%v` to integer: %v", bid, err)
		}

		cards := match[cardsIdx]
		if len(cards) != cardsPerHand {
			return nil, fmt.Errorf("got unexpected cards length: %v", len(cards))
		}

		hand := Hand{
			Bid: bid,
		}

		for i := 0; i < cardsPerHand; i++ {
			hand.Cards[i] = cards[i]
		}

		hands = append(hands, hand)
	}

	return hands, nil
}
