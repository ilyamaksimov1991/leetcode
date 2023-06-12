package binarysearch

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array
// 34. Find First and Last Position of Element in Sorted Array
// time O(logn)
// memory O(1)
func searchRange(nums []int, target int) []int {
	left := binSearch(nums, target, true)
	right := binSearch(nums, target, false)
	return []int{left, right}
}

func binSearch(nums []int, target int, isLeftBin bool) int {
	left, right := 0, len(nums)-1

	res := -1
	for left <= right {
		med := left + (right-left)/2
		v := nums[med]
		switch {
		case v == target:
			res = med
			if isLeftBin {
				right = med - 1
			} else {
				left = med + 1
			}

		case v > target:
			right = med - 1
		case v < target:
			left = med + 1
		}
	}

	return res
}
