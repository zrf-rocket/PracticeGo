package main

import "fmt"

// 全局变量studentName  go语言中推荐使用驼峰命名
var studentName string

// 批量声明全局变量
// 声明全局变量建议使用批量声明，方便追加
var (
	num1   int
	name   string
	age    int32
	isNull bool
	salary float32
)

// 声明变量
var myName string
var myAge int
var isSex bool

// 批量初始化变量(一次初始化多个变量)
var name2, age2, wechat = "SteveRocket", 28, "CTO Plus"

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

func variable02() {
	// 非全局变量（局部变量）声明了就必须使用，不使用就无法编译通过（全局变量声明或初始化了不使用也可以编译通过）
	// 不使用报错：blog declared and not used
	var blog = "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"
	fmt.Println(blog)

	//同一个作用域，不能重复声明同名变量
	//blog := "" // no new variables on left side of :=
	blog = "" // 但可以重新赋值
	fmt.Println(blog)

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
	// 匿名变量 是一个特殊的变量： _
	// 匿名变量使用_指定，用来作为一个占位符，匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
	return 28, "SteveRocket", true, salary
}

func main() {
	variable()
	variable02()
	variable_default_value()
	age, _, sex, _ := annoymous_variable(0)
	fmt.Println(age, sex) // 28 true
	//_, _, _, _ := annoymous_variable(0)  // no new variables on left side of :=
	_, _, _, salary := annoymous_variable(0)
	fmt.Println(salary)
}
