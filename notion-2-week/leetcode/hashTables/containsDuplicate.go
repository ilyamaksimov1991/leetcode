package hashTables

// time O(n)
// memory O(n)
func containsDuplicate(nums []int) bool {
	dublicates := make(map[int]struct{}, len(nums))

	for _, val := range nums {
		if _, ok := dublicates[val]; ok {
			return true
		}
		dublicates[val] = struct{}{}
	}

	return false
}
