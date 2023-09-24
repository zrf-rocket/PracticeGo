package main

import "fmt"
import "log"

/*
类型系统
	为类型添加方法
	值语义和引用语义
	结构体
初始化
匿名组合
可见性

*/

/*
面向对象的特性

Go语言并没有沿袭传统面向对象编程中的诸多概念，比如继承、虚函数、构造函数和析构函数、隐藏的 this 指
针等。优雅之处在于，Go语言对面向对象编程的支持是语言类型系统中的天然组成部分。
整个类型系统通过接口串联，浑然一体。

类型系统才是一门编程语言的地基
类型系统是指一个语言的类型体系结构。类型系统通常包含如下
基础类型，如 byte 、 int 、 bool 、 float 等；
复合类型，如数组、结构体、指针等；
可以指向任意对象的类型（ Any 类型） ；
值语义和引用语义；
面向对象，即所有具备面向对象特征（比如成员方法）的类型；
接口。

类型系统描述的是这些内容在一个语言中如何被关联。
在Java语言中， 存在两套完全独立的类型系统： 一套是值类型系统， 主要是基本类型， 如 byte 、int 、 boolean 、 char 、 double 等，这些类型基于值语义；一套是以 Object 类型为根的对象类型
系统，这些类型可以定义成员变量和成员方法，可以有虚函数，基于引用语义，只允许在堆上创建（通过使用关键字 new ） 。Java语言中的 Any 类型就是整个对象类型系统的根—— java.lang.Object
类型，只有对象类型系统中的实例才可以被 Any 类型引用。值类型想要被 Any 类型引用，需要装箱（boxing）过程，比如 int 类型需要装箱成为 Integer 类型。另外，只有对象类型系统中的类型才可
以实现接口，具体方法是让该类型从要实现的接口继承。

相比之下，Go语言中的大多数类型都是值语义，并且都可以包含对应的操作方法。在需要的时候，你可以给任何类型（包括内置类型） “增加”新方法。而在实现某个接口时，无需从
该接口继承（事实上，Go语言根本就不支持面向对象思想中的继承语法） ，只需要实现该接口要求的所有方法即可。
任何类型都可以被 Any 类型引用。
Any 类型就是空接口， 即 interface{} 。


“在Go语言中没有隐藏的 this 指针”这句话的含义是：
方法施加的目标（也就是“对象” ）显式传递，没有被隐藏起来；
方法施加的目标（也就是“对象” ）不需要非得是指针，也不用非得叫 this 。

如果要求对象必须以指针传递，这有时会是个额外成本，因为对象有时很小（比如4字节） ，用指针传递并不划算。

*/
func main() {
	// TypeAddFunc()
	// ValueSemanticsAndReferenceSemantics()
	Struct()

	a := new(Rect2)
	_ = a
	b := &Rect2{Width: 100, Height: 200}
	fmt.Println(b.area())
}

type Integer int

//为类型添加方法
// 可以给任意类型（包括内置类型，但不包括指针类型）添加相应的方法
// 定义了一个新类型 Integer ，它和 int 没有本质不同，只是它为内置的int 类型增加了个新方法 Less()
//面向对象的例子
func (a Integer) Less(b Integer) bool {
	return a < b
}
func TypeAddFunc() {
	var a Integer = 1
	// 面向对象的用法
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}

	// 面向过程的用法
	if IntegerLess(a, 2) {
		fmt.Println(a, "Less 2")
	}

	a.Add1(2)
	fmt.Println(a) //3

	a.Add2(3)
	fmt.Println(a) //3   维持原来的值   因为Go语言和C语言一样，类型都是基于值传递的。要想修改变量的值，只能传递指针。
}

//面向过程的例子
func IntegerLess(a Integer, b Integer) bool {
	return a < b
}

// 只有在你需要修改对象的时候，才必须用指针。它不是Go语言的约束，而是一种自然约束。
// 为 Integer 类型增加了 Add1() 方法。由于 Add1() 方法需要修改对象的值，所以需要用指针引用。
func (a *Integer) Add1(b Integer) {
	*a += b
}

// 实现成员方法时传入的不是指针而是值（即传入Integer ，而非 *Integer ）
func (a Integer) Add2(b Integer) {
	a += b
}

// 如 http 包中关于HTTP头部信息的 Header 类型就是通过Go内置的 map 类型赋予新的语义来实现的
// （参见$GOROOT/src/pkg/http/header.go）
// Header类型用于表达HTTP头部的键值对信息
// type Header map[string][]string
// Header 类型其实就是一个 map ， 但通过为 map 起一个 Header 别名并增加了一系列方法， 它就
// 变成了一个全新的类型，但这个新类型又完全拥有 map 的功能。

// Add()方法用于添加一个键值对到HTTP头部 如果该键已存在，则会将值添加到已存在的值后面
// func (h Header) Add(key, value string){
// 	textproto.MIMEHeader(h).Add(key, value)
// }

// Set()方法用于设置某个键对应的值，如果该键已存在，则替换已存在的值
// func (h Header) Set(key, value string){
// 	textproto.MIMEHeader(h).Set(key, value)
// }

// 值语义和引用语义
// 值语义和引用语义的差别在于赋值
/*
b=a
b.Modify()
如果 b 的修改不会影响 a 的值，那么此类型属于值类型。如果会影响 a 的值，那么此类型是引用类型。

Go语言中的大多数类型都基于值语义
基本类型，如 byte 、 int 、 bool 、 float32 、 float64 和 string 等
复合类型，如数组（array） 、结构体（struct）和指针（pointer）等


Go语言中有4个类型比较特别，看起来像引用类型
数组切片：指向数组（array）的一个区间。
map：极其常见的数据结构，提供键值查询能力。
channel：执行体（goroutine）间的通信设施。
接口（interface）：对一组满足某个契约的类型的抽象。

*/

// 数组切片本质上是一个区间，可以大致将 []T 表示为
type slice struct {
	// first *T
	first *string
	len   int
	cap   int
}

// 因为数组切片内部是指向数组的指针，所以可以改变所指向的数组元素并不奇怪。数组切片类型本身的赋值仍然是值语义。

// map 本质上是一个字典指针，可以大致将 map[K]V 表示为t
// type Map_K_V struct{
// 	//...
// }
// type map[K][V] struct{
// 	impl *Map_K_V
// }
// 基于指针，我们完全可以自定义一个引用类型
// type IntergerRef struct{
// 	impl *int
// }

// channel和map类似，本质上是一个指针。将它们设计为引用类型而不是统一的值类型的原因是，完整复制一个channel或map并不是常规需求。

// 接口具备引用语义，是因为内部维持了两个指针
// type interface struct{
// 	data *void
// 	itab *Itab
// }

func ValueSemanticsAndReferenceSemantics() {
	// Go语言中的数组和基本类型没有区别，是很纯粹的值类型
	var a = [3]int{11, 22, 33}
	var b = a
	fmt.Println(a, b) //[11 22 33] [11 22 33]
	b[1]++
	// 表明 b=a 赋值语句是数组内容的完整复制。要想表达引用，需要用指针
	fmt.Println(a, b) //[11 22 33] [11 23 33]
	b[2] = 100
	fmt.Println(a, b) //[11 22 33] [11 23 100]

	var a1 = [3]int{33, 44, 55}
	var b1 = &a1
	// 表明 b1=&a1 赋值语句是数组内容的引用。变量 b1 的类型不是 [3]int ，而是 *[3]int 类型
	fmt.Println(a1, *b1) //[33 44 55] [33 44 55]
	b1[1]++
	fmt.Println(a1, *b1) //[33 45 55] [33 45 55]
}

// Go语言的结构体（struct）和其他语言的类（class）有同等的地位，但Go语言放弃了包括继承在内的大量面向对象特性，只保留了组合（composition）这个最基础的特性。
// 组合甚至不能算面向对象特性，因为在C语言这样的过程式编程语言中，也有结构体，也有组合。组合只是形成复合类型的基础。
// 定义一个矩形类型
type Rect struct {
	x, y          float64
	width, height float64
}

// 定义成员方法 Area() 来计算矩形的面积
func (r *Rect) Area() float64 {
	return r.width * r.height
}

// 在Go语言中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以NewXXX 来命名，表示“构造函数”
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

func Struct() {
	// 如何创建并初始化 Rect 类型的对象实例
	// 未进行显式初始化的变量都会被初始化为该类型的零值，例如 bool 类型的零值为 false ， int 类型的零值为0， string 类型的零值为空字符串。
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}
	fmt.Println(rect1.Area()) //0
	fmt.Println(rect2.Area()) //0
	fmt.Println(rect3.Area()) //20000
	fmt.Println(rect4.Area()) //20000
}

// 匿名组合   go通过组合来实现继承
type Base struct {
	Name string
}

// 定义了一个 Base 类（实现了 Foo() 和 Bar() 两个成员方法）
func (base *Base) Foo() {}
func (base *Base) Bar() {}

// 定义了一个Foo 类，该类从 Base 类“继承”并改写了 Bar() 方法
//（该方法实现时先调用了基类的 Bar()方法）
type Foo struct {
	Base
}

// 在“派生类” Foo 没有改写“基类” Base 的成员方法时，相应的方法就被“继承”   调用 foo.Foo() 和调用 foo.Base.Foo() 效果一致
func (foo *Foo) Bar() {
	foo.Base.Bar()
}

// Go语言很清晰地告诉你类的内存布局是怎样的
// 在Go语言中你还可以随心所欲地修改内存布局
type Foo2 struct {
	//...//其他成员
	Base
}

// 还可以以指针方式从一个类型“派生”
type Foo3 struct { //这段Go代码仍然有“派生”的效果，只是 Foo3 创建实例的时候，需要外部提供一个 Base 类实例的指针。
	*Base
	//...
}

// 在C++ 语言中其实也有类似的功能，那就是虚基类

// 定义如下的类型，它匿名组合了一个 log.Logger 指针
type Job struct {
	Command string
	*log.Logger
	// 对于 Job 的实现者来说，他甚至根本就不用意识到 log.Logger 类型的存在，这就是匿名组合的
	// 魅力所在。在实际工作中，只有合理利用才能最大发挥这个功能的价值。
}

// 在合适的赋值后，我们在 Job 类型的所有成员方法中可以很舒适地借用所有 log.Logger 提供的方法。
func (job *Job) Start() {
	// job.Log("starting now...")

	// job.Log("started")
	// Job 例子，即使组合后调用的方式变成了 job.Log(...) ，但 Log 函数的接收者仍然是
	// log.Logger 指针，因此在 Log 中不可能访问到 job 的其他成员方法和变量。
}

// 需要注意的是，不管是非匿名的类型组合还是匿名组合，被组合的类型所包含的方法虽然都
// 升级成了外部这个组合类型的方法，但其实它们被组合方法调用时接收者并没有改变。

// 毕竟被组合的类型并不知道自己会被什么类型组合，当然就没法在实现方法时去使用那个未知的“组合者”的功能了

// 接口组合中的名字冲突问题
type X struct {
	Name string
}
type Y struct {
	X
	Name string
}

// 组合的类型和被组合的类型都包含一个 Name 成员，会不会有问题呢？答案是否定的。所有
// 的 Y 类型的 Name 成员的访问都只会访问到最外层的那个 Name 变量， X.Name 变量相当于被隐藏起来了。

type Logger struct {
	Level int
}

// 匿名组合类型相当于以其类型名称（去掉包名部分）作为成员变量的名字。 按此规则， Y 类型中就相当于存在两个名为 Logger 的成员，
// 虽然类型不同。因此，我们预期会收到编译错误。
type Y2 struct {
	*Logger
	Name string
	// *log.Logger  //duplicate field Logger
}

// 有意思的是，这个编译错误并不是一定会发生的。假如这两个 Logger 在定义后再也没有被
// 用过，那么编译器将直接忽略掉这个冲突问题，直至开发者开始使用其中的某个 Logger 。

// 可见性
// Go语言对关键字的增加非常吝啬，其中没有 private 、 protected 、 public 这样的关键
// 字。要使某个符号对其他包（package）可见（即可以访问） ，需要将该符号定义为以大写字母开头
type Rect2 struct {
	X, Y          float64
	Width, Height float64
}

// Rect 类型的成员变量就全部被导出了，可以被所有其他引用了 Rect 所在包的代码访问到。

// 成员方法的可访问性遵循同样的规则  Rect 的 area() 方法只能在该类型所在的包内使用
func (r *Rect2) area() float64 {
	return r.Width * r.Height
}

// 需要注意的一点是，Go语言中符号的可访问性是包一级的而不是类型一级的。在上面的例子中，尽管 area() 是 Rect 的内部方法，但同一个包中的其他类型也都可以访问到它。这样的可
// 访问性控制很粗旷，很特别，但是非常实用。如果Go语言符号的可访问性是类型一级的，少不了还要加上 friend 这样的关键字，以表示两个类是朋友关系，可以访问彼此的私有成员。
