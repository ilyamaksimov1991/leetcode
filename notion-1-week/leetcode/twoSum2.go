package leetcode

func twoSum2(nums []int, target int) []int {
	for i, val := range nums {
		for j := i + 1; j <= len(nums)-1; j++ {
			result := val + nums[j]
			println(result)
			if result == target {
				return []int{i + 1, j + 1}
			}
			if result > target {
				break
			}
		}
	}
	return []int{}
}
