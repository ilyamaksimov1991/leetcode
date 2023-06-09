package binarysearch

// time O(logN)
// memory O(1)
func mySqrt(x int) int {
	left, right := 0, x
	res := 0

	for left <= right {
		med := left + ((right - left) / 2)
		val := med * med
		if val == x {
			return med
		}

		if val < x {
			left = med + 1
			res = med
		} else {
			right = med - 1
		}
	}

	return res
}
