package src

import (
	"bufio"
	"fmt"
	"os"
)

// handler is a callback
type Handler func(line string)

// open card lists and read optimally to handle large bytes of data from a list, using callback on each line.
func OpenFileByName(filename string, handler Handler) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for {
		if scanner.Scan() {
			line := scanner.Text()
			handler(line)
			continue
		}
		break
	}
	return nil
}

// save live cards in an output.
func SaveCard(cc Card, output string, r Result) error {
	f, err := os.OpenFile(output, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	cardLine := fmt.Sprintf("%s|%s|%s|%s|%s", cc.CardNumber, cc.ExpMonth, cc.ExpYear, cc.Cvv, r.Code)
	_, err = fmt.Fprintln(f, cardLine)
	if err != nil {
		f.Close()
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
