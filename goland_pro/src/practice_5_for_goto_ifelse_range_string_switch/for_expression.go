package main

import "fmt"

/*
for语句是Go编程语言中唯一的循环结构
1、最基本的类型，只具有单一条件。
2、一个经典的初始化/条件/之后的for循环。
3、对于没有条件的函数将重复循环，直到跳出循环或从包闭函数中返回。
*/
func main() {

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break //结束当前循环
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue //跳出本次循环，继续下面的循环
		}
		fmt.Println(n)
	}
}
