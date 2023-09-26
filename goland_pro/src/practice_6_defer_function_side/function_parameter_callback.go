package main

import (
	"fmt"
)

func main() {
	callback(1, Add)

	//函数可以被赋值给变量
	funcVar := func01
	fmt.Printf("类型为：%T   返回值为：%d  %d\n", funcVar, func01(), funcVar()) // 类型为：func() int   返回值为：1000  1000
	//函数作为参数传递给函数
	funcParamter(funcVar)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b) // The sum of 1 and 2 is: 3
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

func func01() int {
	return 1000
}

// 将函数作为函数参数
func funcParamter(function func() int) {
	result := function()
	fmt.Printf("值为：%v  类型为：%T  参数类型为：%T", result, result, function) // 值为：1000  类型为：int  参数类型为：func() int
}
