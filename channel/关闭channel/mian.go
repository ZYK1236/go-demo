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
		// 关闭 channel，如果没有关闭，主线程一直在等待管道的值，会死锁
		// 当然一般情况都不需要主动关闭 channel，除非你可以确定这个 channel 不再存储值
		fmt.Println("close channel")
		close(c)
	}()
	
	time.Sleep(2 * time.Second)

	// 如果 ok 为 true，即能够取到管道的值
	for {
		fmt.Println("main running...")
		// 关闭通，把通道中的值都取掉后，ok 返回 false
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main ended, channel's len is 0")
}
