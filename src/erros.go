package src

import (
	"log"
)

// simple error handling
func HandleError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
