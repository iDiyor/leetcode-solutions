package removeelement

import "testing"

func TestRemoveElement(t *testing.T) {
	type test struct {
		Nums     []int
		Val      int
		Expected int
	}

	tests := []test{
		{
			Nums:     []int{3, 2, 2, 3},
			Val:      3,
			Expected: 2,
		},
		{
			Nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			Val:      2,
			Expected: 5,
		},
		{
			Nums:     []int{},
			Val:      2,
			Expected: 0,
		},
		{
			Nums:     []int{3, 3},
			Val:      3,
			Expected: 0,
		},
		{
			Nums:     []int{2, 2, 3},
			Val:      2,
			Expected: 1,
		},
		{
			Nums:     []int{3, 3, 3},
			Val:      3,
			Expected: 0,
		},
	}

	for _, test := range tests {
		result := removeElement(test.Nums, test.Val)
		if result != test.Expected {
			t.Errorf("Expected %v, got %d for %v", test.Expected, result, test.Nums)
		}
	}
}
