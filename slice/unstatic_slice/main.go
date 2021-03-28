package main

import "fmt"

func logArr(arr [] int) {
	for _, value := range arr {
		fmt.Println(value)
	}
	// 动态数组直接修改下标会影响到原数组，和 js 一样
	arr[0] = 10
}

func main() {
	myArray := []int{1, 2, 3, 4} //动态数组
	logArr(myArray)
	fmt.Println(myArray)
}
