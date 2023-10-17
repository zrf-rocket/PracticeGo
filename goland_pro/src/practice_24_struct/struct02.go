package main

import "fmt"

type MyStruct struct {
	Name string
	Age  int
}

// struct的方法  有两种写法 一种是指针一种是变量
// func (self *ST)f1(){}
// func (self ST)f2(){}

// 对于SayName和SayAge来说，没有区别，声明方式也不造成影响
func (self *MyStruct) SayName() {
	fmt.Println(self.Name)
}

func (self MyStruct) SayAge() {
	fmt.Println(self.Age)
}

// 但是对于SetName和SetAge来说，SetName能够改变，而SetAge却不起作用，声明方式也不造成影响。
func (self *MyStruct) SetName(name string) {
	self.Name = name
}

func (self MyStruct) SetAge(age int) {
	self.Age = age
}

// 区别在于，引用能够改变结构字段的值。
func main() {
	fmt.Println("结构体的使用")
	myStruct := &MyStruct{}
	myStruct.SayName()
	myStruct.SayAge()

	myStruct.SetName("GoLang struct")
	myStruct.SetAge(25)
	fmt.Printf("%+v\n", myStruct)

	myStruct2 := MyStruct{"GoLang2", 22}
	myStruct3 := &MyStruct{"GoLang3", 23}
	fmt.Printf("%+v\n", myStruct2)
	fmt.Printf("%+v\n", myStruct3)
}
