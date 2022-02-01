package maxconsecutiveones

// Given a binary array nums, return the maximum number of consecutive 1's in the array.
/* Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
*/
func findMaxConsecutiveOnes(nums []int) int {
	var max, count int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}

	if count > max {
		max = count
	}

	return max
}
