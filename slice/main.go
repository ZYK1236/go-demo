package main

import "fmt"

func modifyArray(arr [4]int) {
	// 这样子并不能修改原先数组的值，因为是一个拷贝
	arr[0] = 100
}

func main() {
	// 这种声明是数组
	var arr = [4]int{1, 2, 3, 4}

	// 第一种遍历方式
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 第二种遍历方式
	for index, value := range arr {
		fmt.Println("index, value", index, value)
	}

	modifyArray(arr)

	fmt.Println("-=-==-=-=-")
	
	for index, value := range arr {
		fmt.Println("index, value", index, value)
	}

}
