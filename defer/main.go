package main

func printReturn() int{
	println("return")

	return 0
}

func deferDemo() int{
	defer println("defer")

	return printReturn()
}

func main() {
	// defer: 函数/流程结束之前触发，其实就是函数生命周期结束后调用
	// defer 是一种类压栈的形式，所以先进后出，end2 先输出，再输出 end1

	// defer println("end1")
	// defer println("end2")
	// println(1)
	// println(2)

	// 先执行函数内的 return 语句，再 defer
	deferDemo()
}
