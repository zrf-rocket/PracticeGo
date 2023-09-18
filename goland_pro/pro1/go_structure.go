package main

import (
	"fmt"
	"unsafe"
)

// 声明结构体
type name_struct struct {
}

// 声明接口
type name_inteface interface {
}

func main() {
	fmt.Println(unsafe.Sizeof(name_struct{})) // 0

	fmt.Println("my name is SteveRocket") // my name is SteveRocket

	age := 28
	var name = "SteveRocket"
	fmt.Println(age, name) // 28 SteveRocket
	name = "Golang"
	fmt.Println(name) // Golang

	const blog = "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"
	fmt.Println("this is my blog:", blog) // this is my blog: https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q
	//blog = ""                             // const定义的常量无法再次进行修改   cannot assign to blog (untyped string constant
}
