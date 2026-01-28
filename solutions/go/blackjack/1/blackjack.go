package blackjack

import "strings"

var cardMap = map[string]int{
	"one":   0,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"ace":   11,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value, ok := cardMap[strings.ToLower(card)]
	if !ok {
		value = 0
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	result := ""
	card1 = strings.ToLower(card1)
	card2 = strings.ToLower(card2)
	dealerCard = strings.ToLower(dealerCard)

	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	totalCardsValue := card1Value + card2Value

	dealerCardValue := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		result = "P"
	case totalCardsValue == 21:
		if dealerCardValue < 10 {
			result = "W"
		} else {
			result = "S"
		}
	case totalCardsValue >= 17 && totalCardsValue <= 20:
		result = "S"
	case totalCardsValue >= 12 && totalCardsValue <= 16:
		if dealerCardValue >= 7 {
			result = "H"
		} else {
			result = "S"
		}
	case totalCardsValue <= 11:
		result = "H"
	}

	return result
}
