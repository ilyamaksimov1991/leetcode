package matrix

// TODO правильная сложность?
// time O(n)
// memory O(1)
func diagonalSum(mat [][]int) int {
	res := 0

	rows := len(mat)
	for i := 0; i < rows; i++ {
		res += mat[i][i]
		if i != rows-i-1 {
			res += mat[i][rows-i-1]
		}
	}

	return res
}
