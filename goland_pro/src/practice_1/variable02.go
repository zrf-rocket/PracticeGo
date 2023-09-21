package main

import "fmt"

// 全局变量studentName  go语言中推荐使用驼峰命名
var studentName string

// 批量声明全局变量
var (
	num1   int
	name   string
	age    int32
	isNull bool
	salary float32
)

// 批量初始化变量
var name2, age2, wechat = "SteveRocket", 28, "CTO Plus"

//或者一次初始化多个变量
//var name, age = "Q1mi", 20

func variable() {
	Age := 30        // 变量区分大小写
	fmt.Println(Age) // 30

	age := 28
	fmt.Println(age) // 28
	age = 30         // 赋值
	fmt.Println(age) //30
	age = 25.0       // 重新赋值
	fmt.Println(age) // 25

	//在赋值时，需要注意变量的类型必须与所赋的值的类型一致，否则会导致编译错误。
	//age = 25.55  // cannot use 25.55 (untyped float constant) as int value in assignment (truncated)

	fmt.Println(age, Age) // 25 30
	age, Age = Age, age
	fmt.Println(age, Age) // 25 30
}

func variable_default_value() {
	var name string
	var age int
	var isNull bool
	var salary float64
	// 切片
	var slice []int
	// 函数
	var fn func()
	// 指针变量
	var ptr *int
	fmt.Println(name, age, isNull, salary, slice, fn, ptr) // "" 0 false 0 [] <nil> <nil>
}

func annoymous_variable(salary float64) (int, string, bool, float64) {
	return 28, "SteveRocket", true, salary
}
func main() {
	//variable()
	//variable_default_value()
	age, _, sex, _ := annoymous_variable(0)
	fmt.Println(age, sex) // 28 true
	//_, _, _, _ := annoymous_variable(0)  // no new variables on left side of :=
	_, _, _, salary := annoymous_variable(0)
	fmt.Println(salary)
}
