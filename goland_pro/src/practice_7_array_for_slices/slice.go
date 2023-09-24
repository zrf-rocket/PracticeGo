package main

import (
	"fmt"
	"os"
)

func main() {
	target := "Hello SteveRocket"
	// :=  该符号声明并初始化了一个字符串变量。Go语言虽然是一门静态强类型的语言，
	//但是在使用:=操作符时，Go语言会根据符号右边的值推导出符号左边变量的类型。
	fmt.Println(target)

	if len(os.Args) > 1 {
		//os.Args是一个字符串切片(slice) ，切片内容是传递给Go程序的参数
		target = os.Args[1]
		fmt.Println(len(target))
	}
	// 切片是一个可以动态增长的数组，可以通过len()内置函数计算切片的长度，通过cap计算切片的容量
	// 通过slice[n]的方式访问切片中的第n个元素，而slice[n:]则返回从第n个元素到最后一个元素的切片，像极了Python中的切片。
	fmt.Println("Hello:", target)
}
