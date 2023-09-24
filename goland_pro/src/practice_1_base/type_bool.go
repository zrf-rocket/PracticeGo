package main

import "fmt"

func main() {
	BooleanType1()
	BooleanType2()
	BooleanType3()
}

// 类型
// 类型 Types of
// 	布尔类型   bool
func BooleanType1() {
	var v1 bool
	v1 = true
	v2 := (1 == 2)
	fmt.Println(v1, v2) //true false
	// v3 := ("" == false)  //invalid operation: "" == false (mismatched types string and bool)
	// v3 := ("" == '')  //invalid operation: "" == '\u0027' (mismatched types string and rune)
	v3 := ("" == "")
	fmt.Println(v3) //true

	// 布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换
	var b bool
	// b = 1  // 编译错误
	// b = bool(1)  // 编译错误

	b = (1 != 0)
	fmt.Println(b) //true
}

func BooleanType2() {

}

func BooleanType3() {

}
