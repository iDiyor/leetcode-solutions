package longestcommonprefix

import "strings"

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string ""

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lower-case English letters.
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	dict := make(map[string][]int)
	var prefix string

	for idx, str := range strs {
		for _, char := range str {
			prefix += string(char)
			idxList, exist := dict[prefix]
			if exist {
				if !contains(idx, idxList) {
					if idxList == nil {
						dict[prefix] = make([]int, 0)
					}
					dict[prefix] = append(dict[prefix], idx)
				}
			} else {
				dict[prefix] = []int{idx}
			}
		}
		prefix = ""
	}

	for pfx, occurenceList := range dict {
		if len(occurenceList) == len(strs) {
			if len(pfx) > len(prefix) {
				prefix = pfx
			}
		}
	}

	return prefix
}

func contains(num int, nums []int) bool {
	if nums == nil {
		return false
	}
	for _, n := range nums {
		if num == n {
			return true
		}
	}
	return false
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for !strings.Contains(strs[i], prefix) {
			prefix = prefix[0 : len(prefix)-1]
		}
	}

	return prefix
}
