package main

import "fmt"

/*
18种基本类型
	bool,string,byte,rune,int/uint,int8/uint8,int16/uint16,int32/uint32,int64/uint64,float32,float64,complex64,complex128
	除了bool,string其他的都是数值类型
	rune可以看做是uint32类型的一个别名

	基本类型的名称都属于预定义标识符   基本类型也可以称为预定义类型
8种复合类型
	Array,Struct,Function,Interface,Slice,Map(字典),Channel(通道),Pointer(指针)
	静态类型：在变量声明中示出的那个类型，绝大多数类型的变量只拥有静态类型，唯独接口类型
		接口类型同时具备静态和动态类型
	动态类型：在运行时与该变量绑定在一起的值的实际类型。这个实际类型可以是实现了这个接口类型的任何类型。接口类型的变量的动态类型可以在执行期间变化
		因为所有实现了这个接口类型的类型的值都可以被赋给这个变量。但是这个变量的静态类型永远只是他声明时被指定的那个类型，也就是说接口类型的变量的静态
		类型永远会是这个接口的类型。

    一个数组类型的变量的声明中的类型决定了在这个变量中可以存放哪一个类型的元素。
    即一个数组类型的潜在类型决定了在该类型的变量中可以存放哪一个类型的元素。
*/
func main() {
	Type()
}

//定义名为Book类型的结构体   潜在类型就是Book
type Book2 struct {
	Name       string
	ISBN       string
	Press      string
	PageNumber uint16
}

//声明一个自定义类型
type MyString string //可以把MyString看作string类型的一个别名类型  MyString类型的潜在类型就是string

//一个类型的潜在类型存在可传递性
type iString MyString

// iString的潜在类型与MyString类型相同   都是string类型

type MyStrings [3]string

// MyStrings的潜在类型不是 [3]string    [3]string既不是一个预定义类型，也不是一个由类型字面量构造的复合类型
//而是一个元素类型为string的数组类型   应该把[3]string类型的潜在类型作为类型MyStrings的潜在类型
//而[3]string的潜在类型是string类型   因此MyStrings的潜在类型就是string类型

func Type() {
	//变量声明
	var bookName string
	fmt.Println("bookName:", bookName) //什么也不输出

	var Age int = 22
	fmt.Println(Age) //22

	//如何将一个值从一个类型转为另一个类型
	fmt.Println(uint(123)) //会将字面量123转换为uint类型的值

	//初始化一个MyString类型的Hobe变量
	var Hobe MyString = "Pingpong"
	fmt.Println(Hobe)

	//类型转换  这种类型转换不会创建新的值
	fmt.Println(MyString("这是一个字符串"))
	fmt.Println(string(MyString("这是另一个字符串")))
}
