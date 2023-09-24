package main

import "fmt"

func main() {
	nums := []int{11, 22, 44}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for index, num := range nums {
		if num == 22 {
			fmt.Println("index:", index)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
