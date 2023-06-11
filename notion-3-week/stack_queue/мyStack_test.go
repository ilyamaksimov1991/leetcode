package stack_queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor1(t *testing.T) {
	myStack := Constructor1()

	myStack.Push(1)
	myStack.Push(2)
	assert.Equal(t, 2, myStack.Top())
	assert.Equal(t, 2, myStack.Pop())
	assert.Equal(t, false, myStack.Empty())
}
