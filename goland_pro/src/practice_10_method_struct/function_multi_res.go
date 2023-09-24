package main

import "fmt"

func sum(num1, num2 int) int {
	return num1 + num2
}

//函数参数为两个int类型的参数   返回值为两个int类型的数值
func reverse(num1, num2 int) (int, int) {
	return num1, num2
}

func main() {
	//将函数返回值赋值给变量res
	res := sum(12, 34)
	fmt.Println(res)

	// multi_res := reverse(111, 333)
	//这种赋值方式无法输出multi_res的值

	fmt.Println(reverse(111, 333))
}
