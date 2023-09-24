package main

import "fmt"

func runloop() {
	var i int = 5

	for {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
		if i < 0 {
			break
		}
	}
	/*
		The variable i is now: 4
		The variable i is now: 3
		The variable i is now: 2
		The variable i is now: 1
		The variable i is now: 0
		The variable i is now: -1
	*/
}

func runloop02() {
	for {
		fmt.Println("死循环")
	}
}

func main() {
	//runloop()
	runloop02()
}
