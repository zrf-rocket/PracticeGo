package main

import (
	"fmt"
	"strings"
)

func func01(num int) int { return -num }
func func03(num1, num2 int) int {
	return num1 + num2
}

var func04 func(int) int

func funcValue() {
	fValue := func01
	fmt.Println(fValue) // 0xdade20

	//fValue = func03     // cannot use func03 (value of type func(num1 int, num2 int) int) as func(num int) int value in assignment

	fValue1 := func03
	fmt.Println(fValue1) // 0xdade40

	//fmt.Println(func04(123)) // 此处func04的值为nil, 会引起panic错误
	//	panic: runtime error: invalid memory address or nil pointer dereference

	// 函数值可以与nil比较
	if func04 != nil {
		fmt.Println(func04(123))
	}
}

func concatStr(c rune) rune { return c + 1 }

func main() {
	funcValue()

	fmt.Println(strings.Map(concatStr, "CTO Plus"))     // DUP!Qmvt
	fmt.Println(strings.Map(concatStr, "1234abcdABCD")) // 2345bcdeBCDE
	fmt.Println(strings.Map(concatStr, "中文"))           //丮斈
}
