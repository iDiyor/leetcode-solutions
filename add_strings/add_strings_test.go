package add_strings

import "testing"

func TestAddStrings(t *testing.T) {
	type test struct {
		str1   string
		str2   string
		result string
	}

	tests := []test{
		{
			"11",
			"123",
			"134",
		},
		{
			"111",
			"122",
			"233",
		},
		{
			"211",
			"125",
			"336",
		},
		{
			"219",
			"126",
			"345",
		},
	}

	for _, test := range tests {
		actualResult := addStrings(test.str1, test.str2)
		if actualResult != test.result {
			t.Errorf("expected %s, got %s", test.result, actualResult)
		}
	}
}
