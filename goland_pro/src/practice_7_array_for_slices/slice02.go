package main

import "fmt"

func main() {
	slices := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slices[:])   //[1 2 3 4 5 6 7]
	fmt.Println(slices[2:3]) // [3]
	fmt.Println(slices[:3])  // [1 2 3]   从起始位置开始
	fmt.Println(slices[2:])  // [3 4 5 6 7]  不包括指定的起始位置
	// fmt.Println(slices[:-1]) 不支持这种    Python支持

	//按需创建
	fmt.Println(make([]int, 5)) //[0 0 0 0 0]    Python的列表之间的元素使用逗号间隔
	slices2 := make([]int, 0, 5)
	fmt.Println(slices2) // []
}
