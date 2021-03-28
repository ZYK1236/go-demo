package main

import (
	"go-demo/init_import/package"
	"go-demo/init_import/package2"
)

func main()  {
	// 先调用每个包的 init 方法，然后再调用方法
	// init 是一个关键字
	pkg1.Demo1()
	pkg2.Demo2()
}