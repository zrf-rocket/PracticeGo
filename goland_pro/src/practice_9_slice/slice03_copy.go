package main

import "fmt"

func sliceCopy() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println(src) // [1 2 3 4 5]
	fmt.Println(dst) // [1 2 3 4 5]

	slice := []int{11, 22, 33, 44, 55}
	fmt.Printf("%v\n", slice) // [11 22 33 44 55]

	slice2 := make([]int, 10)
	fmt.Printf("%v\n", slice2) // [0 0 0 0 0 0 0 0 0 0]

	copy(slice2, slice)
	fmt.Printf("%v\n", slice2) // [11 22 33 44 55 0 0 0 0 0]

	fmt.Println(append(slice2, slice...)) // [11 22 33 44 55 0 0 0 0 0 11 22 33 44 55]
}

func main() {
	sliceCopy()
}
