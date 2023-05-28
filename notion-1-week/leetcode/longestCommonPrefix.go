package leetcode

// time O(n*2)
// memory O(1)
func longestCommonPrefix(strs []string) string {
	for index, val := range strs[0] {
		for _, str := range strs[1:] {
			if len(str) < index || str[index] != byte(val) {
				return str[:index]
			}
		}
	}
	return ""
}

func longestCommonPrefix2(strs []string) string {
	res := ""
	b := 0

loop:
	for {
		var resLocal byte
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) <= b {
				break loop
			}
			if i == 0 {
				resLocal = strs[i][b]
			} else {
				if resLocal != strs[i][b] {
					break loop
				}
			}

		}
		b++
		res += string(resLocal)
		println(res)
	}

	return res
}
