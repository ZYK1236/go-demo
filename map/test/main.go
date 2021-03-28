package main

import "fmt"

func logMap(cityMap map[string]string) {
	// cityMap 是一个引用传递，和 js 一样
	for key, value := range cityMap {
		fmt.Println("key", key, " value:", value)
	}

	// 原 map 会被影响到
	cityMap["demo"] = "ZYK"

	fmt.Println("-==--=--=-=-=--")
}

func main() {
	cityMap := make(map[string]string)

	cityMap["CHINA"] = "NO_1"
	cityMap["AMERICA"] = "NO_100"
	cityMap["INDIA"] = "NO_-1"

	logMap(cityMap)

	delete(cityMap, "INDIA")
	cityMap["AMERICA"] = "NO_-100"

	logMap(cityMap)
}
