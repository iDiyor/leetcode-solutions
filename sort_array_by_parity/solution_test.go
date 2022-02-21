package sortarraybyparity

import (
	"testing"

	"github.com/iDiyor/leetcode-solutions/utils"
)

func TestSortArrayByParity(t *testing.T) {
	type test struct {
		Nums     []int
		Expected []int
	}

	tests := []test{
		{
			Nums:     []int{3, 1, 2, 4},
			Expected: []int{2, 4, 3, 1},
		},
		{
			Nums:     []int{0},
			Expected: []int{0},
		},
		{
			Nums:     []int{2, 3, 4},
			Expected: []int{2, 4, 3},
		},
		{
			Nums:     []int{0, 2, 1, 4},
			Expected: []int{0, 2, 4, 1},
		},
	}

	for _, test := range tests {
		result := sortArrayByParity(test.Nums)
		if !utils.IsSliceEqual(result, test.Expected) {
			t.Errorf("Expected %v, got %v", test.Expected, result)
		}
	}
}
