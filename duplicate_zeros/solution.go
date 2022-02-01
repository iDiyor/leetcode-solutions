package duplicatezeros

/*
Given a fixed-length integer array arr, duplicate each occurrence of zero,
shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written.
Do the above modifications to the input array in place and do not return anything.

Example 1:

Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

Example 2:

Input: arr = [1,2,3]
Output: [1,2,3]
Explanation: After calling your function, the input array is modified to: [1,2,3]

Example 3:

Input: arr = [1,0,3]
Output: [1,0,0]
*/

func duplicateZeros(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			shiftToRightFromIdx(i, arr)
		}
	}
}

func shiftToRightFromIdx(idx int, arr []int) {
	for i := len(arr) - 2; i >= idx; i-- {
		arr[i+1] = arr[i]
	}
}

func duplicateZeros2(arr []int) {
	zeros := 0

	for _, v := range arr {
		if v == 0 {
			zeros++
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if zeros+i < len(arr) {
				arr[zeros+i] = 0
			}

			if zeros-1+i < len(arr) {
				arr[zeros-1+i] = 0
			}

			zeros--
		} else if i+zeros < len(arr) {
			arr[zeros+i] = arr[i]
		}
	}
}
