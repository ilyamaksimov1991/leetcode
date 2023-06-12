package matrix

// https://leetcode.com/problems/spiral-matrix/submissions/969748701/
// 54. Spiral Matrix
// time O(n)
// memory O(1)
func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}

	res := make([]int, 0, len(matrix)*len(matrix[0]))
	left, right := 0, len(matrix[0])-1
	top, bottom := 0, len(matrix)-1

	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}

		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return res
}
