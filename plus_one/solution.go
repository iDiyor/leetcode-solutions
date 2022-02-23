package plusone

/*
You are given a large integer represented as an integer array digits,
where each digits[i] is the ith digit of the integer.
The digits are ordered from most significant to least significant in left-to-right order.
The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.

Example 1:
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].

Example 2:
Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].

Example 3:
Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].

Constraints:
1 <= digits.length <= 100
0 <= digits[i] <= 9
digits does not contain any leading 0's.

PAY ATTENTION TO CONSTRAINTS. BECAUSE digits.length can be <= 100.
IF YOU TRY TO CREATE INTEGER NUMBER FROM DIGITS EVEN INT64 CANNOT HOLD THE VALUE. BECAUSE MAX DIGITS FOR INT64 is 19 (digits.length <= 19)
*/

// FIRST SOLUTION
// func plusOne(digits []int) []int {
// 	var number int

// 	for idx, digit := range digits {
// 		number += digit * int(math.Pow(10, float64(len(digits)-idx-1)))
// 	}

// 	number += 1
// 	numberDigits := countDigits(number)

// 	result := make([]int, numberDigits)

// 	for i := 0; i < numberDigits; i++ {
// 		digit := number % 10
// 		result[numberDigits-i-1] = digit
// 		number = number / 10
// 	}

// 	return result
// }

// func countDigits(num int) int {
// 	var count int
// 	for num > 0 {
// 		count += 1
// 		num = num / 10
// 	}
// 	return count
// }

func plusOne(digits []int) []int {
	carry := 1 // let's assume plus one means that we have a carry of 1 from some non-existing plus operation

	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		if sum >= 10 {
			digits[i] = sum % 10
			carry = 1
		} else {
			digits[i] = sum
			carry = 0
			return digits
		}
	}

	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}
