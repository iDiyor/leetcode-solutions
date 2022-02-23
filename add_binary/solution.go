package addbinary

/*
Given two binary strings a and b, return their sum as a binary string.


Example 1:
Input: a = "11", b = "1"
Output: "100"

Example 2:
Input: a = "1010", b = "1011"
Output: "10101"


Constraints:
1 <= a.length, b.length <= 104
a and b consist only of '0' or '1' characters.
Each string does not contain leading zeros except for the zero itself.
*/

func addBinary(a string, b string) string {
	var result []byte

	carry := 0

	i, j := len(a)-1, len(b)-1

	appendDigit := func(sum int) {
		switch sum {
		case 3:
			carry = 1
			result = append([]byte("1"), result...)
		case 2:
			carry = 1
			result = append([]byte("0"), result...)
		case 1:
			carry = 0
			result = append([]byte("1"), result...)
		case 0:
			carry = 0
			result = append([]byte("0"), result...)
		}
	}

	for i >= 0 && j >= 0 {
		digitA := a[i] - '0'
		digitB := b[j] - '0'
		sum := int(digitA) + int(digitB) + carry
		appendDigit(sum)

		i--
		j--
	}

	for i >= 0 {
		digit := a[i] - '0'
		sum := int(digit) + carry
		appendDigit(sum)
		i--
	}

	for j >= 0 {
		digit := b[j] - '0'
		sum := int(digit) + carry
		appendDigit(sum)
		j--
	}

	if carry > 0 {
		result = append([]byte("1"), result...)
	}

	return string(result)
}

func addBinary2(a string, b string) string {
	var result []byte

	if len(a) > len(b) {
		b = fillWithZero(b, len(a)-len(b)) // O(n)
	} else {
		a = fillWithZero(a, len(b)-len(a))
	}

	carry := 0

	for i := len(a) - 1; i >= 0; i-- { // O(n)
		digitA := a[i] - '0'
		digitB := b[i] - '0'
		sum := int(digitA) + int(digitB) + carry

		switch sum {
		case 3:
			carry = 1
			result = append([]byte("1"), result...)
		case 2:
			carry = 1
			result = append([]byte("0"), result...)
		case 1:
			carry = 0
			result = append([]byte("1"), result...)
		case 0:
			carry = 0
			result = append([]byte("0"), result...)
		}
	}

	if carry > 0 {
		result = append([]byte("1"), result...)
	}

	return string(result)
}

func fillWithZero(s string, len int) string {
	var result []byte
	for i := 0; i < len; i++ {
		result = append(result, []byte("0")...)
	}
	result = append(result, []byte(s)...)
	return string(result)
}
