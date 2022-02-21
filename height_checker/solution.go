package heightchecker

/*
A school is trying to take an annual photo of all the students.
The students are asked to stand in a single file line in non-decreasing order by height.
Let this ordering be represented by the integer array expected where expected[i] is the expected height of the ith student in line.

You are given an integer array heights representing the current order that the students are standing in.
Each heights[i] is the height of the ith student in line (0-indexed).

Return the number of indices where heights[i] != expected[i].

Example 1:
Input: heights = [1,1,4,2,1,3]
Output: 3
Explanation:
heights:  [1,1,4,2,1,3]
expected: [1,1,1,2,3,4]
Indices 2, 4, and 5 do not match.

Example 2:
Input: heights = [5,1,2,3,4]
Output: 5
Explanation:
heights:  [5,1,2,3,4]
expected: [1,2,3,4,5]
All indices do not match.

Example 3:
Input: heights = [1,2,3,4,5]
Output: 0
Explanation:
heights:  [1,2,3,4,5]
expected: [1,2,3,4,5]
All indices match.


Constraints:
1 <= heights.length <= 100
1 <= heights[i] <= 100
*/

func heightChecker(heights []int) int {
	if len(heights) == 0 || len(heights) == 1 {
		return 0
	}

	expectedHeights := mergeSort(heights, 0, len(heights))
	var counter int

	for idx, height := range heights {
		if expectedHeights[idx] != height {
			counter++
		}
	}

	return counter
}

func mergeSort(arr []int, low, hight int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := low + (hight-low)/2
	left := arr[:middle]
	right := arr[middle:]

	left = mergeSort(left, 0, middle)
	right = mergeSort(right, middle, len(right))
	return merge(left, right)
}

func merge(left, right []int) []int {
	var i, j, k int
	result := make([]int, len(left)+len(right))

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
