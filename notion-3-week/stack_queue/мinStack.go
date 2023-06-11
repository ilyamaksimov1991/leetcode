package stack_queue

type MinStack struct {
	stack []int
}

func Constructor2() MinStack {
	return MinStack{
		stack: make([]int, 0, 0),
	}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
}

func (s *MinStack) Pop() int {
	if s.Empty() {
		return 0
	}
	res := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return res
}

func (s *MinStack) Top() int {
	if s.Empty() {
		return 0
	}
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) Empty() bool {
	return len(s.stack) == 0
}
