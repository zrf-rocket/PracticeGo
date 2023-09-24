package main

import "fmt"

func main() {
	var i int = 5

	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}

	j := 1
	for j <= 3 {
		fmt.Println(j)
		j += 1
	}
}
