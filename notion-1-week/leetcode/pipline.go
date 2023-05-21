package leetcode

import (
	"fmt"
	"sync"
)

func main2() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{20, 10, 30} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{300, 200, 100} {
			c <- num
		}
		close(c)
	}()

	for num := range joinChannels(a, b, c) {
		fmt.Println(num)
	}
}

func joinChannels(chs ...chan int) <-chan int {
	res := make(chan int)
	go func() {
		for _, val := range chs {
			val := val
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				for val2 := range val {
					res <- val2
				}
			}()
			wg.Wait()

		}

		close(res)
	}()

	return res
}
