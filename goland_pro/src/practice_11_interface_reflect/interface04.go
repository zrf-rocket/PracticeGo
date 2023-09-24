package main

import "fmt"

/*
接口
	其他语言接口
	非侵入式接口
	接口赋值
	接口查询
	类型查询
	接口组合
	Any类型 区别Scala中的Any类型
*/
func main() {
	Interface1()
	Interface2()
	Interface3()
	Interface4()
	Interface5()
	OtherLanguageInterface1()
	OtherLanguageInterface2()
	OtherLanguageInterface3()
	OtherLanguageInterface4()
	OtherLanguageInterface5()
	NonIntrusiveInterface1()
	NonIntrusiveInterface2()
	NonIntrusiveInterface3()
	NonIntrusiveInterface4()
	NonIntrusiveInterface5()
	InterfaceAssignment1()
	InterfaceAssignment2()
	InterfaceAssignment3()
	InterfaceAssignment4()
	InterfaceAssignment5()
	InterfaceQuery1()
	InterfaceQuery2()
	InterfaceQuery3()
	InterfaceQuery4()
	InterfaceQuery5()
	TypeQuery1()
	TypeQuery2()
	TypeQuery3()
	TypeQuery4()
	TypeQuery5()
	InterfaceCombination1()
	InterfaceCombination2()
	InterfaceCombination3()
	InterfaceCombination4()
	InterfaceCombination5()
	AnyType1()
	AnyType2()
	AnyType3()
	AnyType4()
	AnyType5()
}

// 接口
func Interface1() {
	/*
	   接口在Go语言有着至关重要的地位。如果说goroutine和channel 是支撑起Go语言的并发模型的基石，让Go语言在如今集群化与多核化的时代成为一道极为亮丽的风景，
	   那么接口是Go语言整个类型系统的基石，让Go语言在基础编程哲学的探索上达到前所未有的高度。
	*/

}

func Interface2() {

}

func Interface3() {

}

func Interface4() {

}

func Interface5() {

}

// 其他语言接口
func OtherLanguageInterface1() {
	/*
	   Go语言的接口并不是其他语言（C++、Java、C#等）中所提供的接口概念。在Go语言出现之前，接口主要作为不同组件之间的契约存在。
	   对契约的实现是强制的，你必须声明你的确实现了该接口。为了实现一个接口，你需要从该接口继承
	*/

}

func OtherLanguageInterface2() {

}

func OtherLanguageInterface3() {

}

func OtherLanguageInterface4() {

}

func OtherLanguageInterface5() {

}

/*
侵入式接口。“侵入式”的主要表现在于实现类需要明确声明自己实现了
某个接口。这种强制性的接口继承是面向对象编程思想发展过程中一个遭受相当多置疑的特性。

Go语言的非侵入式接口
其一，Go语言的标准库，再也不需要绘制类库的继承树图。
在Go中，类的继承树并无意义，你只需要知道这个类实现了哪些方法，每个方法是啥含义就足够了。

其二，实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才合理。接口由使用方按需定义，而不用事前规划。

其三，不用为了实现一个接口而导入一个包，因为多引用一个外部的包，就意味着更多的耦
合。接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口。
*/
// 非侵入式接口
func NonIntrusiveInterface1() {
	// 在Go语言中，一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口
	fmt.Println("接口")

	// 尽管 File 类并没有从这些接口继承，甚至可以不知道这些接口的存在，但是 File 类实现了这些接口，可以进行赋值
	var file1 IFile = new(File)
	var file2 IReader = new(File)
	var file3 IWriter = new(File)
	var file4 ICloser = new(File)
	_, _, _, _ = file1, file2, file3, file4

}

// 定义了一个 File 类
type File struct{}

// 并实现有 Read() 、 Write() 、 Seek() 、 Close() 等方法
func (f *File) Read(buf []byte) (n int, err error)                { return } //{return n, err}
func (f *File) Write(buf []byte) (n int, err error)               { return } //{return n, err}
func (f *File) Seek(off int64, whence int) (pos int64, err error) { return } //{return  pos, err}
func (f *File) Close() error                                      { return nil }

//第一个括号用于继承类  第二个括号是函数参数   第三个括号是函数返回值

//接口
type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}

type IReader interface {
	Read(buf []byte) (n int, err error)
}

type IWriter interface {
	Write(buf []byte) (n int, err error)
}

type ICloser interface {
	Close() error
}

func NonIntrusiveInterface2() {

}

func NonIntrusiveInterface3() {

}

func NonIntrusiveInterface4() {

}

func NonIntrusiveInterface5() {

}

/*
接口赋值在Go语言中分为如下两种情况：
将对象实例赋值给接口；
将一个接口赋值给另一个接口。
*/
type Integer int

// 接口赋值
func InterfaceAssignment1() {
	// 将某种类型的对象实例赋值给接口，这要求该对象实例实现了接口要求的所有方法
	var a Integer = 1
	var b LessAdder = &a //1语句  正确的采用方式
	// var c LessAdder = a  //2语句
	_ = b
}

//选择1  的原因在于，Go语言可以根据下面的函数
func (a Integer) Less(b Integer) bool {
	return a < b
}

// 自动生成一个新的 Less() 方法
// func (a *Integer) Less(b Integer) bool{  //这样，类型 *Integer 就既存在 Less() 方法，也存在 Add() 方法，满足 LessAdder 接口
// 	return (*a).Lass(b)
// }

//无法根据下面这个方法
func (a *Integer) Add(b Integer) {
	*a += b
}

//无法自动生成以下的成员方法
// func (a Integer) Add(b Integer){
// 	(&a).Add(b)
// }
// 因为 (&a).Add() 改变的只是函数参数 a ，对外部实际要操作的对象并无影响，这不符合用
// 户的预期。所以，Go语言不会自动为其生成该函数。因此，类型 Integer 只存在 Less() 方法，
// 缺少 Add() 方法，不满足 LessAdder 接口，故此上面的语句(2)不能赋值。

//定义接口
type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func InterfaceAssignment2() {
	// 定义一个 Integer 类型的对象实例，将其赋值给 Lesser 接口
	var a Integer = 1
	var b1 Lesser = &a //1语句
	var b2 Lesser = a  //2语句
	//上面两个语句均通过编译
	_, _ = b1, b2
}

type Lesser interface {
	Less(b Integer) bool
}

// 将一个接口赋值给另一个接口
// 在Go语言中，只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等同的，可以相互赋值。
func InterfaceAssignment3() {
	var a = new(ReadWriter)
	var b = new(IStream)
	fmt.Println(a)
	fmt.Println(b)
	// a = b

	var file1 IStream = new(File)
	fmt.Println(file1) //&{}
	var file2 ReadWriter = file1
	fmt.Println(file1, file2) //&{} &{}
	var file3 IStream = file2
	fmt.Println(file3) //&{}

	var file4 Writer = file1 //允许将父级(多的)赋值给子集(少的)

	var file5 Writer = new(File)
	// var file6 IStream = file5  //编译不通过
	// 原因是显然的： file5 并没有 Read() 方法。
	_, _ = file4, file5
}

//定义两个接口
type ReadWriter interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}
type IStream interface {
	Write(buf []byte) (n int, err error)
	Read(buf []byte) (n int, err error)
}

// 在Go语言中，这两个接口实际上并无区别，因为
// 任何实现了 one.ReadWriter 接口的类，均实现了 two.IStream ；
// 任何 one.ReadWriter 接口对象可赋值给 two.IStream ，反之亦然；
// 在任何地方使用 one.ReadWriter 接口与使用 two.IStream 并无差异。

// 接口赋值并不要求两个接口必须等价。如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A。
// 可以将上面的 one.ReadWriter 和 two.IStream 接口的实例赋值给 Writer 接口  但是反过来并不成立
// 只允许多的付给少的  少的不能付给多的
type Writer interface {
	Write(buf []byte) (n int, err error)
}

func InterfaceAssignment4() {

}

func InterfaceAssignment5() {

}

// 有办法让上面的 Writer 接口转换为 two.IStream 接口么？有。那就是接口查询语法
// 接口查询
func InterfaceQuery1() {
	// var file1 Writer = ...
	// if file5, ok := file1.(two.IStream); ok {
	// ...
	// }

	// 这个 if 语句检查 file1 接口指向的对象实例是否实现了 two.IStream 接口，如果实现了，则执行特定的代码。
	// 接口查询是否成功，要在运行期才能够确定。它不像接口赋值，编译器只需要通过静态类型检查即可判断赋值是否可行。

	// 在Windows下做过开发的人，通常都接触过COM，知道COM也有一个接口查询（QueryInterface）。
	// 是的，Go语言的接口查询和COM的接口查询非常类似，都可以通过对象（组件）的某个接口来查询对象实现的其他接口。
	// 在Go语言中，对象是否满足某个接口，通过某个接口查询其他接口，这一切都是完全自动完成的。
	// 在COM中实现接口查询的过程非常繁复，但接口查询是COM体系的根本。

	// 在Go语言中是 interface{} ，在COM中是 IUnknown

	// 可以询问接口它指向的对象是否是某个类型
	// if file6, ok := file1.(*File); ok {
	// ...
	// }

	// 这个 if 语句判断 file1 接口指向的对象实例是否是 *File 类型，如果是则执行特定代码。
	// 查询接口所指向的对象是否为某个类型的这种用法可以认为只是接口查询的一个特例。接口
	// 是对一组类型的公共特性的抽象

	// 在C++、Java、C# 等语言中，也有类似的动态查询能力，比如查询一个对象的类型是否继承
	// 自某个类型（基类查询），或者是否实现了某个接口（接口派生查询），但是它们的动态查询与Go的动态查询很不一样。
}

func InterfaceQuery2() {

}

func InterfaceQuery3() {

}

func InterfaceQuery4() {

}

func InterfaceQuery5() {

}

// 类型查询
func TypeQuery1() {
	// 在Go语言中，还可以更加直截了当地询问接口指向的对象实例的类型
	// 类型查询并不经常使用。它更多是个补充，需要配合接口查询使用

	// 对于内置类型， Println() 采用穷举法，将每个类型转换为字符串进行打印。对
	// 于更一般的情况，首先确定该类型是否实现了 String() 方法，如果实现了，则用 String() 方
	// 法将其转换为字符串进行打印。否则， Println() 利用反射功能来遍历对象的所有成员变量进行打印。
	// 是的，利用反射也可以进行类型查询

}

func TypeQuery2() {

}

func TypeQuery3() {

}

func TypeQuery4() {

}

func TypeQuery5() {

}

// 接口组合
func InterfaceCombination1() {
	// 之前介绍的类型组合一样，Go语言同样支持接口组合。

}

// ReadWriter接口将基本的Read和Write方法组合起来
// type ReadWriter interface{
// 	Reader
// 	Writer
// }
// 这个接口组合了 Reader 和 Writer 两个接口，它完全等同于如下写法
// type ReadWriter interface{
// 	Read(p []byte)(n int, err error)
// 	Write(p []byte)(n int, err error)
// }
// 因为这两种写法的表意完全相同： ReadWriter 接口既能做 Reader 接口的所有事情，又能做
// Writer 接口的所有事情。在Go语言包中，还有众多类似的组合接口，比如 ReadWriteCloser 、
// ReadWriteSeeker 、 ReadSeeker 和 WriteCloser 等。

// 可以认为接口组合是类型匿名组合的一个特定场景，只不过接口只包含方法，而不包含任何成员变量。

func InterfaceCombination2() {

}

func InterfaceCombination3() {

}

func InterfaceCombination4() {

}

func InterfaceCombination5() {

}

// Any类型
func AnyType1() {
	// 由于Go语言中任何对象实例都满足空接口 interface{} ，所以 interface{} 看起来像是可以指向任何对象的 Any 类型
	var v1 interface{} = 1                  // 将int类型赋值给interface{}
	var v2 interface{} = "interface golang" // 将string类型赋值给interface{}
	var v3 interface{} = &v2                // 将*interface{}类型赋值给interface{}
	var v4 interface{} = struct{ X int }{1}
	var v5 interface{} = &struct{ X int }{1}

	_, _, _, _, _ = v1, v2, v3, v4, v5
	// 当函数可以接受任意的对象实例时，我们会将其声明为 interface{}
}

func AnyType2() {

}

func AnyType3() {

}

func AnyType4() {

}

func AnyType5() {

}
