package src

import (
	"log"
)

// handle errors
func HandleError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
