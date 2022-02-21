package findallnumbersdisappearedinarray

import (
	"testing"

	"github.com/iDiyor/leetcode-solutions/utils"
)

type test struct {
	Nums     []int
	Expected []int
}

func TestFindDisappearedNumbers(t *testing.T) {
	tests := []test{
		{
			Nums:     []int{4, 3, 2, 7, 8, 2, 3, 1},
			Expected: []int{5, 6},
		},
		{
			Nums:     []int{1, 1},
			Expected: []int{2},
		},
	}

	for _, test := range tests {
		result := findDisappearedNumbers(test.Nums)
		if !utils.IsSliceEqual(result, test.Expected) {
			t.Errorf("expected %v, got %v", test.Expected, result)
		}
	}
}

func TestFindDisappearedNumbers2(t *testing.T) {
	tests := []test{
		{
			Nums:     []int{4, 3, 2, 7, 8, 2, 3, 1},
			Expected: []int{5, 6},
		},
		{
			Nums:     []int{1, 1},
			Expected: []int{2},
		},
	}

	for _, test := range tests {
		result := findDisappearedNumbers2(test.Nums)
		if !utils.IsSliceEqual(result, test.Expected) {
			t.Errorf("expected %v, got %v", test.Expected, result)
		}
	}
}
