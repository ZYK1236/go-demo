package main

import (
	"fmt"
	// 如果想使用包的 init，而不使用它的方法，用匿名别名包，不使用它的方法不报错
	// 如果把 _ 改成其他名字，就是重命名包了 
	_ "go-demo/_import/package"
)

func main()  {
	fmt.Print("666\n")
}