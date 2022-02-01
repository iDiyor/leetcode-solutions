package sortedsquares

import "testing"

func TestSortedSquares(t *testing.T) {
	type test struct {
		Input    []int
		Expected []int
	}

	tests := []test{
		{
			Input:    []int{-4, -1, 0, 3, 10},
			Expected: []int{0, 1, 9, 16, 100},
		},
		{
			Input:    []int{-7, -3, 2, 3, 11},
			Expected: []int{4, 9, 9, 49, 121},
		},
		{
			Input:    []int{-7, -3, 2, 3, 11, 22},
			Expected: []int{4, 9, 9, 49, 121, 484},
		},
	}

	for _, test := range tests {
		result := SortedSquares(test.Input)
		if !isEqual(result, test.Expected) {
			t.Errorf("Expected %v, got %v for input %v", test.Expected, result, test.Input)
		}
	}
}

func TestSortedSquares2(t *testing.T) {
	type test struct {
		Input    []int
		Expected []int
	}

	tests := []test{
		{
			Input:    []int{-4, -1, 0, 3, 10},
			Expected: []int{0, 1, 9, 16, 100},
		},
		{
			Input:    []int{-7, -3, 2, 3, 11},
			Expected: []int{4, 9, 9, 49, 121},
		},
		{
			Input:    []int{-7, -3, 2, 3, 11, 22},
			Expected: []int{4, 9, 9, 49, 121, 484},
		},
	}

	for _, test := range tests {
		result := SortedSquares2(test.Input)
		if !isEqual(result, test.Expected) {
			t.Errorf("Solution #2: Expected %v, got %v for input %v", test.Expected, result, test.Input)
		}
	}
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, e := range a {
		if e != b[i] {
			return false
		}
	}

	return true
}
