package merge

import (
	"fmt"
	"testing"
)

func Test_mergesort(t *testing.T) {
	lst := []int{2, 123, 443, 1, 223, 3, 222, 5, 6, 432, 1}
	lst = mergesort(lst)
	fmt.Println(lst)
}
