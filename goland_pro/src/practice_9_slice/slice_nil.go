package main

import "fmt"

// nil 和 empty slice
func Slice2() {
	// 创建一个 nil slice，创建一个 nil slice 的方法是声明它但不初始化它
	var slice []int                //声明slice
	var slice1 = []int{1, 3, 5, 7} //初始化
	// 使用 slice 字面值创建
	slice2 := []int{} //初始化     空的slice
	// 使用 make 创建
	slice3 := make([]int, 0)

	fmt.Println(slice)      //[]
	fmt.Println(len(slice)) //0
	// fmt.Println(slice[0])  //panic: runtime error: index out of range

	// if not slice {   //error
	// if slice {  		//error
	if slice == nil {
		fmt.Println("切片Slice是空的")
	} else {
		fmt.Println("切片Slice是的内容：", slice)
	}
	if slice1 == nil {
		fmt.Println("切片Slice1是空的")
	} else {
		fmt.Println("切片Slice1是的内容：", slice1)
	}
	if slice2 == nil || 0 == len(slice2) { //!len(slice2)行不通
		fmt.Println("切片Slice2是空的")
	} else {
		fmt.Println("切片Slice2是的内容：", slice2)
	}
	if slice3 == nil || 0 == len(slice3) { //!len(slice2)行不通
		fmt.Println("切片Slice3是空的")
	} else {
		fmt.Println("切片Slice3是的内容：", slice3)
	}

	fmt.Println(slice2, len(slice2))
	// _ = slice2[0]  //error  panic: runtime error: index out of range
	fmt.Println(slice3, len(slice3))
	// _ = slice3[0]  //error  panic: runtime error: index out of range
	// empty slice 包含0个元素并且底层数组没有分配存储空间。当我们想要表示一个空集合时它很有用处，比如一个数据库查询返回0个结果。
	// 不管我们用 nil slice 还是 empty slice，内建函数 append，len和cap的工作方式完全相同。
}

func main() {
	Slice2()
}
