package main

import (
	"fmt"
	"strings"
)

func main() {
	// StringsFunc1();
	// StringsFunc2();
	// StringsFunc3();
	// StringsFunc4();
	// StringsFunc5();
	StringsFunc6()
}

// 函数: ContainsAny(s, chars string) bool
// 说明: 判断字符串 s 中是否有字符 chars,如果有返回 true,否则返回false
func StringsFunc1() {
	str := "Hello GoLang strings"
	fmt.Println(strings.ContainsAny(str, "G")) //true  此处的G不能使用单引号
}

// 函数: Count(s, sep string) int
// 说明: 判断字符 sep 在字符串 s 中出现的次数，如果成功返回总个数  没有则返回0
func StringsFunc2() {
	str := "Hello GoLang strings nn"
	fmt.Println(strings.Count(str, "n"))
	fmt.Println(strings.Count(str, "o"))
	fmt.Println(strings.Count(str, "z")) //不存在则返回0
}

// 函数: EqualFold(s, t string) bool
// 说明: 判断字符串 s 是否与字符串 t 相等，并且不区分大小写
func StringsFunc3() {
	str := "Hello GoLang strings"
	str1 := "hello golang strings"
	fmt.Println(strings.EqualFold(str, str1))           //true
	fmt.Println(strings.EqualFold(str, "hello golang")) //false
}

// 函数: Fields(s string) []string
// 说明: 将字符串 s 以空格为分隔符拆分成若干个字符串，若成功则返回分割后的字符串切片
func StringsFunc4() {
	str := "Hello GoLang strings"
	for _, v := range strings.Fields(str) {
		fmt.Println(v)
	}
	fmt.Println(str)
}

// 函数: HasPrefix(s, prefix string) bool
// 说明: 判断字符串 s 是否是以字符 prefix 开头，如果是返回 true 否则返回 false
func StringsFunc5() {
	str := "Hello GoLang strings"
	fmt.Println(strings.HasPrefix(str, "hello")) //false  区分大小写
	fmt.Println(strings.HasPrefix(str, "Hello")) //true
}

// 函数: Index(s, sep string) int
// 说明: 判断字符 sep 在字符串 s 中第一次出现的位置，如果成功则返回sep位置的索引，如果字符 sep 不在字符串 s 中则返回 -1
func StringsFunc6() {
	str := "Hello GoLang strings"
	fmt.Println(strings.Index(str, "l")) //下标从0开始
	fmt.Println(strings.Index(str, "o"))
	fmt.Println(strings.Index(str, "h"))
	fmt.Println(strings.Index(str, "Z")) //不存在则返回-1
}
