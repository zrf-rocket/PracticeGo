package main

import "os"
import "fmt"

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:] //切片操作，取出第一个以后的所有参数
	arg := os.Args[3]              //下标从0开始，此处取的是第4个元素

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
