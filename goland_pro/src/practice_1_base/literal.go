package main

import "fmt"

/*
3种字面量
用于表示基础数据类型的各种字面量，如：浮点类型值的12E-3

用于构建各种自定义的复合数据类型的类型字面量，如：下面的字面量用于表示一个名称为Person的自定义结构体类型
	type Person struct{
		Name string
		Age uint8
		Address string
	}

用于表示复合数据类型的值的复合字面量
	会被用来构造类型Struct(结构体),Array(数组),Slice(切片),Map(字典)的值
	复合字面量一般由值的字面类型(指的是隶属于结构体、数组、切片或字典类型的自定义数据类型的名称)和由花括号括起来的复合元素的列表组成
*/

func main() {
	Literal()
}

type Person struct {
	Name    string
	Age     uint8
	Address string
}

func Literal() {
	//此复合字面量中的类型名称是Person 代表一个Person类型的值
	//对复合字面量的每次求值都会导致一个新的值被创建
	p := Person{Name: "SteveRocket", Age: 22, Address: "Beijing, China"}
	fmt.Println(p.Name)
	fmt.Println(p.Age)
	fmt.Println(p.Address)

	//此类的复合字面量中不允许为多个复合元素指定同一个字段名称或其他常量形式的键
	//结构体类型值
	//Person{Name:"SteveRocket", Age:22, Name:"Golang"}  //error Name字段重复  duplicate field name in struct literal: Name
	//字典类型值
	// _ = map[string]string{"Name":"Golang","Age":"23","Age":"2"} //error   此处Age字段重复

	_ = map[string]string{"Name": "Golang", "Age": "23"} //OK  由于是string 所以Age的键值都需要是字符串类型

	//切片类型值
	// _ = []string{Age:"0"} //error   由于是切片所以不能是Age 只能是数字
	_ = []string{0: "0", 1: "-1"} // ok
	// _ = []string{0:"0",1:"-1",0:"True"}  //error  键0不能有重复的
	_ = []string{0: "0", 1: "-1", 2: "True"}   // ok
	_ = []string{0: "0", 1: "-1", 222: "True"} // ok
}
