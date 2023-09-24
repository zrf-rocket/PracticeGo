package main

import (
	"fmt"
	"unicode/utf8"
)

func rangeStr() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	fmt.Println()

}

func forRange() {
	url := "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"

	for i := 0; i < len(url); i++ {
		fmt.Printf("%c", url[i])
	}
	fmt.Println()

	for index, char := range url {
		fmt.Printf("%v %c\n", index, char)
	}

	for _, char := range url {
		fmt.Printf("%c", char)
	}
}

func runeRange() {
	url := "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range url {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}

func rangeUnicode() {
	str2 := "Chinese: 中文，微信公众号：CTO Plus"
	fmt.Printf("The length of str2 is: %d\n", len(str2)) //The length of str2 is: 44

	// 使用range循环遍历字符串中的字符
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}

	// 使用utf8.RuneCountInString()函数获取字符串中的字符数量
	charCount := utf8.RuneCountInString(str2)
	fmt.Printf("字符数量: %d\n", charCount) // 字符数量: 26

	for i := 0; i < len(str2); {
		r, size := utf8.DecodeRuneInString(str2[i:])
		fmt.Printf("%c", r)
		i += size
	} // Chinese: 中文，微信公众号：CTO Plus
}

func main() {
	//rangeStr()
	//forRange()
	//runeRange()
	rangeUnicode()
}
