package main

import "fmt"

// 声明全局变量
var GlobalA int = 1000

// := 语法只能用在函数体内

func main() {
	/*
		四种方式
	*/

	// 默认 0
	var a int
	fmt.Print(a, "\n")

	// 初始化
	var b int = 100
	fmt.Print(b, "\n")

	// 数据自动检测
	var c = 1000
	fmt.Print(c, "\n")

	// 类型检测
	fmt.Printf("%T\n", c)

	var str = "123"
	fmt.Printf("%s, %T\n", str, str)

	// 最常用的定义变量方法
	e := 10.2
	fmt.Printf("%T\n", e)
 
	print(GlobalA)

	// 声明多个变量
	var aa, bb int = 1, 2
	var zyk, mxw = 1, '1'
	print(aa, bb)
	print(zyk, mxw,"\n")

	// 声明多个变量 第二种方法
	var (
		npm int = 10
		node string = "10"
	)

	print(npm, node)
}
