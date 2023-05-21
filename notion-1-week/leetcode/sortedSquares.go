package leetcode

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	i, j := 0, len(nums)-1
	for k := len(nums) - 1; k >= 0; k-- {
		westVal := nums[i]
		eastVal := nums[j]
		westSquare := westVal * westVal
		eastSquare := eastVal * eastVal

		if westSquare > eastSquare {
			res[k] = westSquare
			i++
		} else {
			res[k] = eastSquare
			j--
		}
	}

	return res
}
