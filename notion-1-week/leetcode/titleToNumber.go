package leetcode

import "math"

func titleToNumber(columnTitle string) int {
	res := 0
	for i := 0; i < len(columnTitle); i++ {
		charNum := int(columnTitle[i] - 'A' + 1)
		position := len(columnTitle) - i - 1
		res += charNum * int(math.Pow(26, float64(position)))
	}

	return res
}

/*
func titleToNumber(columnTitle string) int {
    result := 0
    for i := len(columnTitle) - 1; i >= 0; i-- {
        c := columnTitle[i]
        result += int(c-'A'+1) * int(math.Pow(26, float64(len(columnTitle)-i-1)))
    }
    return result
}
*/
