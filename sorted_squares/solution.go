package sortedsquares

/*
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

Example 1:

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Example 2:

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in non-decreasing order.

Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?
*/

func SortedSquares(nums []int) []int {
	for idx := range nums { // O(n)
		nums[idx] = square(nums[idx])
	}

	return sortAsc(nums)
}

func sortAsc(nums []int) []int {
	return mergeSort(nums)
}

func square(num int) int { // O(1)
	return num * num
}

func mergeSort(nums []int) []int { // O(n * logn)
	if len(nums) <= 1 {
		return nums
	}

	middle := int(len(nums) / 2)

	left := nums[0:middle]
	right := nums[middle:]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	var i, j, k int

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}
