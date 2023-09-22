package main

import "fmt"

// 变量作用域
var age int

func scopeFunc() {
	fmt.Println(age) // 全局变量age 0
	age = 30
	fmt.Println(age) // 全局变量age 30
	var age = 28
	fmt.Println(age) // 28
}

func scopeFunc2() {
	fmt.Println(age) // 此处使用到的是全局变量age  30

	// 函数作用域 定义局部变量
	var name = "steverocket"
	blog := "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"
	fmt.Println(name) // steverocket
	fmt.Println(blog) // https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q
} // 程序执行到此处  局部变量就会被释放

func scopeBlock() {
	block1 := 123
	block2 := 555
	{
		block1 := 456
		{
			block1 := 789
			block3 := 333
			{
				block1 := 999
				fmt.Println(block1) // 999
				fmt.Println(block2) // 使用最外层的块作用域  555
			}
			fmt.Println(block1) // 789
			fmt.Println(block3)
		}
		//fmt.Println(block3) // undefined: block3  外层块无法访问内层块作用域的变量
		fmt.Println(block1) // 456
	}
	fmt.Println(block1) // 123
}

func main() {
	fmt.Println(age) // 全局变量age 0
	scopeFunc()
	scopeFunc2()
	scopeBlock()
	// 此处无法访问scopeFunc2中定义的两个局部变量 编译报错
	//fmt.Println(name) // undefined: name
	//fmt.Println(blog) //undefined: blog
}
