package main

import "fmt"

func main() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf, len(mf)) // map[1:0xbccdc0 2:0xbccde0 5:0xbcce00] 3
	fmt.Println(mf[1], mf[1]()) // 0xbccdc0 10
	fmt.Println(mf[2], mf[2]()) // 0xbccde0 20
	fmt.Println(mf[5], mf[5]()) // 0xbccde0 20
}
