package nanditsdoubleexist

import "testing"

func TestCheckIfExist(t *testing.T) {
	type test struct {
		Arr      []int
		Expected bool
	}

	tests := []test{
		{
			Arr:      []int{10, 2, 5, 3},
			Expected: true,
		},
		{
			Arr:      []int{7, 1, 14, 11},
			Expected: true,
		},
		{
			Arr:      []int{3, 1, 7, 11},
			Expected: false,
		},
	}

	for _, test := range tests {
		result := checkIfExist(test.Arr)
		if result != test.Expected {
			t.Errorf("Expected %v, got %v for %v", test.Expected, result, test.Arr)
		}
	}
}
