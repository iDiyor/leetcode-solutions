package thirdmaximumnumber

import "testing"

func TestThirdMax(t *testing.T) {
	type test struct {
		Nums     []int
		Expected int
	}

	tests := []test{
		{
			Nums:     []int{3, 2, 1},
			Expected: 1,
		},
		{
			Nums:     []int{1, 2},
			Expected: 2,
		},
		{
			Nums:     []int{2, 2, 3, 1},
			Expected: 1,
		},
		{
			Nums:     []int{10, 50, 20, 100},
			Expected: 20,
		},
	}

	for _, test := range tests {
		result := thirdMax(test.Nums)
		if result != test.Expected {
			t.Errorf("expected %d, got %d for %v", test.Expected, result, test.Nums)
		}
	}
}
