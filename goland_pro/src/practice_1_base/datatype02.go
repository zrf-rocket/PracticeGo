package main

import (
	"fmt"
	"math"
	"reflect" //获取变量类型
)

func main() {
	FloatingPointType1()
	FloatingPointType2()
	FloatingPointType3()

	ComplexType1()
	ComplexType2()
	ComplexType3()

	String1()
	String2()
	String3()

	CharacterType1()
	CharacterType2()
	CharacterType3()
}

// 	浮点型   float32 、 float64
func FloatingPointType1() {
	// 浮点型用于表示包含小数点的数据，比如1.234就是一个浮点型数据。Go语言中的浮点类型采用IEEE-754标准的表达方式
	// 其中 float32 等价于C语言的 float 类型   float64 等价于C语言的 double 类型
	var fvalue1 float32
	fvalue1 = 12

	fvalue2 := 12.0 // 类型将被自动设为 float64，而不管赋给它的数字是否是用32位长度表示的
	// 如果不加小数点，fvalue2会被推导为整型而不是浮点型

	value2 := 12                         //整型
	fmt.Println(reflect.TypeOf(fvalue1)) //float32
	fmt.Println(reflect.TypeOf(fvalue2)) //float64
	fmt.Println(reflect.TypeOf(value2))  //int

	// fvalue1 = fvalue2  // cannot use fvalue2 (type float64) as type float32 in assignment
	fvalue1 = float32(fvalue2)
	fmt.Println(reflect.TypeOf(fvalue2)) //float64
}

func FloatingPointType2() {
	// 因为浮点数不是一种精确的表达方式，所以像整型那样直接用 == 来判断两个浮点数是否相等是不可行的，这可能会导致不稳定的结果
	// 推荐的做法
	fmt.Println("1.222 1.224", IsEqual(1.222, 1.224, 0.01))
	fmt.Println("1.222 1.224", IsEqual(1.222, 1.224, 0.0001))
}

// p为用户自定义的比较精度，比如0.00001
func IsEqual(f1, f2, p float64) bool {
	// return math.Fdim(f1, f2) < p
	return math.Dim(f1, f2) < p
}

// func IsEqual(f1, f2, p float32)bool{
// 	return math.Fdim(f1, f2) < p
// }

func FloatingPointType3() {

}

// 	复数类型  complex64 、 complex128
func ComplexType1() {
	// 复数实际上由两个实数（在计算机中用浮点数表示）构成，一个表示实部（real），一个表示虚部（imag）

	//复数表示形式
	var value1 complex64 // 由2个float32构成的复数类型
	value1 = 3.2 + 12i
	value2 := 3.2 * 12i                 // value2是complex128类型
	value3 := complex(3.2, 12)          // value3结果同 value2
	fmt.Println(value1, value2, value3) //(3.2+12i) (0+38.4i) (3.2+12i)

	// 实部与虚部
	// 对于一个复数 z = complex(x, y) ，就可以通过Go语言内置函数 real(z) 获得该复数的实
	// 部，也就是 x ，通过 imag(z) 获得该复数的虚部，也就是 y
	fmt.Println("实部:", real(value3), "  虚部:", imag(value3))
}

//复数的实际应用
func ComplexType2() {

}

func ComplexType3() {

}

// 	字符串  string
func String1() {
	// 相比之下， C/C++语言中并不存在原生的字符串类型，通常使用字符数组来表示，并以字符指针来传递。
	var str string                         // 声明一个字符串变量
	str = "this is Golang string"          // 字符串赋值
	ch := str[0]                           // 取字符串的第一个字符
	fmt.Println("str:", str, "   ch:", ch) //此处的字符会输出对应的ASCII编码值
	fmt.Printf("%c\n", ch)                 //输出字符

	fmt.Printf("The length of \"%s\" is %d\n", str, len(str))

	// 字符串的内容可以用类似于数组下标的方式获取，但与数组不同，字符串的内容不能在初始化后被修改
	// str[0]='X'  // cannot assign to str[0]
}

/*
Go语言仅支持UTF-8和Unicode编码 对于其他编码，Go语言标准库并没有内置的编码转换支持。不过，所幸的是我们可以很容易基于 iconv 库用Cgo包装一个
https://github.com/xushiwei/go-iconv

每个中文字符在UTF-8中占3个字节，而不是1个字节
*/
func String2() {
	str1 := "中"
	fmt.Println(len(str1)) //3  一个中文占3个字节长度
	str := "这是Golang中的字符串"
	fmt.Println(len(str))
	// 以字节数组的方式遍历  遍历了27次
	for i := 0; i < len(str); i++ {
		// fmt.Printf("index: %d, ch: %c\n", i, str[i])
		ch := str[i]
		fmt.Println(i, ch)
	}

	fmt.Println()

	// 以Unicode字符遍历   此处只遍历13次   每个字符的类型是 rune （早期的Go语言用 int 类型表示Unicode字符），而不是 byte
	for i, ch := range str {
		fmt.Println(i, ch) //ch的类型为rune
	}

}

func String3() {

}

// 	字符类型  rune
func CharacterType1() {
	// 一个是 byte （实际上是 uint8 的别名），代表UTF-8字符串的单个字节的值；
	// 另一个是 rune ，代表单个Unicode字符。
	// rune 相关的操作，可查阅Go标准库的 unicode 包。另外 unicode/utf8 包也提供了UTF8和Unicode之间的转换
	// Go语言的多数API都假设字符串为UTF-8编码。尽管Unicode字符在标准库中有支持，但实际上较少使用。

}

func CharacterType2() {

}

func CharacterType3() {

}
