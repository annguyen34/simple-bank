package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUserName = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d - %d", minLength, maxLength)
	}
	return nil

}

func ValidateUserName(value string) error {
	if err := ValidateString(value, 6, 100); err != nil {
		return err
	}
	if !isValidUserName(value) {
		return fmt.Errorf("invalid username")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, 6, 100); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letters and spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	if err := ValidateString(value, 6, 100); err != nil {
		return err
	}
	return nil
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("invalid email")
	}
	return nil
}
