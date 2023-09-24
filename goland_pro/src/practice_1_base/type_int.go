package main

import "fmt"

func main() {
	Integer1()
	Integer2()
	Integer3()
}

// 	整型  int8 、 byte 、 int16 、 int 、 uint 、 uintptr 等
func Integer1() {
	var age = 28
	fmt.Println(age)
	fmt.Printf("年龄为：%T  %d\n", age, age)

	// 对于常规的开发来说，用 int
	// 和 uint 就可以了，没必要用 int8 之类明确指定长度的类型，以免导致移植困难。
	// int 和 int32 在Go语言里被认为是两种不同的类型，编译器也不会帮你自动做类型转换
	var value2 int32
	value1 := 64 // value1将会被自动推导为int类型
	// value2 = value1  //error   cannot use value1 (type int) as type int32 in assignment

	// 使用强制类型转换可以解决这个编译错误
	value2 = int32(value1)
	fmt.Println(value1, value2)
	// 在做强制类型转换时，需要注意数据长度被截短而发生的数据精度损失（比如
	// 将浮点数强制转为整数）和值溢出（值超过转换的目标类型的值范围时）问题
}

func Integer2() {
	var i, j, k int
	i, j, k = 0, 1, 2
	if i == j {
		fmt.Println(i, j, " 相等")
	} else {

		fmt.Println(i, j, " 不相等")
	}

	fmt.Println(j % k) //1
	fmt.Println(j / k) //0

	// var result = k/i  //panic: runtime error: integer divide by zero
	// fmt.Println(result)
}

func Integer3() {
	// 两个不同类型的整型数不能直接比较，比如 int8 类型的数和 int 类型的数不能直接比较，但
	// 各种类型的整型变量都可以直接与字面常量（literal）进行比较
	var i int32
	var j int64
	i, j = 1, 2
	// if i == j{} // invalid operation: i == j (mismatched types int32 and int64)
	if i == 1 || j == 2 {
		fmt.Println("符合条件")
	}

	// Go语言的大多数位运算符与C语言都比较类似 除了取反在C语言中是~x   而在Go语言中是 ^x
}
