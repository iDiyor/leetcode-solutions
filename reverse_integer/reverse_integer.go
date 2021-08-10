package reverse_integer

func isPalindrome(num int) bool {
	if num < 0 || (num%10 == 0 && num != 0) {
		return false
	}

	revertedNumber := 0
	for num > revertedNumber {
		revertedNumber = revertedNumber*10 + num%10
		num = num / 10
	}

	return num == revertedNumber || num == revertedNumber/10
}
