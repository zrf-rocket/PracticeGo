package main

import "fmt"

// 给int8类型起一个别名   类似于C中的typedef
type Gender int8

const (
	MALE   Gender = iota
	FEMALE Gender = iota
)

func main() {
	fmt.Println(MALE, FEMALE) // 0  1
	gender := MALE

	switch gender {
	case FEMALE:
		fmt.Println("female")
	case MALE:
		fmt.Println("male")
	default:
		fmt.Println("unknown")
	}
}
