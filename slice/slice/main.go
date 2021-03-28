package main

import "fmt"

func main() {
	// 定义一个 slice / 动态数组，并初始化
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 定义一个 slice，不初始化
	var slice2 []int
	// slice2[0] = 1 会报错，因为长度为 0，没有多余的空间，所以我们要开辟空间
	// 开辟的空间默认值为0
	slice2 = make([]int, 3) // [0, 0, 0]
	fmt.Println("slice2:", slice2)

	// 简写:
	slice3 := make([]int, 0)
	fmt.Println("slice3:", slice3)

	// 判空
	var slice4 []int
	fmt.Println("slice4:", slice4)

	fmt.Printf("len3 = %d, len4 = %d\n", len(slice3), len(slice4))

	// 虽然 slice3通过 make 开辟 0个空间，但它还是有空间的，不为空
	if(slice3 == nil) {
		fmt.Println("slice3为空")
	} else {
		fmt.Println("slice3不为空")
	}

	if(slice4 == nil) {
		fmt.Println("slice4为空")
	} else {
		fmt.Println("slice4不为空")
	}
}
