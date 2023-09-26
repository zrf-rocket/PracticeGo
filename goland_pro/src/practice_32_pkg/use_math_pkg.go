package main

import (
	"fmt"
	"practice_32_pkg/mymaths"
)

// 调用
func FunctionTransfer1() {
	//mymaths.Sub调用自定义的包中的方法
	res := mymaths.Sub(11, 11)
	fmt.Println(res)
}

func main() {
	FunctionTransfer1()
	mymaths.ShowFunc() // this is my maths showFunc
}
