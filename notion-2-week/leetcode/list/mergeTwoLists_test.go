package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
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
						Val:  4,
						Next: nil,
					},
				},
			},
			input2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		"2": {
			input1: nil,
			input2: &ListNode{
				Val:  1,
				Next: nil,
			},
			output: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
		"3": {
			input1: nil,
			input2: nil,
			output: nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, mergeTwoLists(tt.input1, tt.input2))
		})
	}
}
