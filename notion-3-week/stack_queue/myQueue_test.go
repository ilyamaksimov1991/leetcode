package stack_queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	myQueue := Constructor()
	myQueue.Push(1)                         // queue is: [1]
	myQueue.Push(2)                         // queue is: [1, 2] (leftmost is front of the queue)
	assert.Equal(t, 1, myQueue.Peek())      // return 1
	assert.Equal(t, 1, myQueue.Pop())       // return 1, queue is [2]
	assert.Equal(t, false, myQueue.Empty()) // return false
}
