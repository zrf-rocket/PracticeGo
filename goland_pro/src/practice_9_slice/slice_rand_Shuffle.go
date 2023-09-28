package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := []int{11, 22, 33, 44, 55}
	fmt.Println(slice) // [11 22 33 44 55]
	for i := 0; i < len(slice); i++ {
		res := rand.Intn(i + 1)
		slice[i], slice[res] = slice[res], slice[i]
	}
	fmt.Println(slice)
}
