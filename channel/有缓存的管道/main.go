package main

import (
	"fmt"
	"time"
)

func main() {
	// 有缓冲的 channel
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("go func ended")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("发送给管道的元素:", i, "此时管道len:", len(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("接收管道的元素:", num)
	}

	fmt.Println("main ended")
}
