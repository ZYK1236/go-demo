package main

import "fmt"

func demo1(a int, b int) (int, string) {
	fmt.Println("------demo1-----")
	c := a + b
	d := "n"
	return c, d
}

func demo2(a int, b int) ()

func main() {
	res, res2 := demo1(1, 2)

	print(res, res2)
}
