package main

import (
	"fmt"
	"reflect"
)

func sliceDemo1() {
	slice := []int{11, 22, 33, 44}
	// 使用取地址符号&来获取slice下标为0的元素的地址
	p := &slice[0]
	p1 := &slice[1]
	p2 := &slice[2]

	fmt.Printf("%p  %p  %p\n", p, p1, p2) // 0xc0000141e0  0xc0000141e8  0xc0000141f0
	// 通过*来获取底层数组的元素
	fmt.Println(*p, *p1, *p2) // 11 22 33

	*p = 111
	*p1 = 222
	*p2 = 333
	// 通过指针修改底层数组的元素，从而修改slice的元素
	fmt.Println(slice) // [111 222 333 44]
}

// 多维 slice
func Slice6() {
	// 也是同数组一样，slice 可以组合为多维的 slice
	slice := [][]int{{11}, {33, 44, 55}, {77}, {}}
	fmt.Println(slice, len(slice), cap(slice)) //[[11] [33 44 55] [77] []] 4 4
	fmt.Println(&slice[0][0], &slice[1][0])    //0xc04203c200 0xc0420423e0
	// 需要注意的是使用 append 方法时的行为，比如我们现在对 slice[0] 增加一个元素
	slice[0] = append(slice[0], 12)

	// 那么只有 slice[0] 会重新创建底层数组，slice[1] 则不会
	fmt.Println(slice)
	fmt.Println(&slice[0][0], &slice[1][0]) //0xc04203c250 0xc0420423e0
}

func sliceDemo2() {
	// [][]int，是指元素类型为 []int
	data := [][]int{
		[]int{11, 22, 33},
		[]int{1, 2, 3},
		[]int{},
	}
	// 生成的结果是个二维的切片，切片里面的每一个元素类型都为[]int，不能为其他任何类型
	fmt.Println(data) // [[11 22 33] [1 2 3] []]
	// 在此处对一第3个元素的第一个元素进行修改会引发panic
	//data[2][0] = 0  panic: runtime error: index out of range [0] with length 0
	data[2] = []int{111, 222, 333}
	fmt.Println(data) // [[11 22 33] [1 2 3] [111 222 333]]

	data[2][0] = 0
	fmt.Println(data)                 // [[11 22 33] [1 2 3] [0 222 333]]
	fmt.Println(reflect.TypeOf(data)) // [][]int

	// 使用type定义二维slice
	type intSlice []int
	numbers := []intSlice{
		[]int{1, 2, 3, 4},
		[]int{11, 22, 33},
	}

	fmt.Println(numbers, reflect.TypeOf(numbers)) // [[1 2 3 4] [11 22 33]]  []main.intSlice
}

func main() {
	//sliceDemo1()
	sliceDemo2()
}
