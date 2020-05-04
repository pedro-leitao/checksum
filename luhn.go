package checksum

import (
	"errors"
	"strconv"
)

// Luhn implements the Luhn algorithm.
type Luhn struct{}

// The formula verifies a number against its included check digit, which is usually appended to a partial account number to
// generate the full account number. This number must pass the following test:
//
//    - From the rightmost digit (excluding the check digit) and moving left, double the value of every second digit.
//		The check digit is neither doubled nor included in this calculation; the first digit doubled is the digit located
//		immediately left of the check digit. If the result of this doubling operation is greater than 9 (e.g., 8 × 2 = 16),
//		then add the digits of the result (e.g., 16: 1 + 6 = 7, 18: 1 + 8 = 9) or, alternatively, the same final result
//		can be found by subtracting 9 from that result (e.g., 16: 16 − 9 = 7, 18: 18 − 9 = 9).
//    - Take the sum of all the digits.
//    - If the total modulo 10 is equal to 0 (if the total ends in zero) then the number is valid according to the
//		Luhn formula; otherwise it is not valid.
//
func luhn(s string, compute bool) (int, error) {
	c := 0
	isSecond := compute
	for i := range s {
		char := s[len(s)-1-i]       // get characters in reverse order
		if char < 48 || char > 57 { // Checks character is between [0-9]
			return 0, errors.New("Not a numeric string")
		}
		digit := int(char - '0')
		if isSecond {
			digit = digit * 2
			if digit > 9 {
				digit = digit - 9
			}
		}
		c = c + digit
		isSecond = !isSecond
	}

	return (c * 9) % 10, nil
}

// Check the checksum of a given string representing a numeric value, always assuming the last digit is the checksum.
func (*Luhn) Check(s string) (bool, error) {
	c, err := luhn(s, false)
	if err != nil {
		return false, err
	}
	if c != 0 {
		return false, nil
	}
	return true, nil
}

// Compute the checksum digit of a given string representing a numeric value.
func (*Luhn) Compute(s string) (int, string, error) {
	c, err := luhn(s, true)
	if err != nil {
		return 0, "", err
	}
	return c, s + strconv.Itoa(c), nil
}
