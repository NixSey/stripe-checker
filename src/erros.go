package src

// handle errors
func HandleError(e error) {
	if e != nil {
		Fatal(e)
	}
}
