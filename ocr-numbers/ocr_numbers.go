// Package ocr solves an Exercism challenge.
package ocr

import (
	"strings"
)

// These are all valid visual representations of a digit.
// They look terrible, but they're correct.
const (
	vOne   = "   \n  |\n  |\n   "
	vTwo   = " _ \n _|\n|_ \n   "
	vThree = " _ \n _|\n _|\n   "
	vFour  = "   \n|_|\n  |\n   "
	vFive  = " _ \n|_ \n _|\n   "
	vSix   = " _ \n|_ \n|_|\n   "
	vSeven = " _ \n  |\n  |\n   "
	vEight = " _ \n|_|\n|_|\n   "
	vNine  = " _ \n|_|\n _|\n   "
	vZero  = " _ \n| |\n|_|\n   "
)

// recognizeDigit takes a string that, when printed, should
// be a 3x4 rectangle. It attempts to match it with any of
// the above-defined constants, and return the correct
// digit. If there are no matches, '?' is returned.
func recognizeDigit(s string) string {
	switch s {
	case vOne:
		return "1"
	case vTwo:
		return "2"
	case vThree:
		return "3"
	case vFour:
		return "4"
	case vFive:
		return "5"
	case vSix:
		return "6"
	case vSeven:
		return "7"
	case vEight:
		return "8"
	case vNine:
		return "9"
	case vZero:
		return "0"
	}
	return "?"
}

// Recognize passes a string, breaks that string into 3x4 rectangles,
// and passes each rectangle to recognizeDigit. If the input is not
// correctly formatted, then the empty string is returned.
func Recognize(s string) []string {
	s = strings.Trim(s, "\n")
	rows := strings.Split(s, "\n")
	// Make sure the number of rows is divisble by 4.
	if len(rows)%4 != 0 {
		return []string{""}
	}
	toReturn := make([]string, 0)
	// Process each collection of 4 rows, which represents a single number.
	for i := 0; i < len(rows)/4; i++ {
		currentRows := rows[i*4 : (i+1)*4]
		// Make sure every row in currentRows is the same length.
		for i := 1; i < 4; i++ {
			if len(currentRows[0]) != len(currentRows[i]) {
				return []string{""}
			}
		}
		// Make sure that every row length is divisible by 3.
		if len(currentRows[0])%3 != 0 {
			return []string{""}
		}

		vNumbers := make([]string, len(currentRows[0])/3)
		// Create the individual visual digits out of the rows.
		for i := 0; i < 4; i++ { // row counter
			for j := 0; j < len(currentRows[0])/3; j++ { // digit counter
				vNumbers[j] += currentRows[i][j*3 : (j+1)*3] // add three characters to a visual digit
				if i != 3 {
					vNumbers[j] += "\n" // add a newline (as long as we're not on the last row)
				}
			}
		}

		number := ""
		// Make the number out of the visual representation
		for _, vNum := range vNumbers {
			number += recognizeDigit(vNum)
		}

		toReturn = append(toReturn, number)
	}
	return toReturn
}
