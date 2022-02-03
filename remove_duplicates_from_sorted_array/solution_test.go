package removeduplicatesfromsortedarray

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	type test struct {
		Nums     []int
		Expected int
	}

	tests := []test{
		{
			Nums:     []int{1, 1, 2},
			Expected: 2,
		},
		{
			Nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			Expected: 5,
		},
		{
			Nums:     []int{1, 2},
			Expected: 2,
		},
		{
			Nums:     []int{1, 2, 2},
			Expected: 2,
		},
		{
			Nums:     []int{1, 1, 2, 2, 3, 3},
			Expected: 3,
		},
		{
			Nums:     []int{1},
			Expected: 1,
		},
	}

	for _, test := range tests {
		result := removeDuplicates(test.Nums)
		if result != test.Expected {
			t.Errorf("Expected %d, got %d", test.Expected, result)
		}
	}
}
