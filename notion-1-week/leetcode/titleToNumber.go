package leetcode

import "math"

// time O(n)
// memory O(1)
func titleToNumber(columnTitle string) int {
	res := 0
	for i := 0; i < len(columnTitle); i++ {
		charNum := int(columnTitle[i] - 'A' + 1)
		position := len(columnTitle) - i - 1
		res += charNum * int(math.Pow(26, float64(position)))
	}

	return res
}
