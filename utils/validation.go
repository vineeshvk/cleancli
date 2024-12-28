package utils

import (
	"errors"
	"strings"
	"unicode/utf8"
)

func ValidateEmptyString(s string) error {
	if utf8.RuneCountInString(strings.TrimSpace(s)) <= 0 {
		return errors.New("input is mandatory")
	}

	return nil
}
