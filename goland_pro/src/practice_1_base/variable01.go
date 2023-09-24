package main

import "fmt"

func main() {
	Variable1()
	Variable2()
	Variable3()
	VariableStatement1()
	VariableStatement2()
	VariableStatement3()
	VariableInitialize1()
	VariableInitialize2()
	VariableInitialize3()
	VariableAssignment1()
	VariableAssignment2()
	VariableAssignment3()
	AnonymousVariables1()
	AnonymousVariables2()
	AnonymousVariables3()
	Constant1()
	Constant2()
	Constant3()
	LiteralConstants1()
	LiteralConstants2()
	LiteralConstants3()
	ConstantDefinition1()
	ConstantDefinition2()
	ConstantDefinition3()
	PredefinedConstants1()
	PredefinedConstants2()
	PredefinedConstants3()
	Enumeration1()
	Enumeration2()
	Enumeration3()
}

// 变量
func Variable1() {
	// 变量相当于是对一块数据存储空间的命名，程序可以通过定义一个变量来申请一块数据存储空间，之后可以通过引用变量名来使用这块存储空间。
	// 	声明
	var v1 int
	fmt.Println(v1) //0
	var v2 string
	var v3 [10]int //数组
	var v4 []int   //数组切片
	var v5 struct {
		f int
	}
	var v6 *int           //指针
	var v7 map[string]int //map  key为string类型  value为int类型
	var v8 func(a int) int

	// var 关键字的另一种用法是可以将若干个需要声明的变量放置在一起，免得程序员需要重复写 var 关键字
	var (
		v9  int
		v10 string
	)

	_, _, _, _, _, _, _, _, _, _ = v1, v2, v3, v4, v5, v6, v7, v8, v9, v10

	// 	初始化
	//以下三种都是正确的初始化方式
	var v11 int = 333
	var v12 = 100 // 编译器自动推导出类型
	v13 := "this is Golang"
	fmt.Println(v11)
	fmt.Println(v12)
	fmt.Println(v13)

	// 	赋值  变量初始化和变量赋值是两个不同的概念
	v1 = 123
	fmt.Println(v1)

	// 交换两个变量的值   多重赋值功能
	fmt.Println("交换前：v1=", v1, "v11=", v11)
	v1, v11 = v11, v1
	fmt.Println("交换后：v1=", v1, "v11=", v11)

	// 	匿名变量
	// 在调用函数时为了获取一个值，却因为该函数返回多个值而不得不定义一堆没用的变量。
	// 在Go中这种情况可以通过结合使用多重返回和匿名变量来避免这种丑陋的写法，让代码看起来更加优雅

	// 若只想获得 nickName ，则函数调用语句可以用如下方式编写
	_, _, nickName := GetName()
	fmt.Println(nickName)
}

func GetName() (firstName, lastName, nickName string) {
	return "zrf", "rocket", "steverocket"
}

// 常量
func Variable2() {
	// 常量是指编译期间就已知且不可改变的值。常量可以是数值类型（包括整型、浮点型和复数类型）、布尔类型、字符串类型等
	// 	字面常量（literal），是指程序中硬编码的常量
	// -12
	// 3.14159265358979323846  // 浮点类型的常量
	// 3.2+12i // 复数类型的常量
	// true // 布尔类型的常量
	// "foo"  // 字符串常量
	// Go语言的字面常量更接近我们自然语言中的常量概念，它是无类型的。
	// 只要这个常量在相应类型的值域范围内，就可以作为该类型的常量，比如上面的常量12，它可以赋值给
	// int 、 uint 、 int32 、int64 、 float32 、 float64 、 complex64 、 complex128 等类型的变量

	// 	常量定义
	const Pi float64 = 3.14159265358979323846
	//Pi = 0 //cannot assign to Pi
	fmt.Println(Pi)

	const zero = 0.0 // 无类型浮点常量
	const (
		size int64 = 1024
		eof        = -1 // 无类型整型常量
	)
	const u, v float32 = 0, 3   // u = 0.0, v = 3.0，常量的多重赋值
	const a, b, c = 3, 4, "foo" // 无类型整型和字符串常量
	fmt.Println("zero=", zero,
		"size=", size,
		"eof=", eof,
		"u=", u,
		"v=", v,
		"a=", a,
		"b=", b,
		"c=", c)
	// Go的常量定义可以限定常量类型，但不是必需的。如果定义常量时没有指定类型，那么它与字面常量一样，是无类型常量。
	// 常量定义的右值也可以是一个在编译期运算的常量表达式
	const mask = 1 << 3
	// 由于常量的赋值是一个编译期行为，所以右值不能出现任何需要运行期才能得出结果的表达
	// 式，比如试图以如下方式定义常量就会导致编译错误
	//const Home = os.GetEnv("HOME")
	// 原因很简单， os.GetEnv() 只有在运行期才能知道返回结果，在编译期并不能确定，所以无法作为常量定义的右值

	// 	预定义常量 true、false和iota
	// iota 比较特殊，可以被认为是一个可被编译器修改的常量，在每一个 const 关键字出现时被重置为0，
	// 然后在下一个 const 出现之前，每出现一次 iota ，其所代表的数字会自动增1
	const ( // iota被重设为0
		c0 = iota //0
		c1 = iota //1
		c2 = iota //2
	)
	fmt.Println(c0, c1, c2)

	const (
		a3 = 1 << iota //1   (iota在每个const开头被重设为0)
		b3 = 1 << iota
		c3 = 1 << iota
	)
	fmt.Println(a3, b3, c3) //1 2 4

	const ()

	const (
		u1         = iota * 42
		v1 float64 = iota * 42
		w1         = iota * 42
	)
	const x1 = iota
	const y1 = iota
	fmt.Println(u1, v1, w1, x1, y1)

	// 如果两个 const 的赋值语句的表达式是一样的，那么可以省略后一个赋值表达式。因此，上
	// 面的前两个 const 语句可简写为
	const (
		c4 = iota
		c5
		c6
	)
	const (
		a4 = 1 << iota
		a5
		a6
	)
	fmt.Println(c4, c5, c6) //0 1 2
	fmt.Println(a4, a5, a6) //1 2 4

	// 	枚举
	// Go语言并不支持众多其他语言明确支持的 enum 关键字
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays // 这个常量没有导出  为包内私有，其他符号则可被其他包访问。
	)
	// 同Go语言的其他符号（symbol）一样，以大写字母开头的常量在包外可见。
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, numberOfDays)
}

func Variable3() {

}

func VariableStatement1() {

}

func VariableStatement2() {

}

func VariableStatement3() {

}

// 	初始化
func VariableInitialize1() {

}

func VariableInitialize2() {

}

func VariableInitialize3() {

}

// 	赋值
func VariableAssignment1() {

}

func VariableAssignment2() {

}

func VariableAssignment3() {

}

// 	匿名变量
func AnonymousVariables1() {

}

func AnonymousVariables2() {

}

func AnonymousVariables3() {

}

// 常量
func Constant1() {

}

func Constant2() {

}

func Constant3() {

}

// 	字面常量
func LiteralConstants1() {

}

func LiteralConstants2() {

}

func LiteralConstants3() {

}

// 	常量定义
func ConstantDefinition1() {

}

func ConstantDefinition2() {

}

func ConstantDefinition3() {

}

// 	预定义常量
func PredefinedConstants1() {

}

func PredefinedConstants2() {

}

func PredefinedConstants3() {

}

// 	枚举
func Enumeration1() {

}

func Enumeration2() {

}

func Enumeration3() {

}
