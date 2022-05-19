package src

import "fmt"

type CardError struct{}

func (cr *CardError) Error() string {
	return "your card-list file contains empty lines"
}

// define a card using a line
func DefineCard(line, separator string) (Card, error) {
	card := GetCardByLine(line, separator)

	if line == "" {
		return card, &CardError{}
	}

	return card, nil
}

// start checking process
func CheckProcess(card Card, cfg Cfg, output string) {
	result = CheckCard(card, cfg)

	if result.Valid {
		SaveCard(card, output, result)

		if result.DeclineCodeValid {
			fmt.Printf("[possible live] %s, %s/%s, %s (%s, %s) \n", card.CardNumber, card.ExpMonth, card.ExpYear, card.Cvv, result.Code, result.DeclinedReason)
			return
		}

		fmt.Printf("[live] %s, %s/%s, %s (%s) \n", card.CardNumber, card.ExpMonth, card.ExpYear, card.Cvv, result.Code)

	} else {
		if result.DeclineCodeValid {
			fmt.Printf("[die] %s, %s/%s, %s (%s, %s) \n", card.CardNumber, card.ExpMonth, card.ExpYear, card.Cvv, result.Code, result.DeclinedReason)
			return
		}
		fmt.Printf("[die] %s, %s/%s, %s (%s) \n", card.CardNumber, card.ExpMonth, card.ExpYear, card.Cvv, result.Code)
	}
}
