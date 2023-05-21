package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_convertToTitle2(t *testing.T) {
	b := 701

	res := []rune{}
	for {
		residual := b % 26
		if residual != 0 {
			res = append(res, rune(residual)+'A'-1)
			b = b - residual
		}
		b = b / 26

		if b <= 26 {
			res = append(res, rune((b)+'A'-1))
			break
		}
	}

	res1 := []rune{}
	for i := len(res) - 1; i >= 0; i-- {
		res1 = append(res1, res[i])
	}

	println(string(res1))
}

func Test_convertToTitle(t *testing.T) {
	tests := map[string]struct {
		input  int
		output string
	}{
		"0": {
			output: "Z",
			input:  26,
		},
		"02": {
			output: "A",
			input:  1,
		},
		"AZ": {
			output: "AZ",
			input:  52,
		},
		"01": {
			output: "AA",
			input:  27,
		},
		"3": {
			output: "AAA",
			input:  703, //676(26+26)+27,
		},
		"4": {
			output: "AAAA",
			input:  18279, //17576(26*26*26)+703
		},
		"2": {
			output: "ZY",
			input:  701,
		},
		"1": {
			output: "AB",
			input:  28,
		},
		"11": {
			output: "BB", // (2*26+2),
			input:  54,
		},
		"5": {
			output: "BBB", // номер_символа*26в_кв
			input:  1406,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, convertToTitle(tt.input))
		})
	}
}
