package hashTables

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	tests := map[string]struct {
		input  string
		output int
	}{
		"0": {
			input:  "leetcode",
			output: 0,
		},
		"1": {
			input:  "loveleetcode",
			output: 2,
		},
		"2": {
			input:  "aabb",
			output: -1,
		},
		"3": {
			input:  "aadadaad",
			output: -1,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, firstUniqChar(tt.input))
		})
	}
}
