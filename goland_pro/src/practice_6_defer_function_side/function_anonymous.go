package main

import (
	"fmt"
	"strings"
)

func anonymousFunc() {
	// 定义一个匿名函数
	func() { fmt.Println("this is anonymous function") }() // this is anonymous function
	// 定义完后的匿名函数自动执行，匿名函数定义完后加()则直接执行
	func(num1, num2 int) { fmt.Println(num1 + num2) }(11, 22) // 33

	fmt.Println(func(numbers ...int) []int { return numbers }(1, 2, 3, 4, 5, 6, 7)) // [1 2 3 4 5 6 7]

	// 将匿名函数赋值给变量
	add := func(numbers ...int) int {
		res := 0
		for _, num := range numbers {
			res += num
		}
		return res
	}
	// 通过变量调用匿名函数
	numbers := []int{11, 22, 33, 44}
	// 给函数传递一个slice
	fmt.Println(add(numbers...)) // 110
	// 给函数传递一组数值
	fmt.Println(add(1, 2, 3, 4, 5)) // 15

	// 使用匿名函数对字符串中的每个字符加1
	fmt.Println(strings.Map(func(r rune) rune {
		return r + 1
	}, "abcd1234ABCD")) // bcde2345BCDE
}

// 定义一个返回匿名函数的函数
func anonymousFunc2() func() int {
	var number int
	return func() int {
		number++
		return number * number
	}
}

func main() {
	anonymousFunc()
	aFunc01 := anonymousFunc2()
	fmt.Println(aFunc01()) // 1
	fmt.Println(aFunc01()) // 4
	fmt.Println(aFunc01()) // 9
	fmt.Println(aFunc01()) // 16
}
