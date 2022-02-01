package atoi

import (
	"math"
)

/*
Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).

The algorithm for myAtoi(string s) is as follows:

1. Read in and ignore any leading whitespace.
2. Check if the next character (if not already at the end of the string) is '-' or '+'. Read this character in if it is either.
This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
3. Read in next the characters until the next non-digit character or the end of the input is reached. The rest of the string is ignored.
4. Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
5. If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that it remains in the range.
Specifically, integers less than -231 should be clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
6. Return the integer as the final result.
Note:

Only the space character ' ' is considered a whitespace character.
Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.


Example 1:

Input: s = "42"
Output: 42
Explanation: The underlined characters are what is read in, the caret is the current reader position.
Step 1: "42" (no characters read because there is no leading whitespace)
         ^
Step 2: "42" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "42" ("42" is read in)
           ^
The parsed integer is 42.
Since 42 is in the range [-231, 231 - 1], the final result is 42.
Example 2:

Input: s = "   -42"
Output: -42
Explanation:
Step 1: "   -42" (leading whitespace is read and ignored)
            ^
Step 2: "   -42" ('-' is read, so the result should be negative)
             ^
Step 3: "   -42" ("42" is read in)
               ^
The parsed integer is -42.
Since -42 is in the range [-231, 231 - 1], the final result is -42.
Example 3:

Input: s = "4193 with words"
Output: 4193
Explanation:
Step 1: "4193 with words" (no characters read because there is no leading whitespace)
         ^
Step 2: "4193 with words" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "4193 with words" ("4193" is read in; reading stops because the next character is a non-digit)
             ^
The parsed integer is 4193.
Since 4193 is in the range [-231, 231 - 1], the final result is 4193.
*/

const (
	whiteSpace = " "
	minusSign  = "-"
	plusSign   = "+"
)

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	var currentIdx int
	var isNegative bool

	// s = strings.TrimSpace(s)
	// trim white spaces before digits or plus and minus sign
	for currentIdx < len(s) && string(s[currentIdx]) == whiteSpace {
		currentIdx++
	}

	// if all char of s were white spaces then return
	if currentIdx == len(s) {
		return 0
	}

	// if the begining of non white space character is "-" sign
	// then isNegative is set to true
	if string(s[currentIdx]) == minusSign {
		isNegative = true
	}

	// Move currentIdx position one step forward to point a characted which is expected to be a digit
	if string(s[currentIdx]) == plusSign || string(s[currentIdx]) == minusSign {
		currentIdx++
	}

	var result int

	// Until we don't meet a non-digit (>= 0 && <=9) character
	// we first calculate the Dec representation of that digit character
	// then we multiple the prev. result by 10 and add Dec value of current digit char.
	for currentIdx < len(s) && s[currentIdx] >= '0' && s[currentIdx] <= '9' {
		digit := int(s[currentIdx] - '0')
		result = result*10 + digit
		currentIdx++
	}

	// if the input is negative, we make the result negative
	if isNegative {
		result = result * -1
	}

	// We check the max and min ranges for the overflow
	if result > math.MaxInt32 {
		return math.MaxInt32
	} else if result < math.MinInt32 {
		return math.MinInt32
	}

	return result
}
