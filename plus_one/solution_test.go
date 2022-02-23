package plusone

import (
	"testing"

	"github.com/iDiyor/leetcode-solutions/utils"
)

func TestPlusOne(t *testing.T) {
	type test struct {
		digits   []int
		expected []int
	}

	tests := []test{
		{
			digits:   []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			digits:   []int{4, 3, 2, 1},
			expected: []int{4, 3, 2, 2},
		},
		{
			digits:   []int{9},
			expected: []int{1, 0},
		},
		{
			digits:   []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6},
			expected: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7},
		},
	}

	for _, test := range tests {
		result := plusOne(test.digits)
		if !utils.IsSliceEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
