package firstnonrepeatingchar

/*
Given a string s, find the first non-repeating character in it and return its index.
If it does not exist, return -1.

Example 1:
Input: s = "leetcode"
Output: 0

Example 2:
Input: s = "loveleetcode"
Output: 2

Example 3:
Input: s = "aabb"
Output: -1
*/

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	charMap := make(map[string]int)

	for _, char := range s {
		charMap[string(char)]++
	}

	for idx, char := range s {
		if charMap[string(char)] == 1 {
			return idx
		}
	}

	return -1
}
