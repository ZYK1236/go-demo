// select 可以监控多个 channel 的读写状态
package main

import "fmt"

// 1,1,2,4,8,16,
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// c 可写，就进入数列计算
		case c <- x:
			temp := x
			x = y
			y = temp + y
		// c 不可写，退出函数
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	/*
	 * select {
	 * 	case <- channel_1:
	 *  如果 channel_1 成功读到数据，执行操作
	 * }
	 * 	case <- channel_2:
	 *  如果 channel_2 成功读到数据，执行操作
	 * }
	 * 	default:
	 *  如果 上面都没有成功，执行 default 操作
	 * }
	 */

	// 斐波拉切数列
	c := make(chan int)
	quit := make(chan int)

	go func() {
		// 这边就是控制 c 可以写多少次
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	fibonacci(c, quit)
}
