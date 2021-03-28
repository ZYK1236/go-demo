package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"007", 2000, []string{"zyk", "mxw"}}
	// 结构体 -> jsonStr
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("json = %s", jsonStr)
	}
}
