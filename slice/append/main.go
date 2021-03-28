package main

import "fmt"

func main() {
	slice := make([]int, 0)
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}
	fmt.Println("slice = ", slice, "len = ", len(slice), "cap = ", cap(slice))
	fmt.Println("截取切片:slice(1, 3): ", slice[1:3])

	// 注意，和 js 不一样，截取得到的切片并不是一个拷贝，修改截取得到的切片原切片会受到影响
	// 如果想达到和 js 一样的效果，可以用 copy 函数，注意 dst slice 的 len 要和 source slice 的长度相等

	tempSlice := make([]int, 2)
	copy(tempSlice, slice[1:3])
	fmt.Println("copySlice:", tempSlice)

	fmt.Println("-=-=-=-=-=-=-=-=-")
	
	// 这种形式定义的 slice 也可以动态 append
	slice2 := []int{1, 2, 3}
	for i := 0; i < 3; i++ {
		slice2 = append(slice2, i)
	}
	fmt.Println("slice2 = ", slice2, "len = ", len(slice2), "cap = ", cap(slice2))
}
