package checksum

import (
	"errors"
	"strconv"
)

var operationTable = map[int]map[int]int{
	0: {0: 0, 1: 3, 2: 1, 3: 7, 4: 5, 5: 9, 6: 8, 7: 6, 8: 4, 9: 2},
	1: {0: 7, 1: 0, 2: 9, 3: 2, 4: 1, 5: 5, 6: 4, 7: 8, 8: 6, 9: 3},
	2: {0: 4, 1: 2, 2: 0, 3: 6, 4: 8, 5: 7, 6: 1, 7: 3, 8: 5, 9: 9},
	3: {0: 1, 1: 7, 2: 5, 3: 0, 4: 9, 5: 8, 6: 3, 7: 4, 8: 2, 9: 6},
	4: {0: 6, 1: 1, 2: 2, 3: 3, 4: 0, 5: 4, 6: 5, 7: 9, 8: 7, 9: 8},
	5: {0: 3, 1: 6, 2: 7, 3: 4, 4: 2, 5: 0, 6: 9, 7: 5, 8: 8, 9: 1},
	6: {0: 5, 1: 8, 2: 6, 3: 9, 4: 7, 5: 2, 6: 0, 7: 1, 8: 3, 9: 4},
	7: {0: 8, 1: 9, 2: 4, 3: 5, 4: 3, 5: 6, 6: 2, 7: 0, 8: 1, 9: 7},
	8: {0: 9, 1: 4, 2: 3, 3: 8, 4: 6, 5: 1, 6: 7, 7: 2, 8: 0, 9: 5},
	9: {0: 2, 1: 5, 2: 8, 3: 1, 4: 4, 5: 3, 6: 6, 7: 7, 8: 9, 9: 0},
}

// Damm implements the Damm algorithm.
type Damm struct{}

// The validity of a digit sequence containing a check digit is defined over a quasigroup. A quasigroup table ready for use can be taken from Damm's dissertation (pages 98, 106, 111).[3] It is useful if each main diagonal entry is 0,[1] because it simplifies the check digit calculation.
// Validating a number against the included check digit
//
//    - Set up an interim digit and initialize it to 0.
//    - Process the number digit by digit: Use the number's digit as column index and the interim digit as row index, take the table entry and replace the interim digit with it.
//    - The number is valid if and only if the resulting interim digit has the value of 0.
//
// Calculating the check digit
//
// Prerequisite: The main diagonal entries of the table are 0.
//
//    - Set up an interim digit and initialize it to 0.
//    - Process the number digit by digit: Use the number's digit as column index and the interim digit as row index, take the table entry and replace the interim digit with it.
//    - The resulting interim digit gives the check digit and will be appended as trailing digit to the number.
//
func damm(s string) (int, error) {
	c := 0
	for i := range s {
		char := s[i] // get characters in order
		digit := 0
		if char < 48 || char > 57 { // Checks character is between [0-9]
			return 0, errors.New("Not a numeric string")
		}
		digit = int(char - '0')
		c = operationTable[c][digit]
	}
	return c, nil
}

// Check the checksum of a given string representing a numeric value, always assuming the last digit is the checksum.
func (*Damm) Check(s string) (bool, error) {
	c, err := damm(s)
	if err != nil {
		return false, err
	}
	if c != 0 {
		return false, nil
	}
	return true, nil
}

// Compute the checksum digit of a given string representing a numeric value.
func (*Damm) Compute(s string) (int, string, error) {
	c, err := damm(s)
	if err != nil {
		return 0, "", err
	}
	return c, s + strconv.Itoa(c), nil
}
