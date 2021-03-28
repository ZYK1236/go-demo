package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个 channel
	channel := make(chan int)

	go func() {
		defer fmt.Println("go func ended -=-=-=-=-=-=-=-=-=-=-")
		// 向 channel 发送数据
		fmt.Println("channel 即将存数据")
		channel <- 666
		fmt.Println("channel <- 666 后的语句")
	}()
	
	time.Sleep(2 * time.Second)
	// 从 channel 取数据
	// channel 有一种 await 的功能，阻塞直到等到 channel 接收到值
	// 这种阻塞不仅发生在 main 中，函数中也会阻塞（main 不消耗，func 一直等待）
	fmt.Println("channel 即将拿数据")

	num := <-channel
	fmt.Println("channel 拿到数据 num =", num)

	fmt.Println("main goroutine 结束")
}
