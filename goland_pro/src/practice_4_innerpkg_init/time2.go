package main

import (
	"fmt"
	"time"
	//"dynamic"
)

func main() {
	// Timestamp1()
	Timestamp2()
	Timestamp3()
	Timestamp4()
}

func Timestamp1() {
	//获取当前时间的时间戳
	t := time.Now().Unix()
	fmt.Println(t)

	//时间戳转为字符串时间
	fmt.Println(time.Unix(t, 0).String())

	//带纳秒的时间戳
	t = time.Now().UnixNano()
	fmt.Println(t)

	//基本格式的时间表示
	fmt.Println(time.Now().String())
	fmt.Println(time.Now().Format("2006year 01month 02day"))
	// fmt.Println(time.Now().Format("2001year 01month 01day")) //3005year 05month 05day
	//时间戳转为date日期类型

}

func Timestamp2() {
	//字符串时间转为时间戳
	//字符串时间转为日期类型

}

func Timestamp3() {
	//直接获取系统时间
	year := time.Now().Year()
	fmt.Println(year)

	month := time.Now().Month()
	month2 := time.Now().Month().String()
	fmt.Println(month)
	fmt.Println(month2)

	day := time.Now().Day()
	fmt.Println(day)

	fmt.Println()
	//引用数据库中时间数据
	// month = time.Unix(dynamic.UpdateTime/1000, 0).Month().String()
	// day = time.Unix(dynamic.UpdateTime/1000, 0).Day()
	// year = time.Unix(dynamic.UpdateTime/1000, 0).Year()

	// 其中dynamic.UpdateTime为从数据库中读取出来的时间字段，先转为Time类型，再去获取月份、日期等。
	// 当然也可以在其结构体中加几个字段（不带json），方便前端的显示。需要注意的是，year和day均为int类型，而month为string类型。
}

func Timestamp4() {

}
