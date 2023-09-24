package main

import "fmt"

/*
切片是Go语言中的关键数据类型，为序列提供了比数组更强大的接口。
*/
func main() {
	s := make([]string, 3) //要创建非零长度的空切片，请使用内置make()函数。这里创建一个长度为3的字符串(初始为零值)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	//可以像数组一样设置和获取字符串的子串值。len()函数返回切片的长度。

	// 内置 append()函数，它返回包含一个或多个新值的切片。注意，需要接收append()函数的返回值，因为可能得到一个新的slice值。
	s = append(s, "d")
	fmt.Println(s)
	s = append(s, "e", "f")
	fmt.Println(s)

	// 可以复制切片。这里创建一个与切片s相同长度的空切片c，并从切片s复制到c中。
	// 切片支持具有语法为slice[low:high]的切片运算符。 例如，这获得元素s[2]，s[3]和s[4]的切片。
	// 这切片到(但不包括)s[5]。这切片从(包括)s[2]。可以在一行中声明并初始化slice的变量。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	//包括开头不包括结尾
	l := s[2:5]
	fmt.Println(l)
	l = s[:5]
	fmt.Println(l)
	l = s[2:]
	fmt.Println(l)

	// 切片可以组成多维数据结构。内切片的长度可以变化，与多维数组不同。
	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
