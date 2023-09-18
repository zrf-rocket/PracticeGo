package main

import "fmt"

func arrayType() {
	//数组
	var arr [5]int                     //创建一个长度为5的int类型数组，初始值为0
	arr2 := [5]int{11, 22, 33, 44, 55} //创建一个包含5个元素的int类型数组
	fmt.Println(arr)
	fmt.Println(arr2)
}
