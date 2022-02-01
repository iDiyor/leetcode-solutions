package sortedsquares

func SortedSquares2(nums []int) []int {
	result := make([]int, len(nums))
	head, tail := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if abs(nums[head]) < abs(nums[tail]) {
			result[i] = nums[tail] * nums[tail]
			tail--
		} else {
			result[i] = nums[head] * nums[head]
			head++
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
