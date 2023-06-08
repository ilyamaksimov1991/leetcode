package binarysearch

// time O(logN)
// memory O(1)
// https://www.code-recipe.com/post/binary-search
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for right >= left {
		mid := (left + right) / 2

		curVal := nums[mid]
		if curVal == target {
			return mid
		}

		if curVal > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
