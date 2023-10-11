package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	arr := [5]int{11, 22, 33, 44, 55}
	fmt.Println(arr, reflect.TypeOf(arr)) // [11 22 33 44 55] [5]int
	// 将数组转换为字符串
	arrStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), ","), "[]")
	fmt.Println("数组转换后的字符串：", arrStr, reflect.TypeOf(arrStr)) // 数组转换后的字符串： 11,22,33,44,55 string

	// 将切片转换为字符串
	slice := []string{"steverocket", "cramer", "zhangsan"}
	str := strings.Join(slice, ",")
	fmt.Println("切片转换后的字符串：", str) // 切片转换后的字符串： steverocket,cramer,zhangsan

	arrStr2 := strings.Replace(strings.Trim(fmt.Sprint(arr), "[]"), " ", ",", -1)
	sliceStr2 := strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", ",", -1)
	fmt.Println(arrStr2, reflect.TypeOf(arrStr2))     // 11,22,33,44,55 string
	fmt.Println(sliceStr2, reflect.TypeOf(sliceStr2)) // steverocket,cramer,zhangsan string
}
