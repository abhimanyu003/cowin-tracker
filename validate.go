package main

import (
	"errors"
	"strconv"
	"time"
)

func validateDate(input string) error {
	_, err := time.Parse("02-01-2006", input)
	if err != nil {
		return errors.New("enter in format MM-DD-YYYY")
	}
	return nil
}

func validateNumber(input string) error {
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.New("invalid number")
	}
	return nil
}
