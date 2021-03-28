package main //程序包名字

import (
	"fmt"
	"time" 
)

func main() {
	// 最好不加分号
	time.Sleep(2 * time.Second)
	fmt.Printf("hello go")
}
