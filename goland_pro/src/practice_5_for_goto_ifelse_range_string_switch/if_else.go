package main

import "fmt"

func ifExpress() {
	var first int = 10 // first is 5 or greater
	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {
		fmt.Printf("first is between 0 and 5\n")
	} else {
		fmt.Printf("first is 5 or greater\n")
	}

	//var cond int   //cond is not greater than 10
	//cond := 6  // cond is not greater than 10
	//cond := 10 // cond is not greater than 10
	cond := 5
	if cond = 5; cond > 10 {
		fmt.Printf("cond is greater than 10\n")
	} else {
		fmt.Printf("cond is not greater than 10\n")
	}

	age1 := 18
	if age1 >= 18 {
		fmt.Println("成年") // 成年
	} else {
		fmt.Println("未成年")
	}
	// 等价下面的写法
	if age := 18; age >= 18 {
		fmt.Println("成年") // 成年
	} else {
		fmt.Println("未成年")
	}
}

func ifMultiExpress() {
	// 注意，Go语言中的条件判断不需要围绕括号，但是大括号是必需的。
	// 在Go语言中没有三元组，所以即使对于基本条件，也需要使用一个完整的if语句。
	if 7%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, " is negative")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}
}

func main() {
	ifExpress()
	ifMultiExpress()
}
