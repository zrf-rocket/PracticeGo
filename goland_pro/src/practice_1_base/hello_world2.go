package main

// package 声明，表示该Go代码所属的包
import "fmt"

// import 语句，用于导入该程序所依赖的包
// 不得包含在源代码文件中没有用到的包，否则Go编译器会报编译错误。

/*
第一个Go程序
main函数是Go可执行程序的执行起点
Go语言的 main() 函数不能带参数，也不能定义返回值。命令行传入的参数在 os.Args 变量
中保存。如果需要支持命令行开关，可使用 flag 包。
*/
func main() {
	// var str1 String
	// str1 = "this is String"
	// fmt.Println(str1)

	//go的字符串必须双引号
	str := "This is GoLang"
	// fmt.Println("Hello Go Programming %s", str)     //此处输出 Hello Go Programming %s This is GoLang
	//此处的输出格式  不需要%s
	fmt.Println("Hello Go Programming ", str)
}
