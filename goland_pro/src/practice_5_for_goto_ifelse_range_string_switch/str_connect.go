package main

import (
	"bytes"
	"fmt"
)

func main() {
	var str1 = "Steve"
	var str2 = "Rocket"
	fmt.Println(str1 + str2) //使用+拼接的方式 效率比较低

	//使用bytes.Buffer进行字符串的拼接
	fmt.Println(string_connect(str1, str2))
	fmt.Println(string_connect("Hello ", "GoLang"))
}

/**
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
