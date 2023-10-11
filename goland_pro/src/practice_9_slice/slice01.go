package main

import (
	"fmt"
	"reflect"
)

func sliceDemo01() {
	// 创建方式1：使用var 创建一个空的string类型切片
	var sElements []string

	// 声明切片 创建一个空的int类型切片
	var s0 []int
	var s []int
	// 使用len和cap 获取切片的长度和容量
	fmt.Printf("切片长度:%d 切片容量%d\n", len(s), cap(s)) // 切片长度:0 切片容量0

	fmt.Println(sElements == nil)    // 结果输出为true  表示为空，没有开辟内存空间
	fmt.Println(s0 == nil, s == nil) // true true

	// 两个切片不能使用==来判断是否相等
	//fmt.Println(s == s0)  // Invalid operation: s == s0 (the operator == is not defined on []int)

	// 创建方式2：使用make创建一个长度为5的int类型切片
	s1 := make([]int, 5)
	fmt.Printf("切片长度:%d 切片容量%d\n", len(s1), cap(s1)) // 切片长度:5 切片容量5

	// 创建方式3：创建一个包含5个元素的int类型切片并初始化
	s2 := []int{11, 22, 33, 44, 55}
	fmt.Printf("切片长度:%d 切片容量%d\n", len(s2), cap(s2)) // 切片长度:5 切片容量5
	// 可以使用下标访问切片中的元素
	fmt.Println(s2[0]) // 输出第一个元素的值：11

	// 通过下标（元素的索引）设置第2个元素的值为2222
	s2[1] = 2222
	fmt.Println(s2) // [11 2222 33 44 55]

	//通过fori遍历切片
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%d ", s2[i])
	} // 11 2222 33 44 55
	fmt.Println()

	// 创建方式4：使用var定义一个整型的切片并初始化
	var nums = []int{1, 2, 3, 4, 5}
	// 直接输出切片
	fmt.Println(nums) // [1 2 3 4 5]

	//在声明的同时初始化slice
	var address = []string{"北京", "上海", "广州", "深圳"}

	//通过for range遍历切片
	for i, s3 := range address {
		fmt.Println(i, s3)
	}
	//0 北京
	//1 上海
	//2 广州
	//3 深圳
	fmt.Println(nums == nil) // 结果输出为false  表示初始化成功，值都不等于nil，为非空，开辟了内存空间
	fmt.Println(address == nil)
}

func sliceDemo02() {
	// 创建方式5：创建一个数组
	var arr1 = [...]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 00}
	slice := arr1[0:5]                         // 此处基于数组做切割
	fmt.Println(arr1, len(arr1), cap(arr1))    // [11 22 33 44 55 66 77 88 99 0] 10 10
	fmt.Println(slice, len(slice), cap(slice)) // [11 22 33 44 55] 5 10
	slice = arr1[4:8]
	//此处slice容量大小等于数组arr1的长度-索引位置
	fmt.Println(slice, len(slice), cap(slice)) // [55 66 77 88] 4 6

	//slice的切片操作
	var slice0 []int = arr1[2:8]
	var slice1 []int = arr1[:]
	var slice2 []int = arr1[:len(arr1)]
	var slice3 []int = arr1[0:len(arr1)]
	var slice4 []int = arr1[:len(arr1)-1]
	fmt.Printf("%v %d %d\n", slice0, len(slice0), cap(slice0)) // [33 44 55 66 77 88] 6 8
	fmt.Printf("%v %d %d\n", slice1, len(slice1), cap(slice1)) // [11 22 33 44 55 66 77 88 99 0] 10 10
	fmt.Printf("%v %d %d\n", slice2, len(slice2), cap(slice2)) // [11 22 33 44 55 66 77 88 99 0] 10 10
	fmt.Printf("%v %d %d\n", slice3, len(slice3), cap(slice3)) // [11 22 33 44 55 66 77 88 99 0] 10 10
	fmt.Printf("%v %d %d\n", slice4, len(slice4), cap(slice4)) // 去除最后一个元素 [11 22 33 44 55 66 77 88 99] 9 10
}

func makeSlice() {
	s1 := []int{}
	fmt.Println(s1, len(s1), cap(s1)) // [] 0 0

	// 省略cap，相当于 cap = len
	var s2 []int = make([]int, 0)
	fmt.Println(s2, len(s2), cap(s2)) // [] 0 0

	// 使用 make 创建slice，并指定 len 和 cap 值
	var s3 []int = make([]int, 0, 10)
	fmt.Println(s3, len(s3), cap(s3)) // [] 0 10

	s4 := [5]int{11, 22, 33, 44, 55}
	fmt.Println(s4, len(s4), cap(s4)) // [11 22 33 44 55] 5 5

	var s5 []int
	fmt.Println(s5, len(s5), cap(s5)) // [] 0 0

	var s6 []int = make([]int, 0, 0)
	fmt.Println(s6, len(s6), cap(s6)) // [] 0 0

	s7 := []int{11, 22, 33, 44, 55}
	fmt.Println(s7, len(s7), cap(s7)) // [11 22 33 44 55] 5 5

	sValue := make([]int, 5, 10)
	fmt.Println(sValue, len(sValue), cap(sValue)) //[0 0 0 0 0] 5 10
}

func sliceDemo03() {
	data := [...]int{11, 22, 33, 44, 55, 66}
	fmt.Println(data, reflect.TypeOf(data), len(data), cap(data)) // [11 22 33 44 55 66] [6]int 6 6
	slice := data[:]
	fmt.Println(slice, reflect.TypeOf(slice), len(slice), cap(slice)) // [11 22 33 44 55 66] []int 6 6

	fmt.Printf("%T  %T\n", data, slice) // [6]int  []int

	newSlice := append(slice, 123)
	fmt.Println(newSlice, reflect.TypeOf(newSlice), len(newSlice), cap(newSlice)) // [11 22 33 44 55 66 123] []int 7 12

	slice[0] += 100
	slice[1] += 200
	fmt.Println(slice) // [111 222 33 44 55 66]

	// 可直接创建 slice 对象，自动分配底层数组
	// 通过初始化表达式构造，可使用索引号
	slice2 := []int{11, 22, 33, 10: 99}
	fmt.Println(slice2) //[11 22 33 0 0 0 0 0 0 0 99]
}

func main() {
	//sliceDemo01()
	//sliceDemo02()
	//makeSlice()
	sliceDemo03()
}
