package functions

import "log"

// error handling function
func CheckError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
