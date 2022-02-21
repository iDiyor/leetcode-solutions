package findpivotindex

import "testing"

func TestPivotIndex(t *testing.T) {
	type test struct {
		Nums     []int
		Expected int
	}

	tests := []test{
		{
			Nums:     []int{1, 7, 3, 6, 5, 6},
			Expected: 3,
		},
		{
			Nums:     []int{1, 2, 3},
			Expected: -1,
		},
		{
			Nums:     []int{2, 1, -1},
			Expected: 0,
		},
		{
			Nums:     []int{0, 2, 1, -1},
			Expected: 1,
		},
		{
			Nums:     []int{-1, -1, -1, 1, 1, -1},
			Expected: -1,
		},
	}

	for _, test := range tests {
		result := pivotIndex(test.Nums)
		if result != test.Expected {
			t.Errorf("expected %v, got %v", test.Expected, result)
		}
	}
}
