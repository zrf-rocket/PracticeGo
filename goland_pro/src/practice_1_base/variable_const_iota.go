package main

import "fmt"

// 在const关键字出现时将被重置为0；
// const中每增加一行常量声明，将使 iota 计数一次
func iotaFunc() {
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Println(Monday)    // 输出：1
	fmt.Println(Tuesday)   // 输出：2
	fmt.Println(Wednesday) // 输出：3
	fmt.Println(Thursday)  // 输出：4
	fmt.Println(Friday)    // 输出：5
	fmt.Println(Saturday)  // 输出：6
	fmt.Println(Sunday)    // 输出：7
}

func iotaFunc2() {
	const (
		KB = 1 << (10 * iota) // 1左移10位，即1024
		MB                    // 1左移20位，即1024*1024
		GB                    // 1左移30位，即1024*1024*1024
		TB                    // 1左移40位，即1024*1024*1024*1024
	)

	fmt.Println(KB) // 输出：1
	fmt.Println(MB) // 输出：1024
	fmt.Println(GB) // 输出：1048576
	fmt.Println(TB) // 输出：1073741824
}

func iotaFunc3() {
	const (
		ReadPermission    = 1 << iota // 1左移0位，即1
		WritePermission               // 1左移1位，即2
		ExecutePermission             // 1左移2位，即4
	)
	var permission int = ReadPermission | WritePermission

	fmt.Println(permission&ReadPermission == ReadPermission)       // 输出：true
	fmt.Println(permission&WritePermission == WritePermission)     // 输出：true
	fmt.Println(permission&ExecutePermission == ExecutePermission) // 输出：false
}

func iotafFunc4() {
	//使用_跳过某些值
	const (
		num1 = iota
		_    // _也占了一行，所以_的值相当于是1
		num3
		_
		num5
	)
	// 此处又重新了const  跳跃iota的赋值
	const (
		num6 = iota
		num7       // num6=iota，所以num6的值为0；num7没有= 所以值递推为1
		num8 = 123 //  因为num8=123，而num9没有=，所以和num9的值保持一致都是123
		num9
		num10 = iota
		num11
	)
	fmt.Println(num1, num3, num5, num6, num7, num8, num9, num10, num11) // 0 2 4 0 1 123 123 4 5

	//多个iota定义在一行
	const (
		n1, n2 = iota + 1, iota + 2
		n3, n4
		n5, n6
		n7, n8
		// n9, n10, n11  错误用法，每行变量的个数必须跟第一排保持一致
		// n9, n10, n11, n12 //missing init expr for n11
	)
	fmt.Println(n1, n2, n3, n4, n5, n6, n7, n8) // 1 2 2 3 3 4 4 5

	const (
		n11, n12 = iota + 1, iota + 1
		n13, n14 = iota + 1, iota + 1
	)
	fmt.Println(n11, n12, n13, n14) // 1 1 2 2
}

func main() {
	//iotaFunc()
	//iotaFunc2()
	//iotaFunc3()
	iotafFunc4()
}
