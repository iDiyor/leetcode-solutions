package add_strings

func addStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	carry := byte(0)
	result := ""

	for i >= 0 || j >= 0 {
		asciiSum := byte(0)

		if i >= 0 {
			asciiSum += num1[i] - '0'
			i--
		}

		if j >= 0 {
			asciiSum += num2[j] - '0'
			j--
		}

		asciiSum += carry

		carry = asciiSum / 10
		remainder := asciiSum % 10

		result = string(remainder+'0') + result
	}

	if carry > 0 {
		result = string(carry+'0') + result
	}

	return result
}
