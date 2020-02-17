package utils

import (
	"log"
	"regexp"
)

const (
	emailRegex = `^([a-zA-Z0-9_\-\.]+)@([a-zA-Z0-9_\-\.]+)\.([a-zA-Z]{2,5})$`
)

//CheckEmail ...
func CheckEmail(email string) bool {
	validationResult := false
	r, err := regexp.Compile(emailRegex)
	if err != nil {
		log.Fatal(err)
	}
	validationResult = r.MatchString(email)
	return validationResult
}

//CheckDob ...
func CheckDob(dob string) bool {
	//TODO
	return true
}

//CheckURL ...
func CheckURL(url string) bool {
	//TODO
	return true
}
