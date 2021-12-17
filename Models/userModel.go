package models

import (
	"regexp"
	"unicode"
)

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name" `
	Email       string `json:"email" `
	Username    string `json:"userName"`
	Password    string `json:"password" `
	DateOfBirth string `json:"dateOfBirth" `
	PhoneNumber string `json:"phoneNumber" `
}

//Check if user has entered valid data or not
func IsValidUser(user User) map[string]string {

	var errs = make(map[string]string)
	// check if the name empty
	if user.Name == "" {
		errs["name1"] = "The name is required!"
	}
	// check  name length
	if len(user.Name) < 2 || len(user.Name) > 25 {
		errs["name2"] = "The name field must be between 2-25 chars!"
	}
	if len(user.Username) < 2 || len(user.Username) > 25 {
		errs["name3"] = "The name field must be between 2-25 chars!"
	}
	if len(user.Password) < 4 || len(user.Password) > 10 {
		errs["name4"] = "The name field must be between 4-10 chars!"
	}
	if user.Email == "" {
		errs["email1"] = "The email field is required!"
	}
	regexpEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !regexpEmail.Match([]byte(user.Email)) {
		errs["email2"] = "The email field should be a valid email address!"
	}

	for _, char := range user.PhoneNumber {
		if unicode.IsDigit(char) {
			continue
		} else {
			errs["phoneNumber1"] = "The phoneNumber  must be Numeric!"
			break
		}

	}

	return errs
}
