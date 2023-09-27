package main

import "fmt"

// 切片作为函数的参数进行传递
func sum(numbers []int) int {
	var total int
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 输出切片中所有元素的和
	fmt.Println(sum(numbers)) // 45

	Slice7()
}

// 在函数间传递 slice
func Slice7() {
	// 在函数间传递 slice 是很廉价的，因为 slice 相当于是指向底层数组的指针，让我们创建一个很大的 slice 然后传递给函数调用它
	slice := make([]int, 1e6)
	slice = foo(slice)
	fmt.Println(&slice[0])
}

func foo(slice []int) []int {
	return slice
}
