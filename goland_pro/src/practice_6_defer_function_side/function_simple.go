package main

import "fmt"

func main() {
	var i1 int = MultiPly3Nums(2, 5, 6)
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", i1)                     // Multiply 2 * 5 * 6 = 60
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6)) // Multiply 2 * 5 * 6 = 60

	showHello()                                                                // this is SteveRocket
	showMe("SteveRocket", "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q") // my name is: SteveRocket  this is my blog: https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q

	fmt.Printf("func1 = %T\n", func1) // func1 = func(int, int, int, string)
	fmt.Printf("func2 = %T\n", func2) // func2 = func(int, int, int, string)
	fmt.Printf("func3 = %T\n", func3) // func3 = func(int, int) int
	fmt.Printf("func4 = %T\n", func4) // func4 = func(int, int) int
	fmt.Printf("func5 = %T\n", func5) // func5 = func(int, int) int
	fmt.Printf("func6 = %T\n", func6) // func6 = func(int, int) int
	fmt.Printf("func7 = %T\n", func7) // func7 = func(int, int) int

	fmt.Println(func8(11, 22)) // 33
}

// 函数的定义
func MultiPly3Nums(a int, b int, c int) int {
	// var product int = a * b * c
	// return product
	// 等价下面一行语句
	return a * b * c
}

// 没有参数也没有返回值
func showHello() {
	fmt.Println("this is SteveRocket")
}

// 有参数无返回值的函数
func showMe(name, blog string) {
	fmt.Println("my name is:", name, " this is my blog:", blog)
}

// 多个相同类型的参数简写方式：当参数的类型一致时，可以将连续相同类型的参数后面的类型省略
func func1(num1, num2, num3 int, name string) {}

// 等价于下面的定义方式
func func2(num1 int, num2 int, num3 int, name string) {}

func func3(num1 int, num2 int) int { return num1 + num2 }

// 下面声明了5个 拥有2个int型参数和1个int型返回值的函数 方法
// 返回值可以命名也可以不命名
func func4(num1, num2 int) int          { return num1 + num2 }
func func5(num1, num2 int) (result int) { return num1 + num2 }

// 下划线标识符 _ 可以表示某个参数未被使用
func func6(num1 int, _ int) int { return num1 }
func func7(int, int) int        { return 0 }

// 命名的返回值就相当于在函数中声明一个变量
func func8(num1, num2 int) (result int) {
	// 注意：因为已经在返回值中声明了ret，所以这里用= 而不是:= ,避免重复声明问题
	// result := num1 + num2   // 会报错 no new variables on left side of :=

	result = num1 + num2
	// return result
	// 因为已经在函数体中声明了result，所以在return的时候不需要重复声明，等价于下面的直接return
	return
}
