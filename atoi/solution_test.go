package atoi

import "testing"

func TestAtoi(t *testing.T) {
	type test struct {
		Input    string
		Expected int
	}

	tests := []test{
		{
			Input:    "42",
			Expected: 42,
		},
		{
			Input:    "    -52",
			Expected: -52,
		},
		{
			Input:    "4321 with words",
			Expected: 4321,
		},
		{
			Input:    "   321  ",
			Expected: 321,
		},
	}

	for _, test := range tests {
		result := Atoi(test.Input)
		if result != test.Expected {
			t.Errorf("Expected %d, got %d for input '%s'", test.Expected, result, test.Input)
		}
	}
}
