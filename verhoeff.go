package checksum

import (
	"errors"
	"strconv"
)

var multiplicationTable = map[int]map[int]int{
	0: {0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9},
	1: {0: 1, 1: 2, 2: 3, 3: 4, 4: 0, 5: 6, 6: 7, 7: 8, 8: 9, 9: 5},
	2: {0: 2, 1: 3, 2: 4, 3: 0, 4: 1, 5: 7, 6: 8, 7: 9, 8: 5, 9: 6},
	3: {0: 3, 1: 4, 2: 0, 3: 1, 4: 2, 5: 8, 6: 9, 7: 5, 8: 6, 9: 7},
	4: {0: 4, 1: 0, 2: 1, 3: 2, 4: 3, 5: 9, 6: 5, 7: 6, 8: 7, 9: 8},
	5: {0: 5, 1: 9, 2: 8, 3: 7, 4: 6, 5: 0, 6: 4, 7: 3, 8: 2, 9: 1},
	6: {0: 6, 1: 5, 2: 9, 3: 8, 4: 7, 5: 1, 6: 0, 7: 4, 8: 3, 9: 2},
	7: {0: 7, 1: 6, 2: 5, 3: 9, 4: 8, 5: 2, 6: 1, 7: 0, 8: 4, 9: 3},
	8: {0: 8, 1: 7, 2: 6, 3: 5, 4: 9, 5: 3, 6: 2, 7: 1, 8: 0, 9: 4},
	9: {0: 9, 1: 8, 2: 7, 3: 6, 4: 5, 5: 4, 6: 3, 7: 2, 8: 1, 9: 0},
}

var inverseTable = map[int]int{
	0: 0,
	1: 4,
	2: 3,
	3: 2,
	4: 1,
	5: 5,
	6: 6,
	7: 7,
	8: 8,
	9: 9,
}

var permutationTable = map[int]map[int]int{
	0: {0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9},
	1: {0: 1, 1: 5, 2: 7, 3: 6, 4: 2, 5: 8, 6: 3, 7: 0, 8: 9, 9: 4},
	2: {0: 5, 1: 8, 2: 0, 3: 3, 4: 7, 5: 9, 6: 6, 7: 1, 8: 4, 9: 2},
	3: {0: 8, 1: 9, 2: 1, 3: 6, 4: 0, 5: 4, 6: 3, 7: 5, 8: 2, 9: 7},
	4: {0: 9, 1: 4, 2: 5, 3: 3, 4: 1, 5: 2, 6: 6, 7: 8, 8: 7, 9: 0},
	5: {0: 4, 1: 2, 2: 8, 3: 6, 4: 5, 5: 7, 6: 3, 7: 9, 8: 0, 9: 1},
	6: {0: 2, 1: 7, 2: 9, 3: 3, 4: 8, 5: 0, 6: 6, 7: 4, 8: 1, 9: 5},
	7: {0: 7, 1: 0, 2: 4, 3: 6, 4: 9, 5: 1, 6: 3, 7: 2, 8: 5, 9: 8},
}

// Verhoeff implements the Verhoeff algorithm.
type Verhoeff struct{}

// The Verhoeff checksum calculation is performed as follows:
// 1. Create an array n out of the individual digits of the number, taken from right to left (rightmost digit is n0, etc.).
// 2. Initialize the checksum c to zero.
// 3. For each index i of the array n, starting at zero, replace c with d(c, p(i mod 8, ni)).
// The original number is valid if and only if c = 0.
// To generate a check digit, append a 0, perform the calculation: the correct check digit is inv(c).

func verhoeff(s string) (int, error) {
	c := 0
	for i := range s {
		char := s[len(s)-1-i] // get characters in reverse order
		digit := 0
		if char < 48 || char > 57 { // Checks character is between [0-9]
			return 0, errors.New("Not a numeric string")
		}
		digit = int(char - '0')
		c = multiplicationTable[c][permutationTable[i%8][digit]]
	}
	return c, nil
}

// Check the checksum of a given string representing a numeric value, always assuming the last digit is the checksum.
func (*Verhoeff) Check(s string) (bool, error) {
	c, err := verhoeff(s)
	if err != nil {
		return false, err
	}
	if c != 0 {
		return false, nil
	}
	return true, nil
}

// Compute the checksum digit of a given string representing a numeric value.
func (*Verhoeff) Compute(s string) (int, string, error) {
	c, err := verhoeff(s + "0")
	if err != nil {
		return 0, "", err
	}
	return inverseTable[c], s + strconv.Itoa(inverseTable[c]), nil
}
