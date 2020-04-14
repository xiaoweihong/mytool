package main

import (
	"fmt"
)

func main() {
	var nums []int
	var s1 = make(chan int)

	nums = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("nums 切片", nums)

	go func() {
		for _, v := range nums {
			fmt.Println("开始发送到channel", v)
			s1 <- v
		}
		close(s1)
	}()

	for {
		ok, v := <-s1
		if !v {
			break
		}
		fmt.Println("接收-->", ok)
	}
	//time.Sleep(time.Second)
}
