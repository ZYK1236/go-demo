package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		// runtime.Goexit() ιεΊεη¨
		fmt.Println("go routine")
		channel <- 666
	}()

	data := <-channel

	fmt.Println("main", data)
}
