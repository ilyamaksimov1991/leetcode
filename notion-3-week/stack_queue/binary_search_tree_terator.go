package stack_queue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	queue []int
}

func Constructor3(root *TreeNode) BSTIterator {
	return BSTIterator{
		queue: a(root),
	}
}
func a(r *TreeNode) []int {
	res := make([]int, 0)
	if r == nil {
		return res
	}

	res = append(res, a(r.Right)...)
	res = append(res, r.Val)
	res = append(res, a(r.Left)...)

	return res
}

func (i *BSTIterator) Next() int {
	item := i.queue[len(i.queue)-1]
	i.queue = i.queue[:len(i.queue)-1]
	return item

}

func (i *BSTIterator) HasNext() bool {
	return len(i.queue) != 0
}
