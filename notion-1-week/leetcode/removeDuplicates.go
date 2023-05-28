package leetcode

// time O(n)
// memory O(1)
func removeDuplicates(nums []int) ([]int, int) {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return nums[:k], k
}
