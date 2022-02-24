package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type test struct {
		strs   []string
		output string
	}

	tests := []test{
		{
			strs:   []string{"flower", "flow", "flight"},
			output: "fl",
		},
		{
			strs:   []string{"dog", "racecar", "car"},
			output: "",
		},
	}

	for _, test := range tests {
		result := longestCommonPrefix(test.strs)
		if result != test.output {
			t.Errorf("expected %s, got %s", test.output, result)
		}
	}
}
