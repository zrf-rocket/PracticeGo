package main

import "fmt"
import "math"

const str string = "constant"

func main() {
	fmt.Println(str)

	const num = 500000000
	const d = 3e20 / num
	fmt.Println(num)
	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(math.Sin(num))
}
