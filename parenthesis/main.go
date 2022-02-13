package main

import (
	"fmt"
	"strconv"
)

func generateParenthesis(n int) []string {
	// The min number will be have the correct number of bits to correspond with desired parentheses.
	// For valid parentheses, the first character will always be a "(", which corresponds to 1 with
	// The binary analogy.  n corresponds to pairs of parens, so it needs to be doubled, and the 1
	// contributes a bit, so the number of shifts is (n*2)-1.

	// Min could be further optimized so it's first binary value is the smallest number that
	// represents a valid bit/parenthesis pattern

	// Max can be optimized as well to be the largest possible bit/parenthesis pattern.

	min := 1 << (n*2 - 1)
	max := min << 1
	results := make([]string, 0)

	for i := min; i < max; i++ {
		binaryString := strconv.FormatInt(int64(i), 2)
		if isValidParens(binaryString) {
			// It's a valid binary equivalent, convert it to parens and add to the result set
			results = append(results, binToParen(binaryString))
		}
	}
	return results
}

// Replace the binary digit representation with parentheses.
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
	// This would indicate parens were closed before opening - no good!
	for _, b := range binaryString {
		switch b {
		case '1':
			score++
		case '0':
			score--
		}

		if score < 0 {
			// A negative score means more parens have been closed than open.  That's not valid, so stop now.
			return false
		}
	}

	// Now that the test is complete, check the score.  If it isn't zero, there aren't a
	// valid set of opening and closing parentheses.
	return score == 0
}

func main() {
	fmt.Printf("%v\n", generateParenthesis(5))
}
