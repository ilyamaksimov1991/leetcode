package leetcode

// time O(n)
// memory O(1)
func reverseString(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}

	return s
}

// time O(n)
// memory O(1)
func reverseString2(s []string) []string {
	j := len(s) - 1
	for i, _ := range s {
		if i >= j {
			break
		}
		s[i], s[j] = s[j], s[i]
		j--
	}

	return s
}
