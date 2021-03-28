// interface{} 引用任何的类型，类似于 ts 中的 any

package main

import "fmt"

type Book struct {
	name string
}

func demo(arg interface{}) {
	// 有个问题，如何区分万能类型：使用断言
	value, ok := arg.(string)

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("no string")
	}
}

func main() {
	book := Book{"zyk"}
	demo(1)
	demo("1")
	demo(book)
}
