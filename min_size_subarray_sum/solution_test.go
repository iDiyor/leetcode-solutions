package minsizesubarraysum

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	type test struct {
		target int
		nums   []int
		output int
	}

	tests := []test{
		{
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			output: 2,
		},
		{
			target: 4,
			nums:   []int{1, 4, 4},
			output: 1,
		},
		{
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			output: 0,
		},
		{
			target: 11,
			nums:   []int{1, 2, 3, 4, 5},
			output: 3,
		},
	}

	for _, test := range tests {
		result := minSubArrayLen(test.target, test.nums)
		if result != test.output {
			t.Errorf("expected %d, got %d", test.output, result)
		}
	}
}
