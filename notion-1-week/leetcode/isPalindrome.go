package leetcode

import (
	"strings"
)

// time O(n)
// memory O(1)
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isCharNum(s[l]) {
			println("yess", string(s[l]))
			l++
		}
		for l < r && !isCharNum(s[r]) {
			println("no")
			r--
		}
		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}

		l, r = l+1, r-1
	}

	return true
}

func isCharNum(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}
