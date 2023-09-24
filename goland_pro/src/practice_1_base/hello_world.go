package main

//fmt 是 Go 语言的一个标准库/包，用来处理标准输入输出，类似于Python中的import os。不同点是Go中需要使用双引号把包名括起来
import "fmt"

// main 函数是整个程序的入口，main 函数所在的包名也必须为 main。类似于C语言中的主入口函数main，但是C的main函数定义为int main 或void main。
func main() {
	//调用 fmt 包的 Println 方法，打印出 Hello, SteveRocket!，并自带一个换行。
	fmt.Println("Hello, SteveRocket!")

	// 输出不同的语言
	fmt.Print("Hello, world; or Καλημέρα κόσμε; or こんにちは 世界\n")
}
