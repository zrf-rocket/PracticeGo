package main

import (
	"fmt"
	"reflect"
)

/*
   Go没有三元操作符
   !和<-可以用于一元和二元操作符来使用
   <-接收操作符

   7个算术操作符只能作用于整数：%、&、|、^、&^、<<、>>

   =  用于将一个值赋给一个已被声明的变量或常量
   :=  用于在声明一个变量的同时对这个变量进行赋值，且只能在函数体内使用

   ++、--是语句不是表达式
   	*p--  等同于 (*p)--
*/

func operator01() {
	//+ 连接符
	//只会创建并使用一个新的字符串值来保存操作结果，而不会改变任何操作数的值
	fmt.Println("Hello" + "GoLang" + "!")
	var str1 string = "Numbers:"
	fmt.Println(str1)
	fmt.Println(str1 + string(111)) //Numbers:o   Python中OK
	// fmt.Println(str1+111)  error  字符串和数字不能使用连接符

	//取出变量str1的地址
	fmt.Println(&str1)
	//根据地址取出值
	fmt.Println(*&str1)

	//  <- 接收操作符
	//	用于通道类型的值
	/*
		从一个通道类型的空值nil 接收值的表达式将会永远被阻塞
		从一个已被关闭的通道类型值接收值会永远成功并立即返回一个其元素类型的零值
		一个由接收操作符和通道类型的操作数所组成的表达式可以直接被用于变量赋值或初始化
	*/
	// <- ch

	// v1 := <-ch
	// v2 = <-ch

	// v, ok = <-ch
	// v, ok := <-ch
	// 通过同时对两个变量进行赋值或初始化  第二个变量将会是一个bool类型的值  表示接受操作的成功与否，可以通过它来判断一个通道是否已被关闭，如果值为false  就表示通道已关闭
}

func operator02() {
	num1 := 5
	num2 := 9
	//num1 > num2 ? num1 : num2  // Go不支持这种写法
	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)
	num2 <<= 2
	fmt.Println(num2)
	num2 >>= 2
	fmt.Println(num2)
	num2 &= 2
	fmt.Println(num2)
	num2 |= 2
	fmt.Println(num2)
	num2 ^= 2
	fmt.Println(num2)
	/* outputs
	14
	-4
	45
	0
	5
	36
	9
	0
	2
	0
	*/
}

func operator03() {
	//关系运算符
	num1 := 11
	num2 := 11
	num3 := 0
	isEqual := num1 == num2
	isNotEqual := num1 != num2
	isLess := num1 < num3
	isLessOrEqual := num1 <= num2
	fmt.Println(isEqual)       // true
	fmt.Println(isNotEqual)    // true
	fmt.Println(isLess)        // true
	fmt.Println(isLessOrEqual) // true

	// 逻辑运算符
	addr := &num1             // 取变量num1的地址
	result := *addr           // 解引用变量addr，得到变量num1的值
	fmt.Println(addr, result) // 0xc00001c098 11
	typeOfAddr := reflect.TypeOf(addr)
	typeOfResult := reflect.TypeOf(result)
	fmt.Println(typeOfAddr, typeOfResult) //  *int int
}

func main() {
	//operator01()
	//operator02()
	operator03()
}
