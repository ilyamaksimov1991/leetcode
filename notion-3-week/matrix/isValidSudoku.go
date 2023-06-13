package matrix

// 36. Valid Sudoku
// time O(n**2)
// memory O(n**2)
func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	squares := [3][3][9]bool{}

	for c := 0; c < 9; c++ {
		for r := 0; r < 9; r++ {
			val := board[c][r]
			if val == '.' {
				continue
			}

			digit := int(val-'0') - 1

			if rows[c][digit] || cols[digit][r] || squares[c/3][r/3][digit] {
				return false
			}

			rows[c][digit] = true
			cols[digit][r] = true
			squares[c/3][r/3][digit] = true
		}
	}

	return true
}
