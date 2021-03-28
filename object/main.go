package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (this Person) getName() string {
	return this.Name
}

func (this Person) errorSetName(newName string) {
	// this 是一个对象的拷贝...和 js 又不一样，所以修改不了原对象属性 
	this.Name = newName
}

func (this *Person) rightSetName(newName string) {
	// 用指针
	this.Name = newName
}

func newPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	person := newPerson("zyk", 18)
	fmt.Println("person:", person)
	person.errorSetName("mxw1")
	fmt.Println("person:", person)
	person.rightSetName("mxw2")
	fmt.Println("person:", person)
}
