package leetcode

// time O(n)
// memory O(n+m)
func isIsomorphic(s string, t string) bool {
	stMap := make(map[byte]byte, len(s))
	tsMap := make(map[byte]byte, len(s))

	for i := 0; i < len(s); i++ {
		if res, ok := tsMap[t[i]]; ok && res != s[i] {
			return false
		}
		if res, ok := stMap[s[i]]; ok && res != t[i] {
			return false
		}

		tsMap[t[i]] = s[i]
		stMap[s[i]] = t[i]
	}

	return true
}
