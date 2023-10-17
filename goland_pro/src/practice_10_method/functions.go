package main

import "fmt"

func plus(num1 int, num2 int) int { //接受两个int类型参数并将它们的和作为一个int返回。
	return num1 + num2
}

func plusPlus(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

/*
函数是Go语言编程的核心
*/
func main() {

	res := plus(12, 34)
	fmt.Println(res)

	res = plusPlus(11, 22, 33)
	fmt.Println(res)
}
