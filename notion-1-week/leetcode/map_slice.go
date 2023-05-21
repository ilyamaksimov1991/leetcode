package leetcode

import (
	"fmt"
	"unsafe"
)

func main4() {

}

////////////3 TODO почему?
func main5() {
	myStr := st{}

	fmt.Println(myStr) // false false 0
	mutatePtr(&myStr)
	fmt.Println(myStr) // false false 0 TODO почему?
	mutatePtr2(&myStr)
	fmt.Println(myStr) // false true 3 TODO почему?

	println(unsafe.Sizeof(myStr))
}

type st struct {
	p1 bool
	p3 bool
	p2 int
}

func mutatePtr(in *st) {
	in = &st{
		p1: true,
		p2: 22,
		p3: false,
	}
	in.p2 = 3
}
func mutatePtr2(in *st) {
	in.p2 = 3
	in.p3 = true
}

/////////2
func main33() {
	slice := []int{64, 32, 12, 16}
	out := make([]*int, 4)

	for id, val := range slice { // в качестве оптимизации при range  переменные id val инициализируются 1 раз, а потом перезаписываются
		out[id] = &val
	}

	for id, val := range out {
		println(id, *val) // 16 16 16 16
	}
}

/////////1 TODO НЕ ПОНЯЛ

func main3() {
	slice := make([]int, 0, 4)
	slice = append(slice, 1)
	slice = append(slice, 2) //1,2_,_
	fmt.Println(slice)       //1,2

	append1(slice, 3)  //1,2,3,_
	fmt.Println(slice) //1,2
	fmt.Println(slice[:3])

	//mutate1(slice, 3, 4) // panic
	//fmt.Println(slice)

}

func append1(s []int, i int) {
	s = append(s, i)
}

func mutate1(s []int, key int, val int) {
	s[key] = val
}
