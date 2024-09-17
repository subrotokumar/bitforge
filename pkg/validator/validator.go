package validator

import (
	"errors"
	"strconv"
)

// Validate credit card using Luhn's Algorithm.
func CreditCardValidator(cardNumber string) error {
	len := len(cardNumber)
	if len != 16 {
		return errors.New("invalid card number length")
	}
	sum := 0
	for index, ch := range cardNumber {
		digit, err := strconv.Atoi(string(ch))
		if err != nil {
			return errors.New("invalid credit card number")
		}
		if index%2 != 0 {
			sum += digit
			continue
		}
		digit = 2 * digit
		if digit > 9 {
			sum += (digit / 10) + (digit % 10)
		}
	}
	if sum%10 == 0 {
		return nil
	}
	return errors.New("card is invalid")
}
