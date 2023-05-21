package leetcode

func firstUniqChar(s string) int {
	var counts = make([]int, 26)

	// Insert all the character's count into counts array
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
	}

	// Find the first element with count 1
	for i := 0; i < len(s); i++ {
		if counts[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}
func firstUniqChar2(s string) int {
	type repeat struct {
		count     int
		firstChar int
	}

	rToi := make(map[rune]*repeat, 0)
	for i, val := range []rune(s) {
		if _, ok := rToi[val]; ok {
			rToi[val].count++
		} else {
			rToi[val] = &repeat{
				count:     1,
				firstChar: i,
			}
		}
	}

	min := len([]rune(s))
	have := false
	for _, val := range rToi {
		if val.count != 1 {
			continue
		}
		if min > val.firstChar {
			min = val.firstChar
			have = true
		}
	}
	if have {
		return min
	}

	return -1
}
