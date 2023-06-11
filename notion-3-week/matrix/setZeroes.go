package matrix

// TODO правильная сложность?
// time O(n)
// memory O(1)
func setZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	flagZero := false

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				if r > 0 {
					matrix[r][0] = 0
				} else {
					flagZero = true
				}

			}
		}
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for r := 0; r < rows; r++ {
			matrix[r][0] = 0
		}
	}

	if flagZero {
		for c := 0; c < cols; c++ {
			matrix[0][c] = 0
		}
	}
}

// TODO можно считать за memory O(1)?
func setZeroes2(matrix [][]int) {
	mRow := make(map[int]struct{})
	mCol := make(map[int]struct{})

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if matrix[r][c] == 0 {
				mRow[r] = struct{}{}
				mCol[c] = struct{}{}
			}
		}
	}

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if _, ok := mRow[r]; ok {
				matrix[r][c] = 0
				continue
			}

			if _, ok := mCol[c]; ok {
				matrix[r][c] = 0
			}
		}
	}
}
