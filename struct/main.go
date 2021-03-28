package main

import "fmt"

// 如果类首字母大写，是可以被外部包访问的(public)，和函数一样
type Book struct {
	name string
	author string
	publishTime int64
}

func logBook(book Book) {
	// 传递的是拷贝
	book.name = "mxw"
}

func logBookPointer(book *Book) {
	// 传递的是指针
	book.name = "mxw"
}

func main() {
	book := Book {
		name: "zyk",
		author: "zyk",
		publishTime: 1223123,
	}

	fmt.Println(book)
	logBook(book)
	fmt.Println(book)
	logBookPointer(&book)
	fmt.Println(book)
}