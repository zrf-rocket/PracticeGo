package main

import (
	"fmt"
)

func sum(a []int, result chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	result <- sum
}

// goroutine和channel是两个很重要的因素，因为这两者让并发(并行)变的如此简单
func main() {
	//int类型的切片a
	a := []int{1, 3, 4, -62, 7, 6, 5, 4, 32}
	fmt.Println(a)      // [1 3 4 -62 7 6 5 4 32]
	fmt.Println(len(a)) //9

	for i := 0; i < len(a); i++ {
		if a[i] < 0 {
			fmt.Println(a[i]) //输出小于0的数字
		}
	}

	//通过内置函数make创建了一个能接收和发送int类型的channel。
	result := make(chan int)
	// 值得注意的地方是channel只能通过Go语言内建的函数make(chan type)创建，
	// 其中type指明了该channel能传递的数据类型。
	fmt.Println(result)
	// 通过关键字go执行了两个goroutine，这两个goroutine的功能是分别计算切片a前半部分和后半部分的和。
	go sum(a[:len(a)/2], result)
	go sum(a[len(a)/2:], result)
	x, y := <-result, <-result
	fmt.Println(x, y, x+y)

	//main函数碰到go关键字，派发goroutine执行相应的函数后，立即返回执行剩余的代码，
	// 不会等待goroutine的返回。sum函数中，计算切片的和，然后将结果发送到channel。
	// 接下来main函数，从channel中获取结果，在这里，main函数会阻塞至直到能从channel result中接收到数据

	fmt.Println(13 / 2) //输出6

	/*
	   程序暂停 休眠sleep()


	   //for循环遍历range
	   	for x<- range(len(a)){
	   		fmt.Println(x)
	   	}
	*/
}
