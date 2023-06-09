package binarysearch

// time O(logN)
// memory O(1)
// https://www.youtube.com/watch?v=K-RYzDZkzCI
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		medium := (left + right) / 2
		value := nums[medium]
		switch {
		case value == target:
			return medium
		case value > target:
			right = medium - 1
		case value < target:
			left = medium + 1
		}
	}

	return left
}
