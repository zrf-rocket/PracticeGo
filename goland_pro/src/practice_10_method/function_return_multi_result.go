package main

import "fmt"

func getName() (firstName, middleName, lastName, nickName string) { //指定函数四个参数类型为string
	//返回多个值
	return "GoLang", "M", "rocket", "Love"
}

// 因为返回值都已经有名字，因此各个返回值也可以用如下方式来在不同的位置进行赋值，从而提供了极大的灵活性
func getName2() (firstName, middleName, lastName, nickName string) {
	//firstName := "Python" //这种赋值方式错误
	firstName = "Python"
	middleName = "M"
	lastName = "zrf"
	nickName = "Like"
	return
}

// 并不是每一个返回值都必须赋值，没有被明确赋值的返回值将保持默认的空值。而函数的调
// 用相比C/C++语言要简化很多
func getName3() (firstName, middleName, lastName, nickName string) {
	//middleName没有被赋值   默认返回空字符串
	firstName = "Scala"
	lastName = "rocket"
	nickName = "Lucky"
	return
}

func Compute(value1 int, value2 float64) (result float64, err error) {
	//返回值result和err均没有被赋值，分别默认返回 0 <nil>
	return
}

func main() {
	fmt.Println(getName())
	fmt.Println(getName2())
	fn, mn, ln, nn := getName3()
	fmt.Println("mn = ", mn)
	fmt.Println(fn, mn, ln, nn)

	// 可以直接用下划线作为占位符来忽略其他不关心的返回值。下面的调用表示调用者只希望接收 firstName 的值，
	// 这样可以避免声 明完全没用的变量
	firstName, _, _, _ := getName()
	fmt.Println(firstName)
	fmt.Println(Compute(111, 1.11))
}
