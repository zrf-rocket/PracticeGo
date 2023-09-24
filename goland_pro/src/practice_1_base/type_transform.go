package main

import "fmt"

/*
	类型转换  用于将一种数据类型的变量转换为另外一种类型的变量
*/
func main() {
	fmt.Println(int2float(123))

	var sum int = 1224
	count := 12
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Println(mean)
}

//将整型转化为浮点型，并计算结果，将结果赋值给浮点型变量
func int2float(int64 int64) float64 {
	return float64(int64)
}
