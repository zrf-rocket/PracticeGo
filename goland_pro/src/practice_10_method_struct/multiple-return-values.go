package main

import "fmt"

func vals() (int, int) { //函数签名中的(int，int)表示函数返回2个int类型值
	return 11, 011
}

/*
如果只想要返回值的一个子集，请使用空白标识符 _
*/
func main() {
	num1, num2 := vals()
	fmt.Println(num1)
	fmt.Println(num2)

	_, num := vals()
	fmt.Println(num)
	//fmt.Println(_)  //空白标识符 _ 不能用来输出值
}
