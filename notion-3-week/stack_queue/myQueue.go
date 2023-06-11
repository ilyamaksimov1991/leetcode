package stack_queue

type MyQueue struct {
	queue []int
}

func Constructor() MyQueue {
	return MyQueue{
		queue: make([]int, 0, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.queue = append(q.queue, x)
}

func (q *MyQueue) Pop() int {
	if q.Empty() {
		return 0
	}
	res := q.queue[0]
	q.queue = q.queue[1:]
	return res
}

func (q *MyQueue) Peek() int {
	if q.Empty() {
		return 0
	}

	return q.queue[0]
}

func (q *MyQueue) Empty() bool {
	return len(q.queue) == 0
}
