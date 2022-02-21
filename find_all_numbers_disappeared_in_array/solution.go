package findallnumbersdisappearedinarray

/*
Given an array nums of n integers where nums[i] is in the range [1, n],
return an array of all the integers in the range [1, n] that do not appear in nums.

Example 1:
Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
Explanation: nums len is 8, so n = 8. In the range of [1, 8], 5 and 6 are missing.

Example 2:
Input: nums = [1,1]
Output: [2]
Explanation: n = 2 (len(nums)). The range is [1,2] and we have 2 missing.

Constraints:
n == nums.length
1 <= n <= 10^5
1 <= nums[i] <= n


Follow up: Could you do it without extra space and in O(n) runtime?
You may assume the returned list does not count as extra space.
*/

func findDisappearedNumbers(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	n := len(nums)
	result := make([]int, 0)
	numsMap := make(map[int]int)

	for _, num := range nums {
		numsMap[num]++
	}

	for i := 1; i <= n; i++ {
		if _, exist := numsMap[i]; !exist {
			result = append(result, i)
		}
	}

	return result
}

func findDisappearedNumbers2(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	for _, num := range nums {
		j := abs(num) - 1
		nums[j] = abs(nums[j]) * -1
	}

	result := make([]int, 0)

	for i, num := range nums {
		if num > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
