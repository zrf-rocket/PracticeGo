package main

import "fmt"

func main() {
	a := 1
	goto TARGET // 执行报错 compile error    goto TARGET jumps over declaration of b a
	b := 9
TARGET:
	b += a
	fmt.Printf("a is %v *** b is %v", a, b)
}
