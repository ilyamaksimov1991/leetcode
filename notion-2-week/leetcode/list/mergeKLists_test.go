package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	tests := map[string]struct {
		input1 []*ListNode
		output *ListNode
	}{
		"1": {
			input1: []*ListNode{
				{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
				{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				{
					Val: 2,
					Next: &ListNode{
						Val:  6,
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
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val:  6,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, mergeKLists(tt.input1))
		})
	}
}
