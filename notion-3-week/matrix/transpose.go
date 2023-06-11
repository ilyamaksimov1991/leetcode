package matrix

// TODO правильная сложность?
// time O(n)
// memory O(n)
func transpose(matrix [][]int) [][]int {
	res := make([][]int, 0)

	for i := 0; i < len(matrix[0]); i++ {
		res1 := make([]int, 0)
		for j := 0; j < len(matrix); j++ {
			res1 = append(res1, matrix[j][i])
		}

		res = append(res, res1)
	}

	return res
}
