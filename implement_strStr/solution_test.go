package implementstrstr

import "testing"

func TestStrStr(t *testing.T) {
	type test struct {
		haystack string
		needle   string
		output   int
	}

	tests := []test{
		{
			haystack: "hello",
			needle:   "ll",
			output:   2,
		},
		{
			haystack: "aaaaa",
			needle:   "bba",
			output:   -1,
		},
		{
			haystack: "",
			needle:   "",
			output:   0,
		},
		{
			haystack: "",
			needle:   "a",
			output:   -1,
		},
		{
			haystack: "a",
			needle:   "",
			output:   0,
		},
		{
			haystack: "aaa",
			needle:   "aaaa",
			output:   -1,
		},
		{
			haystack: "mississippi",
			needle:   "issipi",
			output:   -1,
		},
		{
			haystack: "a",
			needle:   "a",
			output:   0,
		},
	}

	for _, test := range tests {
		result := strStr(test.haystack, test.needle)
		if result != test.output {
			t.Errorf("expected %d, got %d", test.output, result)
		}
	}
}
