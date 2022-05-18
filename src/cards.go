package src

import "strings"

type Card struct {
	CardNumber string
	ExpMonth   string
	ExpYear    string
	Cvv        string
}

func GetCardByLine(line, separator string) Card {
	var card Card
	PreFormattedCard := strings.Split(line, separator)
	card.CardNumber, card.ExpMonth, card.ExpYear, card.Cvv = PreFormattedCard[0], PreFormattedCard[1], PreFormattedCard[2], PreFormattedCard[3]
	return card
}
