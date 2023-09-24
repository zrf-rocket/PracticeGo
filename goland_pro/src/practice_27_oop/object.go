package main

import "fmt"

//直接使用结构  无构造函数
type ClassA struct {
	num1, num2, num3 int
}

//没有self this之类的内容
func (classa ClassA) sum() int {
	return classa.num1 + classa.num2
}

func main() {
	fmt.Println(ClassA{11, 22, 33}.sum())
}
