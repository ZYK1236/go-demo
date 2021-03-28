package main

import "fmt"

// 是一个指针
type Animal interface {
	GetColor()
}

// 具体的类，如果实现了接口的方法
type Cat struct {
	Color string
}

type Dog struct {
	Color string
}

func (this *Cat) GetColor() {
	fmt.Println(this.Color)
}

func (this *Dog) GetColor() {
	fmt.Println(this.Color)
}

func showAnimal(animal Animal) {
	animal.GetColor()
}

func main() {
	// 声明一个接口
	cat := Cat{"yellow"}
	dog := Dog{"blue"}

	// interface 是一个指针，所以传地址
	showAnimal(&cat)
	showAnimal(&dog)
}
