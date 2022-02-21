package heightchecker

import "testing"

func TestHeightChecker(t *testing.T) {
	type test struct {
		Heights  []int
		Expected int
	}

	tests := []test{
		{
			Heights:  []int{1, 1, 4, 2, 1, 3},
			Expected: 3,
		},
		{
			Heights:  []int{5, 1, 2, 3, 4},
			Expected: 5,
		},
		{
			Heights:  []int{1, 2, 3, 4, 5},
			Expected: 0,
		},
	}

	for _, test := range tests {
		result := heightChecker(test.Heights)
		if result != test.Expected {
			t.Errorf("Expected %v, got %v", test.Expected, result)
		}
	}
}
