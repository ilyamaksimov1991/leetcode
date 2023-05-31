package leetcode

import (
	"math"
	"strings"
)

// time O(n)
// memory O(1)
// https://www.code-recipe.com/post/string_to_integer
func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}

	start := 0
	signMultiplier := 1
	if s[0] == '-' {
		signMultiplier = -1
		start++
	} else if s[0] == '+' {
		start++
	}

	res := 0
	for i := start; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return res * signMultiplier
		}

		res = res*10 + int(s[i]-'0')

		if res*signMultiplier <= math.MinInt32 {
			return math.MinInt32
		} else if res*signMultiplier >= math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return res * signMultiplier
}
