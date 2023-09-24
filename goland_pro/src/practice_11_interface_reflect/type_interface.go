package main

import "fmt"

// 任意类型：interface{}
func sum2(cnt1, cnt2 interface{}) interface{} {
	return cnt1 + cnt2
}

func main() {
	fmt.Println(sum2(11, 22))
	fmt.Println(sum2(1.1, 2.2))
	fmt.Println(sum2("Hello", "Golang"))
}
