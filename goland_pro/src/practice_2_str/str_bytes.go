package main

import (
	"bytes"
	"fmt"
)

//使用bytes.Buffer对字符串进行高效的拼接
func main() {
	fmt.Println("This is GoLang Env")

	var str1 = "Steve"
	var str2 = "Rocket"
	fmt.Println(str1 + str2) //使用+拼接的方式 效率比较低

	//使用bytes.Buffer进行字符串的拼接
	fmt.Println(string_connect(str1, str2))
	fmt.Println(string_connect("Hello ", "GoLang"))
}

/*
字符串连接函数
将传入的两个string类型的参数进行拼接
返回拼接后的结果
*/
func string_connect(str1, str2 string) string {
	var result bytes.Buffer
	result.WriteString(str1)
	result.WriteString(str2)
	return result.String()
}

//Go中可以使用“+”合并字符串，但是这种合并方式效率非常低，每合并一次，都是创建一个新的字符串,就必须遍历复制一次字符串。
// Java中提供StringBuilder类来解决这个问题。Go中也有类似的机制，那就是Buffer
//var name bytes.Buffer
//fmt.Println(name)
//for i:=0; i< 5; i++{
//	name.WriteString("Go")
//}
//fmt.Println(name.String())
//fmt.Println(name.WriteString("NewString"))
//fmt.Println(name.String())
