package main

import "fmt"

func sliceDel() {
	slice := []int{11, 22, 33, 44, 55}
	fmt.Println(slice)                                    // [11 22 33 44 55]
	index := 2                                            // 需要删除的元素的索引
	newSlice := append(slice[:index], slice[index+1:]...) // append()函数中第一个切片不用拆开，原因是一个作为被接受的一方，是把后面的元素追加到第一个切片中。
	fmt.Println(slice)                                    // [11 22 44 55 55]
	fmt.Println(newSlice)                                 // [11 22 44 55]
}

func sliceClear() {
	slice := []int{11, 22, 33, 44, 55}
	fmt.Println(slice, len(slice), cap(slice)) // [11 22 33 44 55] 5 5
	slice = slice[0:0]
	fmt.Println(slice, len(slice), cap(slice)) // [] 0 5
}

func main() {
	//sliceDel()
	sliceClear()
}
