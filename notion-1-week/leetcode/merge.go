package leetcode

// time O(n)
// memory O(1)
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	iIndex, jIndex, commonIndex := m-1, n-1, m+n-1
	for ; jIndex >= 0; commonIndex-- {
		if iIndex >= 0 && nums1[iIndex] > nums2[jIndex] {
			nums1[commonIndex] = nums1[iIndex]
			iIndex--
		} else {
			nums1[commonIndex] = nums2[jIndex]
			jIndex--
		}
	}

	return nums1
}
