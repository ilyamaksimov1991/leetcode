package leetcode

func distributeCandies(candyType []int) int {
	types := make(map[int]struct{}, len(candyType))

	for _, val := range candyType {
		types[val] = struct{}{}
	}

	res := 0
	if len(candyType)/2 > len(types) {
		res = len(types)
	} else {
		res = len(candyType) / 2
	}

	return res
}
