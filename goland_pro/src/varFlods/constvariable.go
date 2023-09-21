package main

import "fmt"

/*
1. 定义的时候，必须指定值
2. 指定的值类型主要有三类： 布尔，数字，字符串， 其中数字类型包含（rune, integer, floating-point, complex), 它们都属于基本数据类型。
3. 不能使用 :=
*/
func constVar() {
	const num = 1234
	// num = 123 //erro   cannot assign to num (untyped int constant 1234)

	const (
		num1 = 1000
		num2 = 2.222
	)
	// num1 = 123 // error   cannot assign to num1
	fmt.Println(num, num1+num2)
}
