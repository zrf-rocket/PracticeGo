package main

import "fmt"

func variable() {
	var num int      // 如果没有赋值，默认为0
	fmt.Println(num) // 0

	var num2 int = 123 // 声明时赋值
	fmt.Println(num2)  // 123

	//类型推导 因为 456 是 int 类型，所以赋值时，num3 自动被确定为 int 类型，所以类型名可以省略不写
	var num3 = 456    // 声明时赋值
	fmt.Println(num3) // 456

	var nick_name string   // 声明一个保存字符串类型的变量s1
	fmt.Println(nick_name) // 空字符串

	// 更简单的声明赋值   简短变量声明，只能再函数里面用
	name := "SteveRocket"
	fmt.Println(name) // SteveRocket

	var isNull bool // 默认为false
	fmt.Println(isNull)
}

func myFunction(age int) int {
	var count int
	return count + age
}

func simple_variable() {
	var age int = 28
	var name string = "SteveRocket"

	fmt.Println("name:", name, "age:", age) // name: SteveRocket age: 28
	fmt.Println(myFunction(age))            // 28
}

func main() {
	//variable()
	simple_variable()
}
