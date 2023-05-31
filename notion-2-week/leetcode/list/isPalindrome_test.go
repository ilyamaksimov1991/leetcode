package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := map[string]struct {
		input1 *ListNode
		output bool
	}{
		"1": {
			input1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			output: true,
		},
		"2": {
			input1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			output: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, isPalindrome(tt.input1))
		})
	}
}
