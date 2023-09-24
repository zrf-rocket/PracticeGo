package main

import "fmt"

func main() {
	// 从0遍历到4
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}

	//break 和 continue 的用法与其他语言没有区别
	sum := 0
	for i := 0; i < 10; i++ {
		if sum > 50 {
			break
		}
		sum += i
	}

	sum02 := 0
	for sum02 < 10 {
		fmt.Println(sum02)
		//sum02 += 1
		sum02++
	}
}
