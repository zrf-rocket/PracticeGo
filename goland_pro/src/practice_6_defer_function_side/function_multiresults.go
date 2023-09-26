package main

import (
	"fmt"
	"reflect"
)

// 同一种类型返回值
func multiReturnValue() (int, int) {
	// 像这种仅仅只有返回值类型 而没有变量名的函数，对于代码可读性不是很友好，特别是在同类型的返回值出现时，无法区分每个返回参数的意义。
	num1 := 111
	num2 := 222
	return num1, num2
	// 等价于下面
	return 111, 222
}

// 带有变量名的返回值
func multiReturnValue02() (num1 int, num2 int) {
	//num1 = 111
	//num2 = 222

	num2 = 222
	num1 = 111
	return
}

func multiReturnValue03() (num1, num2 int) {
	num2 = 200
	num1 = 100
	return
}

func multiReturnValue04() (int, interface{}, []int) {
	return 1, reflect.Interface, []int{11, 22, 33}
}

func multiReturnValue05() (num1, num2, num3 int, num5 int) {
	return 11, 22, 33, 44
}

/*
func multiReturnValue06() (num1, num2, num3 int, int) {  //  syntax error: mixed named and unnamed parameters
	return 11, 22, 33, 44
}
*/

func main() {
	fmt.Println(multiReturnValue())   // 111 222
	fmt.Println(multiReturnValue02()) // 111 222
	fmt.Println(multiReturnValue03()) // 100 200
	fmt.Println(multiReturnValue04()) // 1 interface  [11 22 33]
	fmt.Println(multiReturnValue05()) // 11 22 33 44
}
