package main

import "fmt"

/*
Go 语言中的基本数据类型包含
bool      the set of boolean (true, false)

uint8      the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8      the set of all signed  8-bit integers (-128 to 127)
int16      the set of all signed 16-bit integers (-32768 to 32767)
int32      the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64      the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32      the set of all IEEE-754 32-bit floating-point numbers
float64      the set of all IEEE-754 64-bit floating-point numbers

complex64      the set of all complex numbers with float32 real and imaginary parts
complex128      the set of all complex numbers with float64 real and imaginary parts

byte      alias for uint8
rune      alias for int32
uint      either 32 or 64 bits
int      same size as uint
uintptr      an unsigned integer large enough to store the uninterpreted bits of a pointer value

string      the set of string value (eg: "hi")

可以将基本类型分为三大类：
布尔类型
数字类型
字符串类型

高级数据类型：
数组 array
切片 slice
Map map
自定义类型 type
结构体 struct
函数 func
方法 method
接口 interface
*/

func intDataType() {
	var age = 28
	var num1 int = 100
	fmt.Printf("类型：%T 数值：%d\n", age, age)   // %T 用于查看变量的类型   类型：int 数值：28
	fmt.Printf("类型：%T 数值：%d\n", num1, num1) // 类型：int 数值：100

	num4 := int8(1)
	num5 := int16(12345)
	num6 := int32(1234567890)
	//num6 := int32(11234567890)  // cannot convert 11234567890 (untyped int constant) to type int32
	num7 := int64(1123434566565687789)
	fmt.Printf("类型：%T 数值：%d\n", num4, num4)
	fmt.Printf("类型：%T 数值：%d\n", num5, num5)
	fmt.Printf("类型：%T 数值：%d\n", num6, num6)
	fmt.Printf("类型：%T 数值：%d\n", num7, num7)
}

func floatDataType() {
	var f1 float32 = 12.2
	fmt.Printf("类型：%T 数值：%f\n", f1, f1) // 类型：float32 数值：12.200000

	f2 := 1234                          // 类型推导为int类型
	fmt.Printf("类型：%T 数值：%d\n", f2, f2) // 类型：int 数值：1234

	f3 := 1234.
	// Go语言中的小数默认类型是：float64
	fmt.Printf("类型：%T 数值：%f\n", f3, f3) // 类型：float64 数值：1234.000000

	var f4 = float32(12)
	fmt.Printf("类型：%T 数值：%f\n", f4, f4) // 类型：float32 数值：12.000000

	f5 := float64(12345)
	fmt.Printf("类型：%T 数值：%f\n", f5, f5) // 类型：float64 数值：12345.000000
}

func boolDataType() {
	var b1 bool // 默认为false
	var b2 = true
	ok := false
	fmt.Println(b1, b2, ok)                             // false true false
	fmt.Printf("类型：%T  %T  值为%v  %v\n", b1, ok, b1, b2) // 类型：bool  bool  值为false  true
}

func complexDataType() {
	// 创建复数
	c1 := complex(3, 4) // 3 + 4i
	c2 := 2 + 5i        // 2 + 5i

	// 复数运算
	sum := c1 + c2      // 加法
	diff := c1 - c2     // 减法
	product := c1 * c2  // 乘法
	quotient := c1 / c2 // 除法

	// 输出结果
	fmt.Println(sum)      // 输出：(5+9i)
	fmt.Println(diff)     // 输出：(1-1i)
	fmt.Println(product)  // 输出：(-14+23i)
	fmt.Println(quotient) // 输出：(0.7647058823529411+0.058823529411764705i)
}

func nilDataType() {
	var ptr *int
	if ptr == nil {
		fmt.Println("ptr is nil") // 输出：ptr is nil
	}

	var slice []int
	if slice == nil {
		fmt.Println("slice is nil") // 输出：slice is nil
	}
}

func byteDataType() {
	var c1 byte = 'a'
	fmt.Printf("类型：%T  %c\n", c1, c1) // 类型：uint8  a

	str := "Hello, SteveRocket!"

	// 将字符串转换为字节数组
	bytes := []byte(str)

	// 遍历字节数组
	for _, b := range bytes {
		fmt.Printf("%c ", b) // 输出：H e l l o ,   S t e v e R o c k e t !
	}
}

func main() {
	//intDataType()
	//floatDataType()
	//boolDataType()
	//complexDataType()
	//nilDataType()
	byteDataType()
}
