package movezeros

import (
	"testing"

	"github.com/iDiyor/leetcode-solutions/utils"
)

func TestMoveZeros(t *testing.T) {
	type test struct {
		Nums     []int
		Expected []int
	}

	tests := []test{
		{
			Nums:     []int{0},
			Expected: []int{0},
		},
		{
			Nums:     []int{0, 1, 0, 3, 12},
			Expected: []int{1, 3, 12, 0, 0},
		},
		{
			Nums:     []int{0, 1, 0, 2},
			Expected: []int{1, 2, 0, 0},
		},
	}

	for _, test := range tests {
		moveZeroes(test.Nums)
		if !utils.IsSliceEqual(test.Nums, test.Expected) {
			t.Errorf("Expected %v, got %v", test.Expected, test.Nums)
		}
	}
}
