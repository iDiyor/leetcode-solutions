package largestnumberatleasttwice

import "testing"

func TestDominantIndex(t *testing.T) {
	type test struct {
		nums     []int
		expected int
	}

	tests := []test{
		{
			nums:     []int{3, 6, 1, 0},
			expected: 1,
		},
		{
			nums:     []int{1, 2, 3, 4},
			expected: -1,
		},
		{
			nums:     []int{1},
			expected: 0,
		},
	}

	for _, test := range tests {
		result := dominantIndex(test.nums)
		if result != test.expected {
			t.Errorf("expected %d, got %d", test.expected, result)
		}
	}
}
