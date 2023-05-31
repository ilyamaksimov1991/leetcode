package leetcode

// https://www.code-recipe.com/post/container-with-most-water

// time O(n)
// memory O(1)
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxRes := 0

	for l < r {
		curArea := min(height[l], height[r]) * (r - l)
		maxRes = max(maxRes, curArea)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
