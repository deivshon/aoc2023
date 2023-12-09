package day7

import (
	"fmt"
)

type CardRank = int

const (
	FiveOfAKind  CardRank = 6
	FourOfAKind  CardRank = 5
	FullHouse    CardRank = 4
	ThreeOfAKind CardRank = 3
	TwoPair      CardRank = 2
	OnePair      CardRank = 1
	HighCard     CardRank = 0
)

type HandComparisonResult = int

const (
	Win  HandComparisonResult = 1
	Draw HandComparisonResult = 0
	Loss HandComparisonResult = -1
)

const cardsPerHand = 5

type Hand struct {
	Cards [cardsPerHand]byte
	Bid   int
}

type SolutionPart = int

const (
	PartOne SolutionPart = 1
	PartTwo SolutionPart = 2
)

const jokerCardIndex = 9
const jokerCard byte = 'J'
const differentCards = 13

func (h Hand) WinsAgainst(opponent Hand, part SolutionPart) (int, error) {
	if part != PartOne && part != PartTwo {
		return 0, fmt.Errorf("unknown solution part passed: `%v`", part)
	}

	var ownRank int
	var opponentRank int
	switch part {
	case PartOne:
		ownPartOneRank, err := PartOneRank(h.Cards)
		if err != nil {
			return 0, fmt.Errorf("could not determine own rank (cards %v): %v", h.Cards, err)
		}

		opponentPartOneRank, err := PartOneRank(opponent.Cards)
		if err != nil {
			return 0, fmt.Errorf("could not determine opponent rank (cards %v): %v", opponent.Cards, err)
		}

		ownRank = ownPartOneRank
		opponentRank = opponentPartOneRank
	case PartTwo:
		ownPartTwoRank, err := h.PartTwoRank()
		if err != nil {
			return 0, fmt.Errorf("could not get own part two rank (cards %v): %v", h.Cards, err)
		}
		opponentPartTwoRank, err := opponent.PartTwoRank()
		if err != nil {
			return 0, fmt.Errorf("could not get opponent part two rank (cards %v): %v", h.Cards, err)
		}

		ownRank = ownPartTwoRank
		opponentRank = opponentPartTwoRank
	}

	if ownRank > opponentRank {
		return Win, nil
	}
	if opponentRank > ownRank {
		return Loss, nil
	}

	playoffResult, err := h.playoffAgainst(opponent, part)
	if err != nil {
		return 0, fmt.Errorf("could not determine playoff result: %v", err)
	}

	return playoffResult, nil
}

func PartOneRank(cards [cardsPerHand]byte) (int, error) {
	cardsMap, err := getCardsMap(cards)
	if err != nil {
		return 0, fmt.Errorf("could not get cards map: %v", err)
	}

	differentCards := 0
	for _, amount := range cardsMap {
		if amount != 0 {
			differentCards++
		}
	}

	switch differentCards {
	case 5:
		return HighCard, nil
	case 4:
		return OnePair, nil
	case 3:
		for _, amount := range cardsMap {
			if amount == 2 {
				return TwoPair, nil
			}
			if amount == 3 {
				return ThreeOfAKind, nil
			}
		}

		return 0, fmt.Errorf("could not find expected amounts in three-different-cards branch")
	case 2:
		for _, amount := range cardsMap {
			if amount == 4 || amount == 1 {
				return FourOfAKind, nil
			}
			if amount == 3 || amount == 2 {
				return FullHouse, nil
			}
		}

		return 0, fmt.Errorf("could not find expected amounts in two-different-cards branch")
	case 1:
		return FiveOfAKind, nil
	default:
		return 0, fmt.Errorf("got unexpected number of different cards: %v", differentCards)
	}
}

func (h Hand) PartTwoRank() (int, error) {
	bestCards, err := getBestCardForRank(h.Cards)
	if err != nil {
		return 0, fmt.Errorf("could not get own best cards for rank (cards %v): %v", h.Cards, err)
	}

	partTwoRank, err := PartOneRank(bestCards)
	if err != nil {
		return 0, fmt.Errorf("could determine own part two rank (cards %v): %v", bestCards, err)
	}

	return partTwoRank, nil
}

func (h Hand) playoffAgainst(opponent Hand, part SolutionPart) (int, error) {
	for i := 0; i < cardsPerHand; i++ {
		ownCurrentCard := h.Cards[i]
		opponentCurrentCard := opponent.Cards[i]

		ownCurrentCardScore, err := getCardPlayoffScore(ownCurrentCard, part)
		if err != nil {
			return 0, fmt.Errorf("could not determine own card score for card `%v`: %v", rune(ownCurrentCard), err)
		}

		opponentCurrentCardScore, err := getCardPlayoffScore(opponentCurrentCard, part)
		if err != nil {
			return 0, fmt.Errorf("could not determine opponent card score for card `%v`: %v", rune(opponentCurrentCard), err)
		}

		if ownCurrentCardScore > opponentCurrentCardScore {
			return Win, nil
		}
		if opponentCurrentCardScore > ownCurrentCardScore {
			return Loss, nil
		}
	}

	return Draw, nil
}

func getCardIndex(card byte) (int, error) {
	switch card {
	case 'A':
		return 12, nil
	case 'K':
		return 11, nil
	case 'Q':
		return 10, nil
	case 'J':
		return jokerCardIndex, nil
	case 'T':
		return 8, nil
	case '9':
		return 7, nil
	case '8':
		return 6, nil
	case '7':
		return 5, nil
	case '6':
		return 4, nil
	case '5':
		return 3, nil
	case '4':
		return 2, nil
	case '3':
		return 1, nil
	case '2':
		return 0, nil
	default:
		return 0, fmt.Errorf("`%v` is not a card", rune(card))
	}
}

func getCardByIndex(index int) (byte, error) {
	switch index {
	case 12:
		return 'A', nil
	case 11:
		return 'K', nil
	case 10:
		return 'Q', nil
	case jokerCardIndex:
		return jokerCard, nil
	case 8:
		return 'T', nil
	case 7:
		return '9', nil
	case 6:
		return '8', nil
	case 5:
		return '7', nil
	case 4:
		return '6', nil
	case 3:
		return '5', nil
	case 2:
		return '4', nil
	case 1:
		return '3', nil
	case 0:
		return '2', nil
	default:
		return 0, fmt.Errorf("`%v` is not a valid card index", index)
	}
}

func getCardPlayoffScore(card byte, part SolutionPart) (int, error) {
	var increase int
	switch part {
	case PartOne:
		increase = 0
	case PartTwo:
		increase = 1
	default:
		return 0, fmt.Errorf("could not accept solution part value `%v`", part)
	}

	switch card {
	case 'A':
		return 12, nil
	case 'K':
		return 11, nil
	case 'Q':
		return 10, nil
	case jokerCard:
		if part == PartOne {
			return 9, nil
		} else {
			return 0, nil
		}
	case 'T':
		return 8 + increase, nil
	case '9':
		return 7 + increase, nil
	case '8':
		return 6 + increase, nil
	case '7':
		return 5 + increase, nil
	case '6':
		return 4 + increase, nil
	case '5':
		return 3 + increase, nil
	case '4':
		return 2 + increase, nil
	case '3':
		return 1 + increase, nil
	case '2':
		return 0 + increase, nil
	default:
		return 0, fmt.Errorf("`%v` is not a card", rune(card))
	}
}

func getCardsMap(cards [cardsPerHand]byte) ([differentCards]int, error) {
	cardsMap := [differentCards]int{}

	for _, c := range cards {
		mappedIndex, err := getCardIndex(c)
		if err != nil {
			return cardsMap, fmt.Errorf("could not recognize card `%v`: %v", c, err)
		}
		if mappedIndex >= len(cardsMap) {
			return cardsMap, fmt.Errorf("got unexpected index `%v`, too high for cards map length (%v)", mappedIndex, len(cardsMap))
		}

		cardsMap[mappedIndex]++
	}

	return cardsMap, nil
}

func getBestCardForRank(cards [cardsPerHand]byte) ([cardsPerHand]byte, error) {
	newCards := [cardsPerHand]byte{}

	cardsMap, err := getCardsMap(cards)
	if err != nil {
		return newCards, fmt.Errorf("could not get cards map: %v", err)
	}
	if cardsMap[jokerCardIndex] == 0 {
		return cards, nil
	}

	var highestCardNumberIndex *int = nil
	for i, amount := range cardsMap {
		if i != jokerCardIndex && (highestCardNumberIndex == nil || amount > cardsMap[*highestCardNumberIndex]) {
			newHighestCardNumberIndex := i
			highestCardNumberIndex = &newHighestCardNumberIndex
		}
	}

	if highestCardNumberIndex == nil {
		return newCards, fmt.Errorf("could not find any existing card amount in cards map")
	}

	mostCommonCard, err := getCardByIndex(*highestCardNumberIndex)
	if err != nil {
		return newCards, fmt.Errorf("could not get card from most common card index `%v`: %v", *highestCardNumberIndex, err)
	}

	for i := 0; i < cardsPerHand; i++ {
		if cards[i] == jokerCard {
			newCards[i] = mostCommonCard
		} else {
			newCards[i] = cards[i]
		}
	}

	return newCards, nil
}
