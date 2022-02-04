package replaceelementswithgreatest

import (
	"testing"

	"github.com/iDiyor/leetcode-solutions/utils"
)

type test struct {
	Arr      []int
	Expected []int
}

func TestReplaceElements(t *testing.T) {
	tests := []test{
		{
			Arr:      []int{17, 18, 5, 4, 6, 1},
			Expected: []int{18, 6, 6, 6, 1, -1},
		},
		{
			Arr:      []int{400},
			Expected: []int{-1},
		},
	}

	for _, test := range tests {
		result := replaceElements(test.Arr)
		if !utils.IsSliceEqual(result, test.Expected) {
			t.Errorf("Expected %v, got %v", test.Expected, result)
		}
	}
}

func TestReplaceElements2(t *testing.T) {
	tests := []test{
		{
			Arr:      []int{17, 18, 5, 4, 6, 1},
			Expected: []int{18, 6, 6, 6, 1, -1},
		},
		{
			Arr:      []int{400},
			Expected: []int{-1},
		},
	}

	for _, test := range tests {
		result := replaceElements2(test.Arr)
		if !utils.IsSliceEqual(result, test.Expected) {
			t.Errorf("Expected %v, got %v", test.Expected, result)
		}
	}
}

func BenchmarkReplaceElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tests := []test{
			{
				Arr:      []int{17, 18, 5, 4, 6, 1},
				Expected: []int{18, 6, 6, 6, 1, -1},
			},
			{
				Arr:      []int{400},
				Expected: []int{-1},
			},
		}

		for _, test := range tests {
			_ = replaceElements(test.Arr)
		}
	}
}

func BenchmarkReplaceElements2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tests := []test{
			{
				Arr:      []int{17, 18, 5, 4, 6, 1},
				Expected: []int{18, 6, 6, 6, 1, -1},
			},
			{
				Arr:      []int{400},
				Expected: []int{-1},
			},
		}

		for _, test := range tests {
			_ = replaceElements2(test.Arr)
		}
	}
}
