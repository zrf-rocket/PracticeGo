package main

import "fmt"

func main() {
	fmt.Println(fibonacci(34))
}

func fibonacci(i int) int {
	if i < 2 {
		return i
	}
	return fibonacci(i-2) + fibonacci(i-1)
}
