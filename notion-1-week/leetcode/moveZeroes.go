package leetcode

// time O(n*2)
// memory O(1)
func moveZeroes(nums []int) []int {
	j := len(nums) - 1
	for i, num := range nums {
		if num != 0 {
			continue
		}
		for k := j; k < j; k-- {
			if nums[k] != 0 {
				nums[i], nums[j+1] = nums[j+1], nums[i]
				j = k
				break
			}
		}
	}

	return nums
}

// time O(n)
// memory O(1)
func moveZeroes2(nums []int) []int {
	j := 0
	for current, num := range nums {
		if num != 0 {
			nums[current], nums[j] = nums[j], nums[current]
			j++
		}
	}

	return nums
}
