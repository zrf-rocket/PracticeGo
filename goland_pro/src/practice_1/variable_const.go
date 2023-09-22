package main

import "fmt"

/*
1. 定义的时候，必须指定值
2. 指定的值类型主要有三类： 布尔，数字，字符串， 其中数字类型包含（rune, integer, floating-point, complex), 它们都属于基本数据类型。
3. 不能使用 :=
*/
//const NUM := 0  // syntax error: unexpected :=, expected =
func constVar() {
	const maxAge int = 100
	const PI = 3.1415926

	fmt.Println(maxAge)
	//maxAge = 0  // cannot assign to maxAge (constant 100 of type int)

	const num = 1234
	// num = 123 //erro   cannot assign to num (untyped int constant 1234)

	const (
		num1 = 1000
		num2 = 2.222
	)
	// num1 = 123 // error   cannot assign to num1
	fmt.Println(num, num1+num2)
}

// 全局常量 多个常量一起声明
const (
	BLOG   = "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"
	WECHAT = "CTO Plus"
)

// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	maxAge = 100
	age2
	age3
)

//const (
//	num1 // missing init expr for num1
//	num2
//	num3
//)

func main() {
	constVar()
	fmt.Println(BLOG)               // https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q
	fmt.Println(WECHAT)             // CTO Plus
	fmt.Println(maxAge, age2, age3) // CTO Plus
	fmt.Println(num1, num2, num3)
}
