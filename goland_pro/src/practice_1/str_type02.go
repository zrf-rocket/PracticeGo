// 字符串操作
package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func strEncry(r rune) rune {
	return r + 1
}

// 定义一个字符串类型
var WEIXIN_URL = "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"

/*
遍历字符串则迭代各个Unicode代码点(称为Go中的rune s)，可能跨越多个字节。
*/
func strings_operator() {
	// 字符串遍历方式1  按Unicode字符遍历字符串
	for key, val := range WEIXIN_URL {
		//代码中的变量 val，实际类型是 rune 类型，以十六进制打印出来就是字符的编码。
		fmt.Printf("key:%d  value:%c-%x\n", key, val, val)
	}

	// 字符串遍历方式2
	//使用strings.Map()函数将一个字符串中的每个字符都往后移一位
	strURL := WEIXIN_URL
	fmt.Println(strURL)
	mapStr := strings.Map(strEncry, strURL)
	fmt.Println(mapStr)

	//获取字符串的长度
	fmt.Println(len(WEIXIN_URL))
	// 字符串遍历方式3  遍历每一个ASCII字符
	for i := 0; i < len(WEIXIN_URL); i++ {
		fmt.Println(WEIXIN_URL[i], rune(WEIXIN_URL[i]), string(rune(WEIXIN_URL[i])))
	}

	// 数字转字符串 并输出变量类型
	num1 := 100
	num1_str := strconv.Itoa(num1)
	fmt.Println(num1_str, reflect.TypeOf(num1_str))

	str := "123"
	// 字符串转数字 并输出变量类型
	str_int, _ := strconv.Atoi(str)
	fmt.Println(str_int, reflect.TypeOf(str_int))

}

func main() {
	strings_operator()
}
