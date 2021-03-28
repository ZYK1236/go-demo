package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		// runtime.Goexit() 退出协程
		fmt.Println("go routine")
		channel <- 666
	}()

	data := <-channel

	fmt.Println("main", data)
}
