package duplicatezeros

import (
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	type test struct {
		Input    []int
		Expected []int
	}

	tests := []test{
		{
			Input:    []int{},
			Expected: []int{},
		},
		{
			Input:    []int{1, 2, 3},
			Expected: []int{1, 2, 3},
		},
		{
			Input:    []int{1, 0, 3},
			Expected: []int{1, 0, 0},
		},
		{
			Input:    []int{0, 0},
			Expected: []int{0, 0},
		},
		{
			Input:    []int{5, 0},
			Expected: []int{5, 0},
		},
		{
			Input:    []int{1, 0, 0, 3, 5},
			Expected: []int{1, 0, 0, 0, 0},
		},
		{
			Input:    []int{1, 2, 0, 4, 5},
			Expected: []int{1, 2, 0, 0, 4},
		},
		{
			Input:    []int{1, 0, 2, 3, 0, 4, 5, 0},
			Expected: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
	}

	for _, test := range tests {
		duplicateZeros(test.Input)
		if !isEqual(test.Input, test.Expected) {
			t.Errorf("Expected %v, got %v", test.Expected, test.Input)
		}
	}
}

func BenchmarkDuplicateZeros(b *testing.B) {
	input := []int{1, 2, 0, 4, 5}
	for i := 0; i < b.N; i++ {
		duplicateZeros(input)
	}
}

func Benchmark2DuplicateZeros(b *testing.B) {
	input := []int{1, 2, 0, 4, 5}
	for i := 0; i < b.N; i++ {
		duplicateZeros2(input)
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
