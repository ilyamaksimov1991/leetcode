package binarysearch

// https://leetcode.com/problems/peak-index-in-a-mountain-array/description/
// 852. Peak Index in a Mountain Array
// time O(logN)
// memory O(1)
func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		med := left + (right-left)/2
		if arr[med] < arr[med+1] {
			left = med + 1
		} else {
			right = med - 1
		}
	}

	return left
}
