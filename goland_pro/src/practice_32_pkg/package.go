package main

import (
	"fmt"
	"runtime"
)

/*****************
* 同一个代码包中，可以存在多个代码包初始化函数，甚至代码包内的每一个源码文件都可以定义多个代码包初始化函数
* Go编译器无法保证同一个代码包中的多个代码包初始化函数的执行顺序
* 如果确实需要按特定顺序执行的话，可以考虑使用GoLang并发编程模型中的channel进行控制或者调整代码包初始化函数的顺序
******************/

// 包初始化函数，作为代码包的初始化，该函数无需任何参数声明和结果声明，且名称必须为init
// init方法从上到下依次被初始化执行
func init() {
	println("this is other two init function()")
	//先被执行  定义的全局变量name未被赋初值，默认为空字符串
	if name != "" {
		println("my name is ", name)
	} else {
		println("name not is init")
	}
	//给全局变量name赋值
	name = "SteveRocket"
}

func init() { //第一条语句在执行时，变量m已经被初始化并赋值，说明当前代码包中所有全局变量的初始化会在代码包初始化函数执行前完成
	println("this is init function")
	fmt.Printf("Map result:%v\n", m)
	info = fmt.Sprintf("OS:%s, Arch:%s", runtime.GOOS, runtime.GOARCH)
}

// 定义一个全局变量  map类型  已被显示赋值
var m map[int]string = map[int]string{11: "AAA", 22: "BBB", 33: "CCC"}

// 全局变量  string类型  未被显示赋值
var info string //info默认会被赋予类型string的零值 空字符串""
// var info string  //error 相同作用域下，不允许存在相同名字的全局变量
var name string

func init() {
	println("this is other one init function()")
	//在第24行中给name变量了赋值
	println("my name is ", name)
	name = "GoLang"
}

func show() {
	fmt.Println("this is show function in package_1 model")
}

/*******************
* 同一个代码包中除了初始化函数以外不允许再出现相同的函数名称
* 'show' redeclared in this block less  Reports names redeclared in this block.
*******************/
//func show(){  //show方法在上面已经被定义
//	fmt.Println("this is other show function")
//}

/*****************************
* GoLang会在整个程序真正执行前对整个程序的依赖进行分析，并初始化相关的代码包，此处会初始化package_1和package_2两个代码包
* 所有的代码包中的初始化函数init都会在main函数之前执行完毕，而且只执行一次，并且当前代码包中的所有全局变量的初始化会在代码
* 包初始化函数执行前完成，避免了在代码包初始化函数对某个变量进行赋值之后又被该变量声明中赋予的值覆盖掉的问题。
******************************/
func main() {
	//init()  //此处不允许手动调用init方法
	// init方法会在main方法执行之前自动被调用，也就是说一个程序，main方法并不是优先执行

	fmt.Println(info)
	show()
}
