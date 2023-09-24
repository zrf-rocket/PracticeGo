package main

/*
#include <stdio.h>
void hello()
{
	printf("This is CGo  -- from Cworld.\n");
}
*/

// 还有另外一个问题，那就是如果这里的C代码需要依赖一个非C标准库的第三方库，怎么办
// 呢？如果不解决的话必然会有链接时错误。Cgo提供了 #cgo 这样的伪C文法，让开发者有机会指
// 定依赖的第三方库和编译选项

// #cgo CFLAGS: -DPNG_DEBUG=1
// #cgo linux CFLAGS: -DLINUX=1
// #cgo LDFLAGS: -lpng
// #include <png.h>

// 示范了如何使用 CFLAGS 来传入编译选项，使用 LDFLAGS 来传入链接选项。 #cgo 还有另
// 外一种更简便一些的用法，如下所示：
// //#cgo pkg-config: png cairo
// //#include <png.h>

import "C"

func Hello() int {
	return int(C.hello())
}

/*
程序在Windows上执行此代码出错：
exec: "gcc": executable file not found in %PATH%

*/

// C程序
// 函数调用
// 对于常规的函数调用，开发者只要在运行 cgo 指令后查看一下生成的Go代码，就可以知道如
// 何写对应的调用代码。  对于常规返回了一个值的函数，调用者可以用以下
// 的方式顺便得到错误码
// n, err := C.atoi("a234")

// 在传递数组类型的参数时需要注意，在Go语言中将第一个元素的地址作为整个数组的起
// 始地址传入，这一点就不如C语言本身直接传入数组名字那么方便了
// n, err := C.f(&array[0]) // 需要显示指定第一个元素的地址

// 编译CGo
// 编译Cgo代码非常容易，我们不需要做任何特殊的处理。Go安装后，会自带一个cgo命令行
// 工具，它用于处理所有带有Cgo代码的Go文件，生成Go语言版本的调用封装代码。而Go工具集
// 对cgo命令行工具再次进行了良好的封装，使构建过程能够自动识别和处理带有Cgo代码的Go源
// 代码文件，完全不给用户增加额外的工作负担。

func main() {
	Hello()
}
