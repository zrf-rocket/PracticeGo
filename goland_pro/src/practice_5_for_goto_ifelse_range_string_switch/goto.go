package main

import "fmt"

func gotoExpress() {
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}

func gotoExpress2() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			goto END
		}
	}
END:
	fmt.Println("程序结束")
}

func main() {
	//gotoExpress()
	gotoExpress2()
}
