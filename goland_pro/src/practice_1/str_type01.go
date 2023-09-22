package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 全局变量的字符串
var WEIXIN_URL = "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"

func strFunc01() {
	str1 := "SteveRocket"  // 使用双引号声明字符串
	str2 := `公众号：CTO Plus` // 使用反引号声明字符串
	fmt.Println(WEIXIN_URL)
	fmt.Println(str1)
	fmt.Println(str2)
	/*
	   outputs
	      https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q
	      SteveRocket
	      公众号：CTO Plus
	*/

	c1 := '1'
	c2 := '周'
	c3 := 'A'
	fmt.Println(c1, c2, c3)              // 1 周 A
	fmt.Printf("%c %c %c\n", c1, c2, c3) // 1 周 A
}

func strFunc02() {
	fmt.Println(len(WEIXIN_URL))
	fmt.Println(WEIXIN_URL[0])
	fmt.Printf("%c\n", WEIXIN_URL[0])
	fmt.Println(WEIXIN_URL[:])
	//fmt.Println(WEIXIN_URL[-1]) // 不支持这种用法 invalid argument: index -1 (constant of type int) must not be negative

	/* outputs
	49
	104
	h
	https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q
	*/

	str := "微信公众号：CTO Plus"
	char := str[0]
	fmt.Println(char)            // 229
	fmt.Println(WEIXIN_URL[0:5]) // https

	// 字符串拼接
	str1 := "微信公众号"
	str2 := "CTO Plus"
	fmt.Println(str1 + "：" + str2) // 微信公众号：CTO Plus
	newStr := fmt.Sprintf("%s：%s", str1, str2)
	fmt.Println(newStr) // 微信公众号：CTO Plus

	var str3 string = ""
	str3 = str3 + "公众号："
	str3 += "CTO Plus"
	fmt.Printf("%s", str3) // 公众号：CTO Plus
}

func strFunc03() {

	name := "SteveRocket"
	age := 28
	fmt.Printf("My name is %s, I'm %d years old.\n", name, age) // My name is SteveRocket, I'm 28 years old.

	str := fmt.Sprintf("My name is %s, I'm %d years old.\n", name, age)
	fmt.Println(str) // My name is SteveRocket, I'm 28 years old.
	fmt.Println("\\ \a \r \n \t ' \\' \"")
}

func strFunc04() {
	str := "微信公众号：CTO Plus"
	bytes := []byte(str)
	// 将字符C改为字符A
	bytes[18] = 'A'
	str = string(bytes)
	fmt.Println(str) // 微信公众号：ATO Plus
}

func strFunc05() {
	str := "微信公众号：CTO Plus" // byte = 6*3+8
	str2 := "Author:SteveRocket"
	fmt.Println(reflect.TypeOf(str2[2]).Kind()) // uint8
	fmt.Println(str[18], string(str[18]))       // 67 C
	fmt.Printf("%d %c\n", str[0], str[0])       // 229 å
	fmt.Printf("%d %c\n", str2[0], str2[0])     // 65 A
	fmt.Println(len(str))                       // 26

	//将 string 转为 rune 数组
	runeArr := []rune(str)
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind()) // int32
	// 取出中文
	fmt.Println(runeArr[2], string(runeArr[2])) // 20844 公
	//此处输出str的长度为14  跟上面的 26已经不一样了
	fmt.Println(len(runeArr)) // 14
}

func strFunc06() {
	multiLine := `微信公众：CTO Plus
第二行内容
https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q
原样输出转移字符：\r\n\ \ \t  \'  \"
`
	fmt.Println(multiLine)
}

func numberToString() {
	//num := 68
	var num int = 67
	fmt.Println(string(num))
	fmt.Printf("%s\n", string(num))

	var num2 int = 1111
	fmt.Printf("%s\n", strconv.Itoa(num2)) // 1111 1111 公众号：CTO Plus
}

func main() {
	//strFunc01()
	//strFunc02()
	//strFunc03()
	//strFunc04()
	//strFunc05()
	//strFunc06()
	numberToString()
}
