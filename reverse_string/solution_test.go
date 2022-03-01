package reversestring

import "testing"

func TestReverseString(t *testing.T) {
	type test struct {
		s      []byte
		output []byte
	}

	tests := []test{
		{
			s:      []byte{'h', 'e', 'l', 'l', 'o'},
			output: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			s:      []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			output: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	for _, test := range tests {
		reverseString(test.s)
		if string(test.s) != string(test.output) {
			t.Errorf("expected %v, got %v", string(test.output), string(test.s))
		}
	}
}
