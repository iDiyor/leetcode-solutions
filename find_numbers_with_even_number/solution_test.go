package findnumberswithevennumber

import "testing"

func TestFindNumbers(t *testing.T) {
	type test struct {
		input    []int
		expected int
	}

	tests := []test{
		{
			input:    []int{12, 345, 2, 6, 7896},
			expected: 2,
		},
		{
			input:    []int{555, 901, 482, 1771},
			expected: 1,
		},
	}

	for _, test := range tests {
		result := findNumbers(test.input)
		if result != test.expected {
			t.Errorf("Expected %d, got %d for input %v", test.expected, result, test.input)
		}
	}
}
