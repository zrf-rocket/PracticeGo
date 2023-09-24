package main

import (
	"fmt"
	// "reflect"  //获取变量类型
	// "math"
	"errors"
	"io"
	"log"
	"mymaths" //导入自定义的包
	"os"
)

//将当前脚本路径放入到GOPATH中

// 流程控制
// 	条件
// 	选择
// 	循环
// 	跳转
// 函数
// 	定义
// 	调用
// 	不定参数
// 	多返回值
// 	匿名函数与闭包
// 错误处理
// 	error接口
// 	defer
// 	panic()
// 	recover()
// 错误类型： error
// 复合类型
//  指针（pointer）
//   通道（chan）
//   结构体（struct）
//   接口（interface）

func main() {
	fmt.Println("顺序编程")
	// _ = os
	// ErrorInterface1()
	// ErrorInterface2()
	// ErrorInterface3()
	// ErrorDefer1()
	// ErrorDefer2()
	// ErrorDefer3()
	ErrorPanic1()
	ErrorPanic2()
	ErrorPanic3()
	ErrorRecover1()
	ErrorRecover2()
	ErrorRecover3()
}

// 错误处理  Error handling
// 	error接口()  Go语言引入了一个关于错误处理的标准模式，即 error 接口
type error interface {
	Error() string
}

// 对于大多数函数，如果要返回错误，大致上都可以定义为如下模式，将 error 作为多种返回值中的最后一个，但这并非是强制要求
func Foo(param int) (n int, err error) {
	return n, err
}

func ErrorInterface1() {
	// 错误处理不是语言规范的一部分，通常只作为一种编程范式存在，比如C语言中的 errno 。但自C++语言以来，语言层面上会增加错误处理的支持，
	// 比如异常（exception）的概念和 try-catch 关键字的引入。Go语言在此功能上考虑得更为深远。漂亮的错误处理规范是Go语言最大的亮点之一。

	// 调用时的代码建议按如下方式处理错误情况
	n, err := Foo(0)
	if err != nil {
		//错误处理
		fmt.Println("错误处理", n, err)
	} else {
		fmt.Println("n:", n, err)
	}
}

// 用Go库中的实际代码来示范如何使用自定义的 error 类型
// 定义一个用于承载错误信息的类型。因为Go语言中接口的灵活性，你根本不需要从
// error 接口继承或者像Java一样需要使用 implements 来明确指定类型和接口之间的关系
/*
type PathError struct{
	Op string
	Path string
	Err error
}
// 如果这样的话，编译器又怎能知道 PathError 可以当一个 error 来传递呢？关键在于下面的代码实现了 Error() 方法
func (e *PathError) Error() string{
	return e.Op + " " + e.Path + ": " + e.Err.Error()
	// 之后就可以直接返回 PathError 变量了
}
// 下面的代码中，当 syscall.Stat() 失败返回 err 时，将该 err 包装到一个 PathError 对象中返回
func Stat(name string)(fi FileInfo, err error){
	var stat syscall.Stat_t

	err = syscall.Stat(name, &stat)

	if err != nil{
		return nil, &PathError{"stat", name, err}
	}
	return fileInfoFromStat(&stat, name), nil
}
func ErrorInterface2(){
	fi, err = os.Stat("a.txt")
	if err != nil{
		if e,ok := err.(*os.PathError); ok && e.Err != nil{
			// 获取PathError类型变量e中的其他信息并处理
		}
	}
}
*/

func ErrorInterface3() {

}

// 	defer
// 一个函数中可以存在多个 defer 语句，因此需要注意的是， defer 语句的调用是遵照
// 先进后出的原则，即最后一个 defer 语句将最先被执行。只不过，当你需要为 defer 语句到底哪
// 个先执行这种细节而烦恼的时候，说明你的代码架构可能需要调整一下了。
func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}

	defer dstFile.Close()

	// 即使其中的 Copy() 函数抛出异常，Go仍然会保证 dstFile 和 srcFile 会被正常关闭。
	// 如果觉得一句话干不完清理的工作，也可以使用在 defer 后加一个匿名函数的做法
	defer func() {
		fmt.Println("释放资源........")
	}()

	return io.Copy(dstFile, srcFile)
}

//将2.txt文件中的内容 拷贝到1.txt文件中
func ErrorDefer1() {
	// 在C/C++中还有另一种解决方案。开发者可以将需要释放的资源变量都声明在函数的开头部
	// 分，并在函数的末尾部分统一释放资源。函数需要退出时，就必须使用 goto 语句跳转到指定位
	// 置先完成资源清理工作，而不能调用 return 语句直接返回。

	// Go语言使用 defer关键字简简单单地解决了这个问题

	fmt.Println(CopyFile("1.txt", "2.txt"))
}

func ErrorDefer2() {

}

func ErrorDefer3() {

}

// 	panic()
func ErrorPanic1() {
	// Go语言引入了两个内置函数 panic() 和 recover() 以报告和处理运行时错误和程序中的错误场景
	// func panic(interface{})
	// func recover(interface{})

	/*
	   当在一个函数执行过程中调用 panic() 函数时，正常的函数执行流程将立即终止，但函数中
	   之前使用 defer 关键字延迟执行的语句将正常展开执行，之后该函数将返回到调用函数，并导致
	   逐层向上执行 panic 流程，直至所属的goroutine中所有正在执行的函数被终止。错误信息将被报
	   告，包括在调用 panic() 函数时传入的参数，这个过程称为错误处理流程。

	   从 panic() 的参数类型 interface{} 我们可以得知，该函数接收任意类型的数据，比如整
	   型、字符串、对象等。调用方法很简单
	*/
	// panic(404)
	// panic("network broken")
	// panic(Error("file not exists"))

	/*
	   recover() 函数用于终止错误处理流程。一般情况下， recover() 应该在一个使用 defer
	   关键字的函数中执行以有效截取错误处理流程。如果没有在发生异常的goroutine中明确调用恢复
	   过程（使用 recover 关键字），会导致该goroutine所属的进程打印异常信息后直接退出。

	   对于 foo() 函数的执行要么心里没底感觉可能会触发错误处理，或者自己在其中明确加
	   入了按特定条件触发错误处理的语句，那么可以用如下方式在调用代码中截取 recover()
	*/

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Runtime error caught:%v", r)
		}
	}()
	// foo()
	// 无论 foo() 中是否触发了错误处理流程，该匿名 defer 函数都将在函数退出时得到执行。假
	// 如 foo() 中触发了错误处理流程， recover() 函数执行将使得该错误处理过程终止。如果错误处
	// 理流程被触发时，程序传给 panic 函数的参数不为 nil ，则该函数还会打印详细的错误信息。
	log.Printf("Runtime....") //2017/05/10 14:48:17 Runtime....
}

func ErrorPanic2() {

}

func ErrorPanic3() {

}

// 	recover()
func ErrorRecover1() {

}

func ErrorRecover2() {

}

func ErrorRecover3() {

}
