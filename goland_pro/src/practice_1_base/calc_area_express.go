package main

import (
	"fmt"
	"math"
)

func calcCircumference() {
	var radius float64 = 5.0
	var area float64 = math.Pi * radius * radius
	var circumference float64 = 2 * math.Pi * radius

	fmt.Printf("半径为 %.2f 的圆的面积为 %.2f，周长为 %.2f\n", radius, area, circumference)
}

func express01() {
	//	避免整数除法的精度丢失
	//在Go语言中，整数除法会丢失小数部分，只保留整数部分。如果需要保留小数部分，可以将其中一个操作数转换为浮点数。例如：
	a := 10
	b := 3
	fmt.Println(a / b) // 3
	//fmt.Println(float64(a) / b) // invalid operation: float64(a) / b (mismatched types float64 and int)
	fmt.Println(float64(a) / float64(b)) // 3.3333333333333335
}

func express02() {
	//注意浮点数的精度问题
	//浮点数在计算机中是以二进制表示的，因此在进行浮点数运算时可能会出现精度问题。例如：
	a := 0.1
	b := 0.2
	fmt.Println(a + b) // 结果并不是0.3，而是一个接近0.3的近似值   0.30000000000000004
	//为了避免精度问题，可以使用math/big包中的Float类型进行高精度计算。

	//注意类型转换的安全性
	//在Go语言中，类型转换是一种将一个类型的值转换为另一个类型的操作。但是，需要注意类型转换的安全性，避免在类型转换时丢失精度或导致溢出。例如：
	num1 := 10000000000
	res1 := int32(num1) // 此时的结果是不确定的，因为a的值超出了int32类型的范围
	res2 := int64(num1)
	fmt.Println(res1, res2) // 1410065408 10000000000
	//为了避免类型转换的安全问题，可以使用类型断言来判断类型是否可转换。
}

// 注意指针的使用
// 在Go语言中，指针用于存储变量的内存地址。使用指针可以直接操作变量的值，而不需要进行值的拷贝。但是，需要注意指针的生命周期，避免在指针指向的变量被释放后继续使用指针。例如：
func getPointer() *int {
	value := 10
	return &value
}

func main() {
	//calcCircumference()
	//express01()
	express02()

	pointer := getPointer()
	result := *pointer
	fmt.Println(result)
}
