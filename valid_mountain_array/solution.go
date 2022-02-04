package validmountainarray

/*
Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:
> arr.length >= 3
> There exists some i with 0 < i < arr.length - 1 such that:
	+ arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
	+ arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

Example 1:
Input: arr = [2,1]
Output: false

Example 2:
Input: arr = [3,5,5]
Output: false

Example 3:
Input: arr = [0,3,2,1]
Output: true

*/

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	i, j := 0, 1
	var reachedPeak bool

	if arr[j] < arr[i] {
		return false
	}

	for j < len(arr) {
		if arr[j] == arr[i] {
			return false
		}

		if arr[j] < arr[i] && j > 1 {
			reachedPeak = true
		} else if arr[j] > arr[i] && reachedPeak {
			return false
		}

		i++
		j++
	}

	if !reachedPeak {
		return false
	}

	return true
}

func validMountainArray2(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	i, j := 0, len(arr)

	// walk up
	for i+1 < j && arr[i+1] > arr[i] {
		i++
	}

	if i == 0 || i == j-1 {
		return false
	}

	// walk down
	for i+1 < j && arr[i+1] < arr[i] {
		i++
	}

	return i == j-1
}
