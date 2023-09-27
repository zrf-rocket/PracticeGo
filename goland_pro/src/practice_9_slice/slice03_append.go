package main

import "fmt"

func sliceReferenceType() {
	// 定义数据
	arr1 := [...]int{11, 22, 33, 44, 55, 66}
	fmt.Println(arr1) // [11 22 33 44 55 66]

	//由数组切割成切片
	s := arr1[2:]
	fmt.Println(s) //[33 44 55 66]

	//切片再次切片
	s2 := s[1:]
	fmt.Println(s2) // [44 55 66]

	//修改原始数组，把下标为3的值由44改为9999
	arr1[3] = 9999
	fmt.Println(arr1) // [11 22 33 9999 55 66]
	// 打印输出切片s、s2的结果即可看到元素的值均由44改为9999
	fmt.Println(s)                                       // [33 9999 55 66]
	fmt.Println(s2)                                      // [9999 55 66]
	fmt.Printf("arr1=%p  s=%p  s2=%p\n", &arr1, &s, &s2) // arr1=0xc0000c0060  s=0xc000096060  s2=0xc000096090

	// 重新定义一个slice
	s1 := []int{}
	fmt.Println(s1) // []
	// 去数组切片并重新赋值给slice
	s1 = arr1[:]
	fmt.Println(s1) // [11 22 33 9999 55 66]

	s3 := make([]int, 3, 3)
	s3 = []int{11, 22, 33, 44, 55}

	s4 := s3
	fmt.Println(s3, s4) // [11 22 33 44 55] [11 22 33 44 55]
	// 改变slice s3的值，s4的值也跟着变，说明s3、s4都指向同一底数组
	s3[0] = 900
	fmt.Println(s3, s4)                    // [900 22 33 44 55] [900 22 33 44 55]
	fmt.Printf("s3=%p  s4=%p\n", &s3, &s4) // s3=0xc000096120  s4=0xc000096138
}

func sliceAppend() {
	var slice1 = []int{11, 22, 33, 44, 55}
	fmt.Printf("%v\n", slice1) // [11 22 33 44 55]

	slice2 := []int{66, 77, 88, 99, 10}
	fmt.Println(slice2) // [66 77 88 99 10]

	//将两个slice合并成一个新的slice
	newSlice := append(slice1, slice2...)
	fmt.Println(newSlice) //[66 77 88 99 10]

	//向 slice 尾部添加数据，返回新的 slice 对象
	newSlice2 := append(newSlice, 1, 2, 3, 4)
	fmt.Println(newSlice2)                                               // [11 22 33 44 55 66 77 88 99 10 1 2 3 4]
	fmt.Printf("%p %p %p %p\n", &slice1, &slice2, &newSlice, &newSlice2) // 0xc000004078 0xc0000040a8 0xc0000040d8 0xc000004108

	slice3 := make([]int, 0, 5)
	//在切片末尾添加5个元素
	newSlice3 := append(slice3, 1, 2, 3, 4, 5)
	fmt.Println(slice3)    // []
	fmt.Println(newSlice3) // [1 2 3 4 5]

	slice3 = append(slice3, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(slice3)                     // [1 2 3 4 5 6 7 8]
	fmt.Printf("%p %p\n", &slice3, &slice3) // 0xc000096120 0xc000096120
	fmt.Println(newSlice3)                  // [1 2 3 4 5]

	numbers := make([]int, 1, 1)
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d 切片长度：%d  切片容量：%d\n", i, len(numbers), cap(numbers))
	}
	fmt.Println(numbers) // [0 0 1 2 3 4 5 6 7 8 9]

	address := []string{"北京", "上海", "广州", "深圳"}
	fmt.Println("长度：", len(address), "容量：", cap(address)) // 长度： 4 容量： 4
	address = append(address, "武汉")
	fmt.Println("长度：", len(address), "容量：", cap(address)) // 长度： 5 容量： 8
}

func sliceAppend2() {
	s := make([]string, 0, 1)
	caps := cap(s)
	for i := 0; i < 60; i++ {
		s = append(s, "a"+string(i))

		if n := cap(s); n > caps {
			fmt.Println("cap=", caps, n)
			caps = n
		}
	}

	dataArr := [...]int{0, 1, 2, 3, 4, 10: 0}
	s2 := dataArr[:2:3]
	s2 = append(s2, 111, 222)
	fmt.Println(s2, dataArr)         // [0 1 111 222] [0 1 2 3 4 0 0 0 0 0 0]
	fmt.Println(&s2[0], &dataArr[0]) // 0xc000120060 0xc000140000
}

func main() {
	//sliceReferenceType()
	//sliceAppend()
	//sliceAppend2()
	//sliceDelete()
	//sliceCopy()
	sliceResize()
}
