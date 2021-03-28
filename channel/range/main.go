package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("send -> channel", i)
		}
		fmt.Println("close channel")
		close(c)
	}()

	time.Sleep(2 * time.Second)

	// 可以用 range 来不断迭代 channel，它会智能结束循环
	for data := range c {
		fmt.Println("get data from channel:", data)
	}

	fmt.Println("main ended, channel's len is 0")
}
