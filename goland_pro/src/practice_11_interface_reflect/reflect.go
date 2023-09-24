package main

/*
#include <stdlib.h>

//其实在这个块注释中，可以写
//任意合法的C源代码，而Cgo都会进行相应的处理并生成对应的Go代码
*/

// 这个注释的内容是有意义的，而不是传
// 统意义上的注释作用。这个例子里用的是一个块注释，实际上用行注释也是没问题的，只要是紧
// 贴在 import 语句之前即可。比如下面也是正确的Cgo写法
// #include <stdio.h>
// #include <stdlib.h>

import (
	"C"
	"fmt"
	"reflect"
)

/*
反射
反射（reflection）是在Java出现后迅速流行起来的一种概念。通过反射，你可以获取丰富的类型信息，并可以利用这些类型信息做非常灵活的工作。

在Java中，你可以读取配置并根据类型名称创建对应的类型，这是一种常见的编程手法。Java
中的很多重要框架和技术（比如Spring/IoC、Hibernate、Struts）等都严重依赖于反射功能。虽然
时，使用Java EE时很多人都觉得很麻烦，比如需要配置大量XML格式的配置程序，但这毕竟不
是反射的错，反而更加说明了反射所带来的高可配置性。

大多数现代的高级语言都以各种形式支持反射功能，除了一切以性能为上的C++语言。Go
语言的反射实现了反射的大部分功能，但没有像Java语言那样内置类型工厂，故而无法做到像Java
那样通过类型字符串创建对象实例。

Type 和 Value ，它们也是Go语言包中 reflect 空间里最重要的两个类型
	type MyReader struct{
		Name string
	}
	func (r MyReader)Read(p []byte)(n int, err error){}

因为 MyReader 类型实现了 io.Reader 接口的所有方法（其实就是一个 Read() 函数），所以
MyReader 实现了接口 io.Reader 。我们可以按如下方式来进行 MyReader 的实例化和赋值
	var	reader io.Reader
	reader = &MyReader("a.txt")

什么是 Type ，什么是 Value
对所有接口进行反射，都可以得到一个包含 Type 和 Value 的信息结构。
对上例的reader 进行反射，也将得到一个 Type 和 Value ， Type 为 io.Reader ， Value 为 MyReader{"a.txt"}
Type 主要表达的是被反射的这个变量本身的类型信息，而 Value 则为该变量实例本身的信息。


	基本用法
1. 获取类型信息
2. 获取值类型
*/

func main() {
	// reflectGetTypeInfo()
	// reflectGetValueType()
	// reflectStruct()

	fmt.Println(Random())
	Seed(100)

}

// 1. 获取类型信息
func reflectGetTypeInfo() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x)) //type: float64

	// Type 和 Value 都包含了大量的方法，其中第一个有用的方法应该是 Kind ，这个方法返回该类型的具体信息： Uint 、 Float64 等。
	// Value 类型还包含了一系列类型方法，比如 Int() ，用于返回对应的值
	v := reflect.ValueOf(x)
	fmt.Println(v)        //3.4
	fmt.Println(v.Type()) //float64
	// fmt.Println(v.Value())
	fmt.Println(v.Kind(), reflect.Float64)                       //float64 float64
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64) //true
	fmt.Println("Value:", v.Float())                             //Value: 3.4
}

// 2. 获取值类型
func reflectGetValueType() {
	// 类型 Type 中有一个成员函数 CanSet() ，其返回值为 bool 类型。如果你在注意到这个函数之前就直接设置了值，很有可能会收到一些看起来像异常的错误处理消息。

	// 为什么要有这么个奇怪的函数，可以设置所有的域不是很好吗？
	// Go语言中所有的类型都是值类型，即这些变量在传递给函数的时候将
	// 发生一次复制。

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println(v, x)

	// v.Set(4.1)  //cannot use 4.1 (type float64) as type reflect.Value in argument to v.Set
	// 最后一条语句试图修改 v 的内容。是否可以成功地将 x 的值改为4.1呢？先要理清 v 和 x 的关系。

	// 在
	// 调用 ValueOf() 的地方，需要注意到 x 将会产生一个副本，因此 ValueOf() 内部对 x 的操作其实
	// 都是对着 x 的一个副本。假如 v 允许调用 Set() ，那么我们也可以想象出，被修改的将是这个 x 的
	// 副本，而不是 x 本身。如果允许这样的行为，那么执行结果将会非常困惑。调用明明成功了，为
	// 什么 x 的值还是原来的呢？
	// 为了解决这个问题Go语言，引入了可设属性这个概念（Settability）。
	// 如果 CanSet() 返回 false ，表示你不应该调用 Set() 和 SetXxx() 方法，否则会收到这样的
	// 错误：
	// panic: reflect.Value.SetFloat using unaddressable value

	// 有些场景下不能使用反射修改值，那么到底什么情况下可以修改的呢？其实
	// 这还是跟传值的道理类似。
	// 直接传递一个 float 到函数时，函数不能对外部的这个
	// float 变量有任何影响，要想有影响的话，可以传入该 float 变量的指针。

	// 成功地用反射的方式修改了变量 x 的值
	p := reflect.ValueOf(&x) // 注意：得到X的地址
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	fmt.Println()

	v1 := p.Elem()
	fmt.Println("settability of v1:", v1.CanSet())

	v1.SetFloat(7.1)
	fmt.Println(v1.Interface())
	fmt.Println(v1, x)
}

// 对结构的反射操作
// 如何获取一个结构中所有成员的值
func reflectStruct() {
	type T struct {
		A int
		B string
	}
	t := T{203, "mh203"}
	s := reflect.ValueOf(&t).Elem()
	typeOff := s.Type()
	fmt.Println("s.NumField()：", s.NumField())
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		//分别输出变量名  类型 以及 值
		fmt.Printf("%d: %s %s = %v\n", i, typeOff.Field(i).Name, f.Type(), f.Interface())
	}
}

// 对于结构的反射操作并没有根本上的不同，只是用了 Field() 方法来按索引获取
// 对应的成员。当然，在试图修改成员的值时，也需要注意可赋值属性。

/*
语言交互性
像Java这样非常重视面向对象身份的语言也都提供了JNI机制，以调用那些C代码库
Go确实也提供了这一功能，称为Cgo
*/
func Random() int {
	return int(C.random())
}
func Seed(i int) {
	C.srandom(C.uint(i))
}

// 在函数中使用了 C包包含的 random() 和 srandom() 函数，顺便还用了一个 C 包中提供的 uint 类型。
// 看起来确实没有问题，但再细想一下，马上就会蹦出很多疑问来。
// (1) Go语言标准库里的包名字都是小写的，这个名字大写的 C 包怎么看都不像是Go自带的，
// 但我也没有装过这个包，它到底从哪里来的呢？
// (2) 为什么要在 import 前面写上那么奇怪的一段完全就是C语法的注释？这段注释是必需
// 的吗？
// (3) 不是说包内类型的可见性是由首个字母的大小写决定的吗？为什么这里能够使用 C 包里
// 以小写字母开头的函数和类型呢？

// 事实上，根本就不存在一个名为 C 的包。这个 import 语句其实就是一个信号，告诉Cgo它应
// 该开始工作了。做什么事情呢？就是对应这条 import 语句之前的块注释中的C源代码自动生成
// 包装性质的Go代码。如果你对以下这些概念有所了解，就相对比较容易理解Cgo这个声称Go代
// 码的过程：Java的JNI、.NET的P/Invoke、RPC和WebService的Proxy/Stub、IDL语言和WSDL语
// 言等。

// 函数调用从汇编的角度看，就是一个将参数按顺序压栈
// （push），然后进行函数调用（call）的过程。Cgo生成的代码只不过是帮你封装了这个压栈和调用
// 的过程，从外面看起来就是一个普通的Go函数调用。

/*
	类型映射
在跨语言交互中，比较复杂的问题有两个：类型映射以及跨越调用边界传递指针所带来的对象生命周期和内存管理的问题。比如Go语言中的 string 类型需要跟C语言中的字符数组进行对
应，并且要保证映射到C语言的对象的生命周期足够长，以避免在C语言执行过程中该对象就已
经被垃圾回收。

对于C语言的原生类型，Cgo都会将其映射为Go语言中的类型： C.char 和 C.schar （对应于
C语言中的 signed char ）、 C.uchar （对应于C语言中的 unsigned char ）、 C.short 和 C.ushort
（对应于 unsigned short ）、 C.int 和 C.uint （对应于 unsigned int ）、 C.long 和 C.ulong
（对应于 unsigned long ）、 C.longlong （对应于C语言中的 long long ）、 C.ulonglong （对
应于C语言中的 unsigned long long 类型）以及 C.float 和 C.double 。C语言中的 void* 指针
类型在Go语言中则用特殊的 unsafe.Pointer 类型来对应。
C语言中的 struct 、 union 和 enum 类型，对应到Go语言中都会变成带这样前缀的类型名称：
struct_ 、 union_ 和 enum_ 。比如一个在C语言中叫做 person 的 struct 会被Cgo翻译为
C.struct_person 。

如果C语言中的类型名称或变量名称与Go语言的关键字相同，Cgo会自动给这些名字加上下
划线前缀。



	字符串映射
因为Go语言中有明确的 string 原生类型，而C语言中用字符数组表示，两者之间的转换是
一个必须考虑的问题。Cgo提供了一系列函数来提供支持： C.CString 、 C.GoString 和
C.GoStringN 。需要注意的是，每次转换都将导致一次内存复制，因此字符串内容其实是不可
修改的（实际上，Go语言的 string 也不允许对其中的内容进行修改）。
由于 C.CString 的内存管理方式与Go语言自身的内存管理方式不兼容，我们设法期待Go语
言可以帮我们做垃圾收集，因此在使用完后必须显式释放调用 C.CString 所生成的内存块，否则
将导致严重的内存泄露。

defer 用法，所有用到 C.CString 的代码大致都可以写成如下的风格：
	var gostr string
	// ... 初始化gostr
	cstr := C.CString(gostr)
	defer C.free(unsafe.Pointer(cstr))
	// 接下来大胆地使用cstr吧，因为保证可以被释放掉了
	// C.sprintf(cstr, "content is: %d", 123)

*/

/*
连接符号
链接符号关心的是如何将语言文法使用的符号转化为链接期使用的符号，在常规情况下，链
接期使用的符号对我们不可见，但是在一些特殊情况下，我们需要关心这一点，比如：在用  gdb
调试的时候，要设置断点： b  <函数名>，这里的<函数名>是指“链接符号”，而非我们平常看到
的语言文法层面使用的符号。



goroutine机理
goroutine就是一种Go语言版本的协程（coroutine）
协程这个术语应该是随着Lua语言的流行而流行起来的，但要刨根究底的话，协
程第一次出现在1963年，用于汇编编程。
Lua和Go语言则可以算是最近几年在语言层
面支持协程的典型代表，但实际上支持协程的语言有三四十种之多，比如C#也在内部支持协程。


	协程  协程机理
协程，也有人称之为轻量级线程，具备以下几个特点。
能够在单一的系统线程中模拟多个任务的并发执行。
  在一个特定的时间，只有一个任务在运行，即并非真正地并行。
  被动的任务调度方式，即任务没有主动抢占时间片的说法。当一个任务正在执行时，外
部没有办法中止它。要进行任务切换，只能通过由该任务自身调用 yield() 来主动出让
CPU使用权。
每个协程都有自己的堆栈和局部变量。
每个协程都包含3种运行状态：挂起、运行和停止。停止通常表示该协程已经执行完成（包
括遇到问题后明确退出执行的情况），挂起则表示该协程尚未执行完成，但出让了时间片，以后
有机会时会由调度器继续执行。


	协程C实现
轻量级协程库 libtask ，这个库的下载地址为http://swtch.com/libtask/

	协程库
	 libtask 库实现了以下几个关键模块
		任务及任务管理
	  任务调度器
	  异步IO
	  channel


	任务


	任务调度



	上下文切换



	通信机制
因为协程原则上不会出现多线程编程中经常遇到的资源竞争问题，所以这个channel的数据结
构甚至在访问的时候都不用加锁（因为Go语言支持多CPU核心并发执行多个goroutine，会造成资
源竞争，所以在必要的位置还是需要加锁的）。

*/

/*
接口机理
曾经深入研究过C++语言中的虚函数以及函数重载原理的读者，可能对于C++中引入的虚表
和虚表指针还有深刻的印象。因为C++中并没有真正的接口，而只有纯虚函数和纯虚类，因此虚
函数的原理就可以认为是C++版本的接口原理。
接口的主要用法包含从类型赋值到接口、接口之间赋值和接口查询等


	类型赋值给接口



	接口查询




	接口赋值



*/
