package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

/*
可变参数的函数可以用任何数量的参数来调用。 例如，fmt.Println()就是一个常见的可变函数。
*/
func main() {
	sum(1, 2, 3)
	sum(2, 3, 4, 5)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
