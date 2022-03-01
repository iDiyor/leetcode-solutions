package longestcommonprefix

import "testing"

type test struct {
	strs   []string
	output string
}

func getTests() []test {
	return []test{
		{
			strs:   []string{"flower", "flow", "flight"},
			output: "fl",
		},
		{
			strs:   []string{"dog", "racecar", "car"},
			output: "",
		},
		{
			strs:   []string{"c", "acc", "ccc"},
			output: "",
		},
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, test := range getTests() {
		result := longestCommonPrefix(test.strs)
		if result != test.output {
			t.Errorf("expected %s, got %s", test.output, result)
		}
	}
}

func TestLongestCommonPrefix2(t *testing.T) {
	for _, test := range getTests() {
		result := longestCommonPrefix2(test.strs)
		if result != test.output {
			t.Errorf("expected %s, got %s", test.output, result)
		}
	}
}
