package leetcode

func convertToTitle(columnNumber int) string {
	b := columnNumber

	res := []rune{}

	if columnNumber <= 26 {
		return string(byte(columnNumber) + 'A' - 1)
	}
	for {
		residual := b % 26
		println(residual)
		if residual != 0 {
			res = append(res, rune(residual)+'A'-1)
			println("yes")
			b = b - residual
			continue
		}

		b = b / 26

		if b == 0 {
			break
		}

		if b <= 26 {
			res = append(res, rune((b)+'A'-1))
			break
		}
	}

	res1 := []rune{}
	for i := len(res) - 1; i >= 0; i-- {
		res1 = append(res1, res[i])
	}

	return string(res1)
}
