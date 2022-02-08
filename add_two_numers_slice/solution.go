package addtwonumersslice

/*
Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:

Input: l1 = [4,7,5], l2 = [5,4]
Output: [9,1,6]
Explanation: 574 + 45 = 619.

Example 3:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
*/

func addTwoNumbers(a, b []int) []int {
	var result []int

	if len(a) > len(b) {
		result = addTwoNumbersFirstSmall(b, a)
	} else {
		result = addTwoNumbersFirstSmall(a, b)
	}

	return result
}

func addTwoNumbersFirstSmall(a, b []int) []int {
	sumResult := make([]int, 0)

	var idx, sum, digit, carry int

	for idx < len(a) || idx < len(b) {
		if idx < len(a) { // if idx is still in the range of a
			sum = a[idx] + b[idx]
			digit = (sum + carry) % 10
		} else { // if idx is in the range of b
			sum = b[idx] + carry
			digit = sum % 10
		}

		sumResult = append(sumResult, digit)

		// finding carry for the next iteration
		sumRemainderDiff := sum - digit
		if sumRemainderDiff > 10 {
			carry = sum % 10
		} else if sumRemainderDiff == 10 {
			carry = 1
		} else if sumRemainderDiff <= 0 {
			carry = 0
		}

		idx++
	}

	if carry > 0 {
		sumResult = append(sumResult, carry)
	}

	return sumResult
}
