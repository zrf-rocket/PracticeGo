package main

import (
	"fmt"
	"strconv"
)

//十进制 八进制 十六进制

func octal() {
	num := 0o77
	fmt.Printf("八进制数%#o表示的十进制数是%d\n", num, num) // 八进制数077表示的十进制数是63

	num2 := 077
	fmt.Printf("十进制：%d  八进制：%o  类型：%T\n", num2, num2, num2) // 十进制：63  八进制：77  类型：int
}

func hexadecimal() {
	// 使用十六进制表示数值
	num := 0xFF
	fmt.Printf("十六进制数%#x表示的十进制数是%d\n", num, num) // 十六进制数0xff表示的十进制数是255

	num2 := 0x1234567
	fmt.Printf("十进制：%d  十六进制数：%x  类型：%T\n", num2, num2, num2) // 十进制：19088743  十六进制数：1234567  类型：int
}

func main() {
	//octal()
	//hexadecimal()

	// 将十进制数转换为八进制和十六进制
	num := 42
	oct := fmt.Sprintf("%o", num)
	hex := fmt.Sprintf("%x", num)
	hex2 := fmt.Sprintf("%X", num)
	fmt.Printf("十进制数%d转换为八进制是%s，转换为十六进制是%s，大写十六进制是%s\n", num, oct, hex, hex2)
	// 十进制数42转换为八进制是52，转换为十六进制是2a，大写十六进制是2A

	// 十进制转二进制
	decimal := 10
	binary := strconv.FormatInt(int64(decimal), 2)
	fmt.Println(binary)

	// 二进制转十进制
	binaryStr := "1010"
	decimal, _ = strconv.ParseInt(binaryStr, 2, 64)
	fmt.Println(decimal) // 输出：10
}
