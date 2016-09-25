// Package say solves an Exercism challenge.
package say

import (
	"strings"
)

// basicNumbers are all numbers below 20.
var basicNumbers = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

// tensNumbers count how many tens are in the number.
var tensNumbers = []string{
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

// scaleWords are all the scale words we can use.
var scaleWords = []string{
	"",
	"thousand",
	"million",
	"billion",
	"trillion",
	"quadrillion",
	"quintillion",
}

// Say takes an integer and returns that integer
// parsed to english.
func Say(n uint64) string {
	// n = 0 is a special case
	if n == 0 {
		return "zero"
	}

	var res []string
	for scaleIndex := 0; n > 0; scaleIndex, n = scaleIndex+1, n/1000 {
		grouping := n % 1000
		groupingEnglish := sayHundreds(grouping)
		if groupingEnglish != "" {
			// Put the grouping and then the scale word on the front of res
			res = append([]string{groupingEnglish, scaleWords[scaleIndex]}, res...)
		}
	}

	return strings.TrimSpace(strings.Join(res, " "))
}

// sayHundreds takes an integer between 0 and 999 and
// returns it's English representation.
func sayHundreds(n uint64) string {
	// n == 0 is a special case
	if n == 0 {
		return ""
	}

	var res []string

	// Hundreds place
	hundreds := n / 100
	if hundreds > 0 {
		res = append(res, basicNumbers[hundreds])
		res = append(res, "hundred")
		n -= 100 * hundreds
	}
	if n == 0 {
		return strings.Join(res, " ")
	}

	// Tens place, or ones place if no tens
	tens := n / 10
	if tens < 2 {
		res = append(res, basicNumbers[n])
		n = 0
	} else {
		res = append(res, tensNumbers[tens-2])
		n -= 10 * tens
	}
	if n == 0 {
		return strings.Join(res, " ")
	}

	// Ones place after some number of tens
	res[len(res)-1] += "-" + basicNumbers[n]

	return strings.Join(res, " ")
}
