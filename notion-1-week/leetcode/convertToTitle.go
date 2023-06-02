package leetcode

// time O(n)
// memory O(n)
func convertToTitle(columnNumber int) string {
	var result string
	for columnNumber > 0 {
		result = string(65+(columnNumber-1)%26) + result
		columnNumber = (columnNumber - 1) / 26
	}
	return result
}
