package maxconsecutiveones

import "testing"

func TestMaxConsecutiveOnes(t *testing.T) {
	type test struct {
		input    []int
		expected int
	}

	tests := []test{
		{
			input:    []int{1, 1, 0, 1, 1, 1},
			expected: 3,
		},
		{
			input:    []int{1, 1, 1, 0, 1, 1, 1, 1, 0, 1},
			expected: 4,
		},
	}

	for _, test := range tests {
		result := findMaxConsecutiveOnes(test.input)
		if result != test.expected {
			t.Errorf("Expected %d, got %d for given input %v", test.expected, result, test.input)
		}
	}
}
