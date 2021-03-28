package main

import (
	"fmt"
	"go-demo/pointer/swap"
)

// 通过指针在函数内部修改值影响外部
func pointer(p *int) {
	// *就是改变指针指向的地址的值
	*p = 10
}

func main() {
	a := 5
	// 传递地址，而不是值的拷贝
	pointer(&a)
	// 打印 10
	println("a =", a)

	/*
	* @description: 下面是交换两个数字的例子
	 */

	b := 10
	c := 20
	fmt.Printf("previous: b = %d, c = %d\n", b, c)
	swap.Swap(&b, &c)
	fmt.Printf("after swap: b = %d, c = %d\n", b, c)

	/*
	* @description: 一级 / 二级指针例子
	 */

	e := 30
	var f *int = &e
	// f := &e

	// 0xc00006cf08 0xc00006cf08
	println(&e, f)

	var g **int = &f
	// g := &f

	// 0xc00006cf28 0xc00006cf28
	println(&f, g)
}
