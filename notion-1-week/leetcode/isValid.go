package leetcode

func isValid(s string) bool {
	closed := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	q := make([]rune, 0)
	for _, val := range []rune(s) {
		if res, ok := closed[val]; ok {
			if len(q) == 0 {
				return false
			}
			if len(q) > 0 && res == q[len(q)-1] {
				q = q[:len(q)-1]
			} else {
				return false
			}
		} else {
			q = append(q, val)
		}
	}
	if len(q) > 0 {
		return false
	}
	return true
}
