package reverse_integer

import "testing"

func TestReverseInteger(t *testing.T) {
	type test struct {
		number       int
		isPalindrome bool
	}

	tests := []test{
		{
			121,
			true,
		},
		{
			125,
			false,
		},
		{
			14541,
			true,
		},
		{
			98589,
			true,
		},
	}

	for _, test := range tests {
		result := isPalindrome(test.number)
		if result != test.isPalindrome {
			t.Errorf("expected %v for %d, got %v", test.isPalindrome, test.number, result)
		}
	}
}
