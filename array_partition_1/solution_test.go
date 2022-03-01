package arraypartition1

import "testing"

func TestArrayPairSum(t *testing.T) {
	type test struct {
		nums   []int
		output int
	}

	tests := []test{
		{
			nums:   []int{1, 4, 3, 2},
			output: 4,
		},
		{
			nums:   []int{6, 2, 6, 5, 1, 2},
			output: 9,
		},
	}

	for _, test := range tests {
		result := arrayPairSum(test.nums)
		if result != test.output {
			t.Errorf("expected %v, got %v", test.output, result)
		}
	}
}
