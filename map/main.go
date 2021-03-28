package main

import "fmt"

func main() {
	myMap := make(map[int]string)

	for i := 0; i < 3; i++ {
		str := fmt.Sprintf("%d", i)
		myMap[i] = str
	}

	fmt.Println(myMap)
}