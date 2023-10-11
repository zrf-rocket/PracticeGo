package main

import "fmt"

func sliceDemo() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片：", slice, len(slice), cap(slice)) // 原始切片： [1 2 3 4 5] 5 5

	// 调整切片的长度为前三个元素
	slice = slice[:3]
	fmt.Println("调整长度后的切片：", slice, len(slice), cap(slice)) // 调整长度后的切片： [1 2 3] 3 5

	// 调整切片的容量
	slice = slice[3:cap(slice)]
	fmt.Println("调整容量后的切片：", slice, len(slice), cap(slice)) // 调整容量后的切片： [4 5] 2 2

	newSlice2 := slice[1:2]
	fmt.Println(len(newSlice2), cap(newSlice2)) // 1 1

	newSlice3 := append(slice[:1], slice[len(slice)-1:]...)
	fmt.Println(len(newSlice3), cap(newSlice3)) // 2 2
}

func sliceDemo3() {
	str := "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"
	protocol := str[:5]
	url := str[8:]
	fmt.Println(str)      // https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q
	fmt.Println(protocol) // https
	fmt.Println(url)      // mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q

	newUrl := []byte(str) // 中文字符需要用[]rune(str)
	// 将https中的s换成空字符
	newUrl[4] = ' '
	fmt.Println(string(newUrl)) // http ://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q

	// 将https换成http
	fmt.Println(string(append(newUrl[:4], newUrl[5:]...))) // http://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q
}

func sliceDemo4() {
	// 声明一个空切片
	// 下面的代码中，切片slice已经初始化，但没有元素，因此不是nil
	slice := []int{}
	slice2 := make([]string, 0, 0)
	slice3 := make([]string, 0)
	slice4 := make([]string, 1, 1)
	var slice5 = []int{}
	// 声明一个空切片
	var slice6 []int

	fmt.Println(slice == nil, len(slice), slice)    // false 0 []
	fmt.Println(slice2 == nil, len(slice2), slice2) // false 0 []
	fmt.Println(slice3 == nil, len(slice3), slice3) // false 0 []

	fmt.Println(slice4 == nil, len(slice4), slice4) // false 1 []
	fmt.Println(slice5 == nil, len(slice5), slice5) // false 0 []

	// 从输出结果可以看出，可以使用nil来判断一个切片是否为空。但是，需要注意的是，只有未初始化的切片才是nil，已经初始化但没有元素的切片不是nil。
	fmt.Println(slice6 == nil, len(slice6), slice6) // true 0 []
	//	因此，要判断一个切片是否为空，需要结合判断切片的长度是否为0来使用。
}

func main() {
	//sliceDemo()
	//sliceDemo3()
	sliceDemo4()
}
