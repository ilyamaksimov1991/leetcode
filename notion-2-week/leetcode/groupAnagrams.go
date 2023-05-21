package leetcode

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	groupToStr := make(map[string][]string)
	for _, val := range strs {
		val1 := strings.Split(val, "")
		sort.Strings(val1)
		key := strings.Join(val1, "")
		groupToStr[key] = append(groupToStr[key], val)
	}

	res := make([][]string, 0, len(groupToStr))
	for _, val := range groupToStr {
		res = append(res, val)
	}
	return res
}
