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
	fmt.Println(sum(numbers))  // 45
}
