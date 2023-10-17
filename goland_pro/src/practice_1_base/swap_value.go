package main

import "fmt"

func main() {
	num1, num2 := 11, 22
	fmt.Println(num1, num2) // 11 22
	num1, num2 = num2, num1
	fmt.Println(num1, num2) // 22 11
}
