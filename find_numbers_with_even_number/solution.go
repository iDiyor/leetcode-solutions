package findnumberswithevennumber

// Given an array nums of integers, return how many of them contain an even number of digits.
/*
Input: nums = [12,345,2,6,7896]
Output: 2
Explanation:
12 contains 2 digits (even number of digits).
345 contains 3 digits (odd number of digits).
2 contains 1 digit (odd number of digits).
6 contains 1 digit (odd number of digits).
7896 contains 4 digits (even number of digits).
Therefore only 12 and 7896 contain an even number of digits.

Example 2:
Input: nums = [555,901,482,1771]
Output: 1
Explanation:
Only 1771 contains an even number of digits.
*/

func findNumbers(nums []int) int {
	var evenCount int

	for _, num := range nums {
		digitCount := getDigit(num)
		if digitCount%2 == 0 {
			evenCount++
		}
	}

	return evenCount
}

func getDigit(num int) int {
	var count int

	for num != 0 {
		num = num / 10
		count++
	}

	return count
}
