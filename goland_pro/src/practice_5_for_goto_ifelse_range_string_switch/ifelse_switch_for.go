package main

import "fmt"

func main() {
	Condition1()
	Choose1()
	Cycle1()
	Cycle2()
	Jump1()
}

// 流程控制 process control
// 	条件()
func Condition1() {
	// 条件语句不需要使用括号将条件包含起来 ()
	if a := 111; a < 5 { //在 if 之后，条件语句之前，可以添加变量初始化语句，使用 ; 间隔；
		fmt.Println("符合条件")
	} else { //无论语句体内有几条语句，花括号 {} 都是必须存在的
		fmt.Println("不符合条件")
	}
	fmt.Println(example(0))  //5
	fmt.Println(example(11)) //11
}

//下面这条在go1.7中允许
// 在有返回值的函数中，不允许将“最终的” return 语句包含在 if...else... 结构中否则会编译失败
func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}

// 	选择
func Choose1() {
	num := 0  //0
	num = 1   //1
	num = 2   //2
	num = 3   //3 4 5
	num = 4   //3 4 5
	num = 5   //3 4 5
	num = 111 //不存在
	//case 的后面不需要像C++/java一样的break语句
	switch num {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3, 4, 5: //单个 case 中，可以出现多个结果选项
		fmt.Println("3 4 5")
	default:
		fmt.Println("不存在")
	}

	num = 6   //4~6
	num = -1  //不存在
	num = 0.0 //0~3
	// switch 后面的表达式甚至不是必需的
	// 可以不设定 switch 之后的条件表达式，在此种情况下，整个 switch 结构与多个if...else... 的逻辑作用等同。
	switch { //左花括号 { 必须与 switch 处于同一行
	case 0 <= num && num <= 3: //条件表达式不限制为常量或者整数
		fmt.Println("0~3")
	case 4 <= num && num <= 6:
		fmt.Println("4~6")
	case 7 <= num && num <= 8:
		fmt.Println("7~8")
	default:
		fmt.Println("不存在")
	}

	// 只有在 case 中明确添加 fallthrough 关键字，才会继续执行紧跟的下一个 case
	num = 0    //024
	num = 4    //4
	num = -0.0 //024
	switch num {
	case 0:
		fmt.Printf("0")
		fallthrough
	case 2:
		fmt.Printf("2")
		fallthrough
	case 4:
		fmt.Printf("4")
	case 6:
		fmt.Printf("6")
	default:
		fmt.Println("不存在")
	}
}

// 	循环
func Cycle1() {
	// Go语言中的循环语句只支持 for 关键字，而不支持 while 和 do-while结构

	// 无限循环
	// for{
	// 	fmt.Println("死循环")
	// }

	//使用break跳出死循环
	num := 0
	for {
		num += 1
		fmt.Println("循环", num)
		if num > 10 {
			break
		}
	}

	// 条件表达式中也支持多重赋值
	a := []int{11, 22, 33, 44, 55, 66, 77}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		fmt.Println("i=", i, "  j=", j)
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
}

func Cycle2() {
	fmt.Println("break JLoop")
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if i > 3 {
				fmt.Println(i, j)
				break
				// continue
				// break JLoop
			}
		}
	}
	// JLoop
	// fmt.Println("终止")
}

//goto跳转语句
func Jump1() {
	// goto 语句的语义非常简单，就是跳转到本函数内的某个标签
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}
