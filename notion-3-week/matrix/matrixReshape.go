package matrix

// TODO правильная сложность?
// time O(n)
// memory O(n)
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != c*r {
		return mat
	}

	res := make([][]int, 0, r)
	for i := 0; i < r; i++ {
		res = append(res, make([]int, c))
	}

	curR := 0
	curC := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			res[curR][curC] = mat[i][j]

			curC++
			if curC == c {
				curC = 0
				curR++
			}
		}
	}

	return res
}
