package stack_queue

// time O(n)
// memory O(n)
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0)
	for index, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			lastIndex := stack[len(stack)-1]
			distance := index - lastIndex
			res[lastIndex] = distance
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, index)
	}

	return res
}
