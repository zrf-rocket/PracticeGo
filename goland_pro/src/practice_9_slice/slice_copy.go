package main

import "fmt"

func sliceDemo() {
	src := []int{11, 22, 33, 44, 55}
	dst := make([]int, len(src))
	fmt.Println(src, len(src), cap(src)) // [11 22 33 44 55] 5 5
	fmt.Println(dst, len(dst), cap(dst)) // [11 22 33 44 55] 5 5
	fmt.Println(&src[0], &dst[0])        // 0xc0000be060 0xc0000be090

	count := copy(dst, src)
	fmt.Println("拷贝元素的个数：", count)       // 拷贝元素的个数： 5
	fmt.Println(dst, len(dst), cap(dst)) // [11 22 33 44 55] 5 5
	fmt.Println(&src[0], &dst[0])        // 0xc0000be060 0xc0000be090
}

func sliceDemo2() {
	src := []int{11, 22, 33, 44, 55}
	//直接声明变量 用赋值号=进行操作
	var src2 = src
	//src2切片和src1引用同一个内存地址
	fmt.Println(&src[0], &src[0])   // 0xc00000e3f0 0xc00000e3f0
	fmt.Println(&src2[0], &src2[0]) // 0xc00000e3f0 0xc00000e3f0
	fmt.Println(src, src2)          // [11 22 33 44 55] [11 22 33 44 55]
}

func main() {
	//sliceDemo()
	sliceDemo2()
}
