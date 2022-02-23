package diagonaltraverse

import (
	"testing"

	"github.com/iDiyor/leetcode-solutions/utils"
)

func TestFindDiagonalOrder(t *testing.T) {
	type test struct {
		Mat      [][]int
		Expected []int
	}

	tests := []test{
		{
			Mat: [][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			Expected: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			Mat: [][]int{
				[]int{1, 2},
				[]int{3, 4},
			},
			Expected: []int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		result := findDiagonalOrder(test.Mat)
		if !utils.IsSliceEqual(result, test.Expected) {
			t.Errorf("expected %v, got %v", test.Expected, result)
		}
	}
}
