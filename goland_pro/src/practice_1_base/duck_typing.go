package main

import "fmt"

type ClassA struct {
	num1, num2, num3 int
}

type ISum interface {
	sum() float64
}

func main() {
	var a ISum
	b := ClassA{1, 2}
	a = b
	a = &b
	fmt.Println(a)
	fmt.Println(b)
}
