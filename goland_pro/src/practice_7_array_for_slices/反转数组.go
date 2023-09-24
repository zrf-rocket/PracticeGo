package main

import "fmt"

func main() {
	ints := []int{11, 22, 33, 44}
	fmt.Println(ints)
	fmt.Printf("%v\n", ints)

	fmt.Println(len(ints)) //获取数组元素个数

	//	正向遍历数组
	for i := 0; i < len(ints); i++ {
		fmt.Print(ints[i])
		fmt.Print("\t")
	}
	fmt.Println()
	//逆向遍历数组
	for i := len(ints) - 1; i >= 0; i-- {
		fmt.Print(ints[i])
		fmt.Print("\t")
	}

	fmt.Println()
	//反转数组
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
	fmt.Printf("%v\n", ints)
}
