package evenindexsquare

import (
	"testing"

	"github.com/iDiyor/leetcode-solutions/utils"
)

func TestSquareEven(t *testing.T) {
	type test struct {
		Arr      []int
		Expected []int
	}

	tests := []test{
		{
			Arr:      []int{9, -2, -9, 11, 56, -12, -3},
			Expected: []int{81, -2, 81, 11, 3136, -12, 9},
		},
	}

	for _, test := range tests {
		result := squareEven(test.Arr)
		if !utils.IsSliceEqual(result, test.Expected) {
			t.Errorf("Expected %v, got %v", test.Expected, result)
		}
	}
}
