package main

import "fmt"

func runeDemo1() {
	information := "微信公众号：CTO Plus"
	for _, i2 := range information {
		//我们发现英文可以正常输出，但是中文通过这种方式输出会乱码
		fmt.Printf("%c", i2) // 微信公众号：CTO Plus
	}

	fmt.Println("\n字符串长度：", len(information)) // 字符串长度： 26
	fmt.Println(information[18:])             // CTO Plus

	bytes := []byte(information)
	fmt.Println(bytes) // [229 190 174 228 191 161 229 133 172 228 188 151 229 143 183 239 188 154 67 84 79 32 80 108 117 115]

	for i, element := range information {
		fmt.Printf("%d  %c\n", i, element)
	}

	runnes := []rune{'微', '信', '公', '众', '号', '：', 'C', 'T', 'O', ' ', 'P', 'l', 'u', 's'}
	str := string(runnes)
	fmt.Println(len(str), str)       // 26 微信公众号：CTO Plus
	fmt.Println(len(runnes), runnes) // 14 [24494 20449 20844 20247 21495 65306 67 84 79 32 80 108 117 115]
}

func runeDemo2() {
	// 字符串修改是不能直接修改的，需要转成rune切片后再修改
}

func main() {
	runeDemo1()

}
