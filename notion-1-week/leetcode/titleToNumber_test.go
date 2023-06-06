package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	tests := map[string]struct {
		input string
		res   int
	}{
		"0": {
			input: "A",
			res:   1,
		},
		"01": {
			input: "AA",
			res:   27,
		},
		"3": {
			input: "AAA",
			res:   703, //676(26+26)+27,
		},
		"4": {
			input: "AAAA",
			res:   18279, //17576(26*26*26)+703
		},
		"2": {
			input: "ZY",
			res:   701,
		},
		"1": {
			input: "AB",
			res:   28,
		},
		"11": {
			input: "BB", // (2*26+2),
			res:   54,
		},
		"5": {
			input: "BBB", // номер_символа*26в_кв
			res:   1406,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.res, titleToNumber(tt.input))
		})
	}
}
