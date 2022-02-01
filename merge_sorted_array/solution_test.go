package mergesortedarray

import "testing"

func TestMerge(t *testing.T) {
	type test struct {
		Num1     []int
		Num2     []int
		M        int
		N        int
		Expected []int
	}

	tests := []test{
		{
			Num1:     []int{1, 2, 3, 0, 0, 0},
			Num2:     []int{2, 5, 6},
			M:        3,
			N:        3,
			Expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			Num1:     []int{1},
			Num2:     []int{},
			M:        1,
			N:        0,
			Expected: []int{1},
		},
		{
			Num1:     []int{0},
			Num2:     []int{1},
			M:        0,
			N:        1,
			Expected: []int{1},
		},
		{
			Num1:     []int{1, 2, 3, 4, 0, 0},
			Num2:     []int{5, 6},
			M:        4,
			N:        2,
			Expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			Num1:     []int{4, 5, 6, 0, 0, 0},
			Num2:     []int{1, 2, 3},
			M:        3,
			N:        3,
			Expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, test := range tests {
		merge(test.Num1, test.M, test.Num2, test.N)
		if !isEqual(test.Num1, test.Expected) {
			t.Errorf("Expected %v, got %v", test.Expected, test.Num1)
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
