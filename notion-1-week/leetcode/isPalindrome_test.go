package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := map[string]struct {
		input1 string
		result bool
	}{
		"0": {
			input1: "A man, a plan, a canal: Panama",
			result: true,
		},
		"1": {
			input1: "race a car",
			result: false,
		},
		"2": {
			input1: "",
			result: true,
		},
		"3": {
			input1: "as",
			result: false,
		},
		"4": {
			input1: "0P",
			result: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.result, isPalindrome(tt.input1))
		})
	}
}

func Test_isPalindrome2(t *testing.T) {
	input1 := "A man, a plan, a canal: Panama"

	for i := 0; i < len(input1); i++ {
		println(string(input1[i]), isCharNum(input1[i]), input1[i])
	}

}
