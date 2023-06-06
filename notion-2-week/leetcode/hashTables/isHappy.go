package hashTables

// time O(n)
// memory O(n)
func isHappy(n int) bool {
	repeat := make(map[int]struct{}, 0)
	for n > 0 {
		n = mySumOfSquare(n)

		if _, ok := repeat[n]; ok {
			return false
		}
		if n == 1 {
			return true
		}
		repeat[n] = struct{}{}
	}

	return false
}

func mySumOfSquare(n int) int {
	res := 0
	for n > 0 {
		digit := n % 10
		res += digit * digit
		n = n / 10
	}

	return res
}
