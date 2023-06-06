package leetcode

// time O(n)
// memory O(n)
func twoSum(nums []int, target int) []int {
	keys := make([]int, 0, 2)
	valToKey := make(map[int]int)
	for key, val := range nums {
		if pos, ok := valToKey[target-val]; ok {
			return []int{pos, key}
		}
		valToKey[val] = key
	}
	return keys
}
