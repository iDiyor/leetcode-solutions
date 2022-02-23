package addbinary

import "testing"

type test struct {
	a        string
	b        string
	expected string
}

func buildTests() []test {
	return []test{
		{
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
		{
			a:        "0",
			b:        "0",
			expected: "0",
		},
		{
			a:        "1111",
			b:        "1111",
			expected: "11110",
		},
		{
			a:        "100",
			b:        "110010",
			expected: "110110",
		},
	}
}

func TestAddBinary(t *testing.T) {
	for _, test := range buildTests() {
		result := addBinary(test.a, test.b)
		if result != test.expected {
			t.Errorf("expected %s, got %s", test.expected, result)
		}
	}
}

func TestAddBinary2(t *testing.T) {
	for _, test := range buildTests() {
		result := addBinary2(test.a, test.b)
		if result != test.expected {
			t.Errorf("expected %s, got %s", test.expected, result)
		}
	}
}

func TestFillWithZero(t *testing.T) {
	type test struct {
		s        string
		len      int
		expected string
	}

	tests := []test{
		{
			s:        "11",
			len:      2,
			expected: "0011",
		},
		{
			s:        "0011",
			len:      1,
			expected: "00011",
		},
		{
			s:        "",
			len:      3,
			expected: "000",
		},
	}

	for _, test := range tests {
		result := fillWithZero(test.s, test.len)
		if result != test.expected {
			t.Errorf("expected %s, got %s", test.expected, result)
		}
	}
}

func BenchmarkAddBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range buildTests() {
			_ = addBinary(test.a, test.b)
		}
	}
}

func BenchmarkAddBinary2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range buildTests() {
			_ = addBinary2(test.a, test.b)
		}
	}
}
