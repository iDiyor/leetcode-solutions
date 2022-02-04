package validmountainarray

import "testing"

type test struct {
	Arr      []int
	Expected bool
}

func TestValidMountainArray(t *testing.T) {
	tests := []test{
		{
			Arr:      []int{2, 1},
			Expected: false,
		},
		{
			Arr:      []int{3, 5, 5},
			Expected: false,
		},
		{
			Arr:      []int{0, 3, 2, 1},
			Expected: true,
		},
		{
			Arr:      []int{0, 2, 3, 4, 5, 2, 1, 0},
			Expected: true,
		},
		{
			Arr:      []int{2, 0, 2},
			Expected: false,
		},
		{
			Arr:      []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			Expected: false,
		},
	}

	for _, test := range tests {
		result := validMountainArray(test.Arr)
		if result != test.Expected {
			t.Errorf("Expected %v, got %v for %v", test.Expected, result, test.Arr)
		}
	}
}

func TestValidMountainArray2(t *testing.T) {
	tests := []test{
		{
			Arr:      []int{2, 1},
			Expected: false,
		},
		{
			Arr:      []int{3, 5, 5},
			Expected: false,
		},
		{
			Arr:      []int{0, 3, 2, 1},
			Expected: true,
		},
		{
			Arr:      []int{0, 2, 3, 4, 5, 2, 1, 0},
			Expected: true,
		},
		{
			Arr:      []int{2, 0, 2},
			Expected: false,
		},
		{
			Arr:      []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			Expected: false,
		},
	}

	for _, test := range tests {
		result := validMountainArray2(test.Arr)
		if result != test.Expected {
			t.Errorf("Expected %v, got %v for %v", test.Expected, result, test.Arr)
		}
	}
}
