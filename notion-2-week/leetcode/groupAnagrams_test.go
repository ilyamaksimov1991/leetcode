package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tests := map[string]struct {
		input  []string
		output [][]string
	}{
		"1": {
			input:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			output: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		"2": {
			input:  []string{""},
			output: [][]string{{""}},
		},
		"3": {
			input:  []string{"a"},
			output: [][]string{{"a"}},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, groupAnagrams(tt.input))
		})
	}
}
