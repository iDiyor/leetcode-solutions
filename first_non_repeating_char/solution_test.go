package firstnonrepeatingchar

import "testing"

func TestFirstUniqChar(t *testing.T) {
	type test struct {
		Input    string
		Expected int
	}

	tests := []test{
		{
			Input:    "leetcode",
			Expected: 0,
		},
		{
			Input:    "loveleetcode",
			Expected: 2,
		},
		{
			Input:    "aabb",
			Expected: -1,
		},
	}

	for _, test := range tests {
		result := firstUniqChar(test.Input)
		if result != test.Expected {
			t.Errorf("Expected %d, got %d for input %s", test.Expected, result, test.Input)
		}
	}
}
