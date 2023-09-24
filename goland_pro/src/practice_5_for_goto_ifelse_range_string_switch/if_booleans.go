package main

import "fmt"

func main() {
	if true {
		fmt.Println("Hello SteveRocket") //Hello SteveRocket
	}

	/*
		Python中可以使用任何对象来做判断，Go中不能这么用
		if 1 {
		}
	*/

	bool1 := true
	if bool1 {
		fmt.Printf("https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q\n")
	} else {
		fmt.Printf("https://blog.csdn.net/zhouruifu2015/\n")
	}
	//	输出结果：https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q

	num := -10 // 输出：值小于零
	if num > 0 {
		fmt.Println("值大于零")
	} else if num < 0 {
		fmt.Println("值小于零")
	} else {
		fmt.Println("值为零")
	}
}
