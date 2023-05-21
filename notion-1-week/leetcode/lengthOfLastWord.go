package leetcode

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == byte(' ') {
			if res == 0 {
				continue
			} else {
				break
			}
		}
		res++
	}
	return res
}
func lengthOfLastWord2(s string) int {
	parts := strings.Fields(s)
	return len(parts[len(parts)-1])
}
