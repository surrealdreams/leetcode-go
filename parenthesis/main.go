package main

import (
	"fmt"
	"strconv"
)

func generateParenthesis(n int) []string {
	min := 1 << (n*2 - 1)
	max := min << 1
	results := make([]string, 0)

	for i := min; i < max; i++ {
		binaryString := strconv.FormatInt(int64(i), 2)
		if isValidParens(binaryString) {
			results = append(results, binToParen(binaryString))
		}
	}
	return results
}

func binToParen(bin string) string {
	result := ""
	for i := 0; i < len(bin); i++ {
		switch bin[i] {
		case '1':
			result += "("
		case '0':
			result += ")"
		}
	}
	return result
}

func isValidParens(binaryString string) bool {
	score := 0
	// To be valid parenthesis, each paren must be opened before it's closed.  Calculate a
	// "score" for the number by adding 1 for each "1" and subtracting 1 for each "0".
	// The score must always end equal to 0 and must never become negative during the calculation.
	// This would indicate parens were closed before opening - invalid!
	for _, b := range binaryString {
		switch b {
		case '1':
			score++
		case '0':
			score--
		}
		if score < 0 {
			return false
		}
	}

	// debug
	// if score == 0 {
	// 	fmt.Printf("%q\n", binaryString)
	// }
	return score == 0
}

func main() {
	fmt.Printf("%v\n", generateParenthesis(5))
}
