package src

import (
	"bufio"
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
