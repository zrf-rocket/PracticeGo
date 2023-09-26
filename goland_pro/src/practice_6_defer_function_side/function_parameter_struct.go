package main

import (
	"fmt"
	"time"
)

type UserInformation struct {
	Author  string
	Age     int
	Blog    string
	Timeout time.Duration
}

// 将结构体作为函数参数进行传递
func UseParameter(userInfo UserInformation) {
	if userInfo.Author == "" {
		userInfo.Author = "SteveRocket"
	}
	if 0 == userInfo.Age {
		userInfo.Age = 28
	}
	if "" == userInfo.Blog {
		userInfo.Blog = "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"
	}
	if 0 == userInfo.Timeout {
		userInfo.Timeout = time.Second
	}
}

func main() {
	// 实例化结构体
	var user01 UserInformation
	UseParameter(user01)
	fmt.Println(user01.Author)
	fmt.Println(user01.Age)
	fmt.Println(user01.Blog)
	fmt.Println(user01.Timeout)

	var user UserInformation
	user = UserInformation{Author: "SteveRocket", Age: 28, Blog: "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"}
	UseParameter(user)
	fmt.Println(user.Author)
	fmt.Println(user.Age)
	fmt.Println(user.Blog)
	fmt.Println(user.Timeout)

	var user02 UserInformation
	user02 = UserInformation{Author: "SteveRocket", Age: 28, Blog: "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q", Timeout: 10000}
	UseParameter(user02)
	fmt.Println(user02.Author)
	fmt.Println(user02.Age)
	fmt.Println(user02.Blog)
	fmt.Println(user02.Timeout)
}
