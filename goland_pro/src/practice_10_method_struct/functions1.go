package main

import (
	"fmt"
	"math"
)

/*
函数
Go 语言最少有个 main() 函数
可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务
函数声明告诉了编译器函数的名称，返回类型，和参数
Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。
如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的函数个数。

func：函数由 func 开始声明




函数参数
函数返回值
函数可变参数
函数参数传递（引用传递，值传递）

Scala 函数中函数的返回值默认值最后一个表达式的值，不需要return
*/
func main() {
	//fmt.Println(len(1234))  //invalid argument 1234 (type int) for len
	fmt.Println(len("hello Golang")) //12

	fmt.Println(max_num(12, 34))

	//定义局部变量
	var num1 int = 123
	var num2 int = 456
	var ret int //定义的局部变量  默认值为0

	fmt.Println(ret)
	ret = max_num(num1, num2)
	fmt.Println(ret)

	res1, res2 := multi_result("Golang", "Scala")
	//此处函数由两个返回值   需要分别使用res1,res2来接收
	fmt.Println(res1, res2)

	res1, res2 = swap("Golang", "Scala")
	fmt.Println(res1, res2)

	//交换两个元素
	var v1 int = 12111
	var v2 int = 34111
	fmt.Println(v1, v2)
	v1, v2 = v2, v1
	fmt.Println(v1, v2)

	func_args()
	use_inner_func()
	fmt.Println()
	//nextNumber 为一个函数，函数 i 为 0
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	//创建新的函数 nextNumberNew，并查看结果
	nextNumberNew := getSequence()
	fmt.Println(nextNumberNew())
	fmt.Println(nextNumberNew())
	fmt.Println(nextNumberNew())

	var c1 Circle
	c1.radius = 10
	fmt.Println("Area of Circle(c1) = ", c1.getArea())

}

//返回最大值
func max_num(num1, num2 int) int {
	//return num1 ? num1 > num2 : num2;

	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

//返回多个值
func multi_result(str1, str2 string) (string, string) {
	return str1, str2
}

//交换两个值，并返回多个值
func swap(str1, str2 string) (string, string) {
	return str2, str1
}

//函数参数
//函数如果使用参数，该变量可称为函数的形参  形参就像定义在函数体内的局部变量
func func_args() {
	//值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。  存在副本机制
	//引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
	//默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

	num1 := 123
	num2 := 456
	//swap_addr(num1, num2)//引用传递不应该这样传  而应该使用下面的地址传递
	swap_addr(&num1, &num2)
	fmt.Println(num1, num2) //最终实现了两数字之间的交换
}

func swap_addr(num1, num2 *int) {
	var temp int
	temp = *num1
	*num1 = *num2
	*num2 = temp
}

//函数定义后可作为值来使用
//在定义的函数中初始化一个变量，该函数仅仅是为了使用内置函数 math.sqrt()
func use_inner_func() {
	//声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
}

//闭包是匿名函数，可在动态编程中使用
//Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明
//创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 方法就是一个包含了接受者的函数
// Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，
// 接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。
//定义一个结构体类型和该类型的一个方法
//定义函数
type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	//c.radius  即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}
