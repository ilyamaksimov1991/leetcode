package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseList(t *testing.T) {
	tests := map[string]struct {
		input1 *ListNode
		input2 *ListNode
		output *ListNode
	}{
		"1": {
			input1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			output: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, reverseList(tt.input1))
		})
	}
}
