package main

import (
	"fmt"
	"os"
)

func calcSum(numbers ...int) int {
	result := 0
	for _, val := range numbers {
		result += val
	}
	return result
}

func calcSum02(blog string, numbers ...int) (result int) {
	result = len(numbers) + len(blog)
	return
}

//可变长参数必须放在函数参数的最后
// 以下可变参数的用法都是错误的
// func calcSum03(numbers ...int, infos ...string) {  // can only use ... with final parameter in list
//
// }

//func calcSum04(numbers ...int, blog string) {
//
//}

func main() {
	//可变长度参数的函数 参数不传递也可以
	fmt.Println(calcSum())                          // 0
	fmt.Println(calcSum(1))                         // 1
	fmt.Println(calcSum(1, 2, 3, 4, 5, 6, 7, 8, 9)) // 45
	//可变长度参数的函数 参数不传递也可以
	fmt.Println(calcSum02("https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"))             // 29
	fmt.Println(calcSum02("https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q", 1, 2, 3, 4)) // 53

	values := []int{11, 22, 33, 44, 55}
	fmt.Println(calcSum(values...)) // 165

	fmt.Printf("func01 = %T\n", func01) // func01 = func(...int)
	fmt.Printf("func02 = %T\n", func02) // func02 = func([]int)

	lineNumber, name := 10101, "SteveRocket"
	errorf(lineNumber, "undefined:%s", name)                      // line 10101:undefined:SteveRocket
	errorf(lineNumber, "undefined:%s", 1, 2, 3, 4, 5, 6)          // line 10101:undefined:%!s(int=1)%!(EXTRA int=2, int=3, int=4, int=5, int=6)
	errorf(lineNumber, "undefined:%s", "SteveRocket", false, 1.2) // line 10101:undefined:SteveRocket%!(EXTRA bool=false, float64=1.2)
}

func func01(...int) {

}

func func02([]int) {

}

func errorf(lineNumber int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "line %d:", lineNumber)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
