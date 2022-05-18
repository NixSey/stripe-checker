package src

import (
	"bufio"
	"fmt"
	"os"
)

type Handler func(line string)

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

func SaveCard(cc Card, output string, r Result) error {
	f, err := os.OpenFile(output, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, fmt.Sprintf("%s|%s|%s|%s|%s", cc.CardNumber, cc.ExpMonth, cc.ExpYear, cc.Cvv, r.Code))
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
