package leetcode

// time O(n)
// memory O(n)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s1 := make(map[rune]int, len(s))
	s2 := make(map[rune]int, len(s))

	for _, val := range s {
		s1[val]++
	}
	for _, val := range t {
		s2[val]++
	}

	for key, val := range s1 {
		if res, ok := s2[key]; !ok || res != val {
			return false
		}
	}

	return true
}
