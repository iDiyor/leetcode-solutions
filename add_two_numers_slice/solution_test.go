package addtwonumersslice

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	type test struct {
		Input    [][]int
		Expected []int
	}

	tests := []test{
		{
			Input: [][]int{
				{2, 4, 3},
				{5, 6, 4},
			},
			Expected: []int{7, 0, 8},
		},
		{
			Input: [][]int{
				{4, 7, 5},
				{5, 4},
			},
			Expected: []int{9, 1, 6},
		},
		{
			Input: [][]int{
				{0},
				{0},
			},
			Expected: []int{0},
		},
		{
			Input: [][]int{
				{9, 9, 9, 9, 9, 9, 9},
				{9, 9, 9, 9},
			},
			Expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, test := range tests {
		result := addTwoNumbers(test.Input[0], test.Input[1])
		if !isEqual(result, test.Expected) {
			t.Errorf("Expected %v, got %v for a - %v and b - %v", test.Expected, result, test.Input[0], test.Input[1])
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
