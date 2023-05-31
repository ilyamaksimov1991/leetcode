package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	tests := map[string]struct {
		input1 *ListNode
		output bool
	}{
		"1": {
			input1: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val:  -4,
							Next: nil,
						},
					},
				},
			},
			output: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, hasCycle(tt.input1))
		})
	}
}
