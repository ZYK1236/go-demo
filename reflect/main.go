package main

import (
	"fmt"
	"reflect"
)

func getType(arg interface{}) {
	user, ok := arg.(User)
	if ok {
		// 获取结构体里面的属性/值
		myType := reflect.TypeOf(user)
		myValue := reflect.ValueOf(user)
		for i := 0; i < myType.NumField(); i++ {
			field := myType.Field(i)
			value := myValue.Field(i).Interface()

			fmt.Printf("%s %v = %v\n", field.Name, field.Type, value)
		}
		// 获取结构体的方法
		myMethod := myType.NumMethod()
		for i := 0; i < myMethod; i++ {
			method := myType.Method(i)
			fmt.Printf("%s: %v", method.Name, method.Type)
		}

	} else {
		myType := reflect.TypeOf(arg)
		fmt.Println(myType.Name())
	}
}

type User struct {
	Id   int
	Name string
}

func (this User) Call() {

}

func main() {
	a := 1.23
	b := "123"
	c := User{1, "zyk"}
	fmt.Println("ni hao")
	getType(a)
	getType(b)
	getType(c)
}
