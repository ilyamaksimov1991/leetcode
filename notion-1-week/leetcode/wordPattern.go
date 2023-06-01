package leetcode

import "strings"

// time O(n+m)
// memory O(n+m)
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	charToWord := make(map[byte]string, 0)
	wordToChar := make(map[string]byte, 0)
	for i := 0; i < len(pattern); i++ {
		if char, ok := wordToChar[words[i]]; ok {
			if char != pattern[i] {
				return false
			}
		}
		if word, ok := charToWord[pattern[i]]; ok {
			if word != words[i] {
				return false
			}
		} else {
			charToWord[pattern[i]] = words[i]
			wordToChar[words[i]] = pattern[i]
		}
	}

	return true
}
