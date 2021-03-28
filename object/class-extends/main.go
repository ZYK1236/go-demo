package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person //继承

	Grade int
}

func (this *Person) eat() {
	fmt.Println("eat...")
}

func (this *Person) sleep() {
	fmt.Println("sleep...")
}

// 重写父类方法
func (this *Student) sleep() string {
	return ("student don't need sleep...")
}

func createStudent(name string, age int) Student {
	return Student{
		Person{name, age},
		100,
	}
}

func main() {
	student := createStudent("zyk", 18)
	fmt.Println("studentName:", student.Name, "\n",
	"studentAge:", student.Age, "\n",
	"student to sleep:", student.sleep())
}
