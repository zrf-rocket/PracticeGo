package main

import (
	"errors"
	"fmt"
	"mymaths"
)

func main() {
	FunctionDefinition1()
	FunctionTransfer1()
	VariableParameter1()
	VariableParameter2()
	VariableParameter3()
	MultipleReturnValue1()
	AnonymousFunctionWithClosures1()
	AnonymousFunctionWithClosures2()
}

// 函数
// 	定义()
func FunctionDefinition1() {
	res, ok := Add(-1, 0)
	fmt.Println(res, ok) //0 Shoud be non-negative numbers
	res, ok = Add(11, 22)
	fmt.Println(res, ok)      //33 <nil>
	fmt.Println(Add2(33, 44)) //77
}

// func Add(num1 int, num2 int)(ret int, err error){  //ok
// 如果参数列表中若干个相邻的参数类型的相同，比如上面例子中的 num1 和 num2 ，则可以在参数列表
func Add(num1, num2 int) (ret int, err error) { //ok
	if num1 < 0 || num2 < 0 { // 假设这个函数只支持两个非负数字的加法
		err = errors.New("Shoud be non-negative numbers")
		return
	}
	return num1 + num2, nil // 支持多重返回值
}

func Add2(num1, num2 int) int {
	res := num1 + num2
	return res
}

// 	调用
func FunctionTransfer1() {
	//mymaths.Sub调用自定义的包中的方法
	res := mymaths.Sub(11, 11)
	fmt.Println(res)
}

// 小写字母开头的函数只在本包内可见，大写字母开头的函数才能被其他包使用。
// 这个规则也适用于类型和变量的可见性。

// 	不定参数
func VariableParameter1() {
	myfunc(1, 3, 5, 7, 9)
	myfunc() //长度为0

	//调用上 需要使用[]int{} 来构造一个数组切片实例
	myfunc2([]int{2, 4, 6, 8, 0})

}

// 不定参数是指函数传入的参数个数为不定数量。为了做到这点，首先需要将函数定义为接受不定参数类型
func myfunc(args ...int) { //接受不定数量的参数，这些参数的类型全部是 int
	fmt.Println(len(args))
	for _, arg := range args {
		fmt.Printf("%d\t", arg)
	}
	fmt.Println()
}

// 形如 ...type 格式的类型只能作为函数的参数类型存在，并且必须是最后一个参数。它是一
// 个语法糖（syntactic sugar），即这种语法对语言的功能并没有影响，但是更方便程序员使用。通
// 常来说，使用语法糖能够增加程序的可读性，从而减少程序出错的机会。
// 类型 ...type 本质上是一个数组切片，也就是 []type ，这也是为
// 什么上面的参数 args 可以用 for 循环来获得每个传入的参数

// 如果不用...type 这样的语法
func myfunc2(args []int) {
	for _, arg := range args {
		fmt.Printf("%d\t", arg)
	}
	fmt.Println()
}

// 不定参数的传递
func VariableParameter2() {
	funcArgs(11, 22, 33, 44, 55, 66, 77, 88, 99, 0)
	// funcArgs()
}
func funcArgs(args ...int) {
	// 按原样传递
	myfunc3(args...)
	// 传递片段，实际上任意的int slice都可以传进去
	myfunc3(args[1:]...)
}

func myfunc3(args ...int) {
	fmt.Println(len(args))
	for _, arg := range args {
		fmt.Printf("%d\t", arg)
	}
	fmt.Println()
}

// 任意类型的不定参数
// 希望传任意类型，可以指定类型为interface{}
func VariableParameter3() {
	printf("formats", "11", "22", "33", "44") //formats [11 22 33 44]
	printf("formats", "")                     //format []

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "GoLang"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)
}

// Go语言标准库中 fmt.Printf() 的函数原型
func printf(format string, args ...interface{}) {
	fmt.Println(format, args)
}

// 用 interface{} 传递任意类型数据是Go语言的惯例用法。使用 interface{} 仍然是类型安全的

func MyPrintf(args ...interface{}) {
	fmt.Println(len(args))
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, " is a int value.")
		case string:
			fmt.Println(arg, " is a string value.")
		case int64:
			fmt.Println(arg, " is a int64 value.")
		default:
			fmt.Println(arg, " is an unknown type.")
		}
	}
}

// 	多返回值
func MultipleReturnValue1() {
	// 比如 File.Read() 函数就可以同时返回读取的字节数和错误信息。如果读取文件成功，则返回值中的 n 为读取的字节数， err 为 nil ，否则 err 为具体的出错信息：
}

// func (file *File) Read(b []byte)(n int, err Error)
// 从上面的方法原型可以看到，我们还可以给返回值命名，就像函数的输入参数一样。返回值被命名之后，
// 它们的值在函数开始的时候被自动初始化为空。在函数中执行不带任何参数的 return 语句时，会返回对应的返回值变量的值。
// Go语言并不需要强制命名返回值，但是命名后的返回值可以让代码更清晰，可读性更强，同时也可以用于文档。
// 如果调用方调用了一个具有多返回值的方法，但是却不想关心其中的某个返回值，可以简单
// 地用一个下划线“ _ ”来跳过这个返回值，比如下面的代码表示调用者在读文件的时候不想关心
// Read() 函数返回的错误码：
// n, _ := f.Read(buf)

// 	匿名函数与闭包
func AnonymousFunctionWithClosures1() {
	// 匿名函数是指不需要定义函数名的一种函数实现方式
	// 函数可以像普通变量一样被传递或使用，这与C语言的回调函数比较类似。不同的是，Go语言支持随时在代码里定义匿名函数。

	//带函数名的一般函数
	// func func_name() {
	// }

	// 匿名函数由一个不带函数名的函数声明和函数体组成
	// 匿名函数可以直接赋值给一个变量或者直接执行
	f := func(a, b int, z float64) bool {
		return a*b < int(z)
	}
	fmt.Println(f(11, 11, 121))

	f2 := func(x, y int) int {
		return x + y
	}
	fmt.Println(f2(11, 22)) //33

	//直接调用
	func() { fmt.Println("最简单的匿名函数") }()
	// func(ch chan int){
	// 	ch <- ACX
	// }(reply_chan)  //// 花括号后直接跟参数列表表示函数调用

	fmt.Println(func(num1, num2 int) bool { return num1 < num2 }(11, 22)) //true
}

func AnonymousFunctionWithClosures2() {
	// Go的匿名函数是一个闭包
	// 闭包是可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者
	// 任何全局上下文中定义，而是在定义代码块的环境中定义。要执行的代码块（由于自由变量包含
	// 在代码块中，所以这些自由变量以及它们引用的对象没有被释放）为自由变量提供绑定的计算环境（作用域）。

	// 闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示
	// 数据还要表示代码。支持闭包的多数语言都将函数作为第一级对象，就是说这些函数可以存储到
	// 变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回。

	// Go语言中的闭包同样也会引用到函数外的变量。闭包的实现确保只要闭包还被使用，那么被闭包引用的变量会一直存在

	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("*i, j:%d, %d\n", i, j)
		}
	}()
	// 变量 a 指向的闭包函数引用了局部变量 i 和 j ， i 的值被隔离，在闭包外不能被修改，改变 j 的值以后，再次调用 a ，发现结果是修改过的值。
	// 在变量 a 指向的闭包函数中，只有内部的匿名函数才能访问变量 i ，而无法通过其他途径访问到，因此保证了 i 的安全性。
	a() // 10 5

	j *= 2
	a() // 10 10
}
