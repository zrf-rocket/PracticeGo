package main

import "fmt"

func sliceDelete() {
	slice := []int{11, 22, 33, 44, 55}
	fmt.Println(slice)                                    // [11 22 33 44 55]
	index := 2                                            // 需要删除的元素的索引
	newSlice := append(slice[:index], slice[index+1:]...) // append()函数中第一个切片不用拆开，原因是一个作为被接受的一方，是把后面的元素追加到第一个切片中。
	fmt.Println(slice)                                    // [11 22 44 55 55]
	fmt.Println(newSlice)                                 // [11 22 44 55]
}

func sliceResize() {
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[:3]                     // 调整切片的大小为前三个元素
	fmt.Println(len(slice), cap(slice))       // 5 5
	fmt.Println(len(newSlice), cap(newSlice)) // 3 5

	newSlice2 := slice[1:2]
	fmt.Println(len(newSlice2), cap(newSlice2)) // 1 4

	newSlice3 := append(slice[:1], slice[len(slice)-1:]...)
	fmt.Println(len(newSlice3), cap(newSlice3)) // 2 5
}

func main() {
	sliceDelete()
	sliceResize()
}
