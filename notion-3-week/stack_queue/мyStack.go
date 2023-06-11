package stack_queue

type MyStack struct {
	stack []int
}

func Constructor1() MyStack {
	return MyStack{
		stack: make([]int, 0, 0),
	}
}

func (s *MyStack) Push(x int) {
	s.stack = append(s.stack, x)
}

func (s *MyStack) Pop() int {
	if s.Empty() {
		return 0
	}
	res := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return res
}

func (s *MyStack) Top() int {
	if s.Empty() {
		return 0
	}
	return s.stack[len(s.stack)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.stack) == 0
}
