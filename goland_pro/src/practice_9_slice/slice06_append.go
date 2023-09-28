package main

import (
	"fmt"
)

func sliceAppend() {
	var slice1 = []int{11, 22, 33, 44, 55}
	fmt.Printf("%v\n", slice1) // [11 22 33 44 55]

	slice2 := []int{66, 77, 88, 99, 10}
	fmt.Println(slice2) // [66 77 88 99 10]

	//将两个slice合并成一个新的slice
	newSlice := append(slice1, slice2...)
	fmt.Println(newSlice) //[66 77 88 99 10]

	//向 slice 尾部添加数据，返回新的 slice 对象
	newSlice2 := append(newSlice, 1, 2, 3, 4)
	fmt.Println(newSlice2)                                               // [11 22 33 44 55 66 77 88 99 10 1 2 3 4]
	fmt.Printf("%p %p %p %p\n", &slice1, &slice2, &newSlice, &newSlice2) // 0xc000004078 0xc0000040a8 0xc0000040d8 0xc000004108

	slice3 := make([]int, 0, 5)
	//在切片末尾添加5个元素
	newSlice3 := append(slice3, 1, 2, 3, 4, 5)
	fmt.Println(slice3)    // []
	fmt.Println(newSlice3) // [1 2 3 4 5]

	slice3 = append(slice3, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(slice3)                     // [1 2 3 4 5 6 7 8]
	fmt.Printf("%p %p\n", &slice3, &slice3) // 0xc000096120 0xc000096120
	fmt.Println(newSlice3)                  // [1 2 3 4 5]

	numbers := make([]int, 1, 1)
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d 切片长度：%d  切片容量：%d\n", i, len(numbers), cap(numbers))
	}
	fmt.Println(numbers) // [0 0 1 2 3 4 5 6 7 8 9]

	address := []string{"北京", "上海", "广州", "深圳"}
	fmt.Println("长度：", len(address), "容量：", cap(address)) // 长度： 4 容量： 4
	address = append(address, "武汉")
	fmt.Println("长度：", len(address), "容量：", cap(address)) // 长度： 5 容量： 8
}

func sliceAppend02() {
	slice := make([]int, 5, 10)
	fmt.Println(slice) // [0 0 0 0 0]
	fmt.Println(cap(slice))
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}
	fmt.Println(slice) // [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
}

func main() {
	//sliceAppend()
	sliceAppend02()
}
