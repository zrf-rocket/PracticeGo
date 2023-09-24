package main

import "fmt"

func main() {
	ArrayTest()
}

func ArrayTest() {
	Array1()
	Array2()
	Array3()
	Array4()
	Array5()
	Array6()

	ArrayMap1()
	ArrayMap2()
	ArrayMap3()

	ArraySlice()
	ArraySlice2()
	ArraySlice3()
	ArraySlice4()
	ArraySlice5()
}

// 数组
// 数组长度在定义后就不可改变   在声明时长度可以为一个常量或者一个常量
// 表达式（常量表达式是指在编译期即可计算结果的表达式）
// 它包含相同类型的连续的元素，这些元素可以是内建类型，像数字和字符串，也可以是结构类型，元素可以通过唯一的索引值访问，从 0 开始。
// 数组是很有价值的数据结构，因为它的内存分配是连续的，内存连续意味着可是让它在 CPU 缓存中待更久，所以迭代数组和移动元素都会非常迅速。
func Array1() {
	// 声明一个长度为5的整数数组
	var array [5]int
	fmt.Println(array) //[0 0 0 0 0]
	// Go 语言中任何变量被声明时，都会被默认初始化为各自类型对应的 0 值，数组当然也不例外。
	// 当一个数组被声明时，它里面包含的每个元素都会被初始化为 0 值。

	// 快速创建和初始化数组的方法是使用数组字面值。数组字面值允许我们声明我们需要的元素个数并指定数据类型
	//var array1 := [5]int{11,22,44}//error
	// 声明一个长度为5的整数数组
	array1 := [5]int{11, 22, 44} //ok
	fmt.Println(array1)

	// var array3 []int{33,55,66,77,89,0}//error
	array[1] = 111
	fmt.Println(array)

	// 一旦数组被声明了，那么它的数据类型跟长度都不能再被改变。如果你需要更多的元素，
	// 那么只能创建一个你想要长度的新的数组，然后把原有数组的元素拷贝过去。
	// len({1,3,5,7,8,9,0})  /error

	// 把长度写成 ...，Go 编译器将会根据你的元素来推导出长度
	array4 := [...]string{"golang", "java", "scala", "python"}
	fmt.Println(len(array4)) //4
	// 数组的长度是该数组类型的一个内置常量

	// 如果我们知道想要数组的长度，但是希望对指定位置元素初始化
	// 为索引为1和4的位置指定元素初始化
	// 剩余元素为空值
	array5 := [5]string{1: "scala", 4: "java"}
	fmt.Println(array5)
}

//定义一个指针数组
func Array2() {
	array := [5]*int{0: new(int), 1: new(int)}
	fmt.Println(array)
	//为索引为0和1的元素赋值
	*array[0] = 11
	*array[1] = 1111
	fmt.Println(array)
	fmt.Println(*array[0]) //取出数组第一个元素

	// 遍历输出指针数组
	for i := 0; i < len(array); i++ {
		// fmt.Printf("\t",*(array[i]))
		fmt.Printf("\t", (array[i]))
	}
}

//数组的复制
func Array3() {
	// 在 Go 语言中数组是一个值，所以可以用它来进行赋值操作。一个数组可以被赋值给任意相同类型的数组：
	var array1 [5]string
	array2 := [5]string{"java", "golang", "python", "scala", "erlang"}
	array1 = array2
	fmt.Println(array1)
	fmt.Println(array2)

	// 注意数组的类型同时包括数组的长度和可以被存储的元素类型，数组类型完全相同才可以互相赋值
	var array3 [4]string
	//被复制的数组的长度不一样  所以无法被复制
	// array3 = array2  //cannot use array2 (type [5]string) as type [4]string in assignment
	var array4 [6]string
	// array4 = array2// 同样无法被复制

	//获取数组array2的长度  作为新声明的数组长度
	var array5 [len(array2)]string //这种做法是OK的   C++中不允许
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(array5)
	array5 = array2
	fmt.Println(array5)
}

//拷贝一个指针数组
func Array4() {
	// 拷贝一个指针数组实际上是拷贝指针值，而不是指针指向的值
	var array1 [5]*string
	array2 := [len(array1)]*string{new(string), new(string), new(string), new(string), new(string)}
	array1 = array2 //array1相当于array2的别名
	// 赋值完成后，两组指针数组指向同一字符串
	fmt.Println(array1)
	fmt.Println(array2)

	*array2[0] = "java"
	*array2[2] = "golang"
	*array2[4] = "scala"

	//对array2赋值  等共同于给array1赋值
	fmt.Println(*array1[0])
	fmt.Println(*array1[2])
	fmt.Println(*array1[4])
	fmt.Println(*array2[0])
	fmt.Println(*array2[2])
	fmt.Println(*array2[4])
}

// 多维数组
func Array5() {
	// 数组总是一维的，但是可以组合成多维的。多维数组通常用于有父子关系的数据或者是坐标系数据
	// 声明一个二维数组
	var array [][]int
	// array[0][0] = 0  //panic: runtime error: index out of range
	fmt.Println(array)      //[]
	fmt.Println(len(array)) // 0

	array2 := [][]int{{}, {}, {}, {}}
	fmt.Println(array2)      //[[] [] [] []]
	fmt.Println(len(array2)) //4

	// 使用数组字面值声明并初始化
	array3 := [][]int{{11, 22}, {33, 44}, {55, 66}, {77, 88}}
	fmt.Println(array3)

	// 指定外部数组索引位置初始化
	array4 := [][]int{1: {22, 33}, 5: {55, 66}}
	fmt.Println(array4, len(array4)) //[[] [22 33] [] [] [] [55 66]] 6

	// 同时指定内外部数组索引位置初始化
	array5 := [][]int{1: {0: 111}, 3: {1: 666}}
	fmt.Println(array5, len(array5)) //[[] [111] [] [0 666]] 4

	var array6 [4][3]int //4行3列
	array6[0][0] = 1
	array6[0][1] = 2
	array6[0][2] = 3
	array6[1][0] = 4
	array6[1][1] = 5
	array6[1][2] = 6
	fmt.Println(array6) //[[1 2 3] [4 5 6] [0 0 0] [0 0 0]]

	num := 10
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			num += 1
			array6[i][j] = num
		}
	}
	fmt.Println(array6) //[[11 12 13] [14 15 16] [17 18 19] [20 21 22]]

	// 同样的相同类型的多维数组可以相互赋值
	// var array7 [][]int //error  这样不行
	//同样需要相同类型和相同长度的数组
	var array7 [4][3]int //ok
	array7 = array6
	fmt.Println(array7)
	fmt.Println(array7[1])

	// 因为数组是值，我们可以拷贝单独的维
	var array8 [3]int
	array8 = array7[1]
	fmt.Println(array8)

	var value int = array7[0][0]
	fmt.Println(value)
}

// 在函数中传递数组
func Array6() {
	// 在函数中传递数组是非常昂贵的行为，因为在函数之间传递变量永远是传递值，所以如果变量是数组，那么意味着传递整个数组，即使它很大
	// 创建一个有百万元素的整形数组，在64位的机器上它需要8兆的内存空间，来看看我们声明它和传递它时发生了什么
	var array [1e6]int
	foo(array)
	// 每一次 foo 被调用，8兆内存将会被分配在栈上。一旦函数返回，会弹栈并释放内存，每次都需要8兆空间

	// 更好的方法来在函数中传递数组，那就是传递指向数组的指针，这样每次只需要分配8字节内存
	fooPointer(&array)
	// 但是注意如果你在函数中改变指针指向的值，那么原始数组的值也会被改变。幸运的是 slice(切片)可以帮我们处理好这些问题
}
func foo(array [1e6]int) {
	// fmt.Println(array)
}

// Go不支持函数重载  foo redeclared in this block
func fooPointer(array *[1e6]int) {
	// fmt.Println(array)
}

func ArrayMap1() {
	//数组声明

	//长度为32的数组  每个元素为一个字节
	arr := [32]byte{0}
	arrLength := len(arr)
	fmt.Println(arr) //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(arrLength)

	arr1 := [12]byte{1, 2, 3, 4}
	fmt.Println(arr1) //[1 2 3 4 0 0 0 0 0 0 0 0]

	// arr2 := [10]byte{1, _, 3, _, 5, _, 7} //此处不能使用占位符 _

	//复杂类型数组
	//[2 * N] struct {x, y int32}

	//指针数组
	// [1000]*float64

	//二维数组   3行5列的二维整型数组
	// [3][5]int

	//三维数组
	// [2][2][2]float64
	//上下等同
	//[2]{[2]([2]float64)}

}

func ArrayMap2() {
	//元素访问   可以使用数组下标来访问数组中的元素。数组下标从0开始
	arr := [5]int{11, 33, 55, 77, 88}
	//循环遍历一维数组
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d\t", arr[i])
	}

	fmt.Println("\n最后一个元素：", arr[len(arr)-1])

	//使用range遍历数组
	for index, value := range arr {
		fmt.Println("Array element[", index, "] = ", value)
	}
	// range 具有两个返回值，第一个返回值是元素的数组下标，第二个返回值是元素的值
}

func ArrayMap3() {
	// 值类型  数组的副本机制
	// 数组是一个值类型（value type）。所有的值类型变量在赋值和作为参数传递时都将产生一次复制动作。如果将数组作为函数的参数类型，则在函数调用时该
	// 参数将发生数据复制。因此，在函数体中无法修改传入的数组的内容，因为函数内操作的只是所传入数组的一个副本。

	//定义并初始化数组
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("修改之前：", arr) //[1 2 3 4 5]
	modify(arr)
	fmt.Println("修改之后：", arr) //[1 2 3 4 5]
}

//此处函数的参数中  数组的长度要与传入的数组的长度相等
func modify(array [5]int) {
	array[0] = 111                                  //试图修改array的值
	fmt.Println("In modify() array values:", array) //[111 2 3 4 5]
}

// 如何才能在函数内操作外部的数据结构呢？
// 数组的长度在定义之后无法再次修改；
// 数组是值类型，每次传递都将产生一份副本。
func ArraySlice() {
	// 数组切片的数组结构可以抽象成以下3个变量
	// 一个指向原生数组的指针；
	// 数组切片中的元素个数；
	// 数组切片已分配的存储空间。

	// 创建数组切片  基于数组和直接创建
	// 数组切片可以只使用数组的一部分元素或者整个数组来创建，甚至可以创建一个比所基于的数组还要大的数组切片
	var myArray = [10]int{} // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(myArray, len(myArray))

	var myArray2 [10]int = [10]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 100}
	fmt.Println(myArray2)
	fmt.Println("Elements of myArray2:")
	for _, value := range myArray2 {
		fmt.Printf("%d\t", value)
	}
	fmt.Println("\nElements of mySlice:")
	//基于数组的前5个元素创建一个数组切片
	var mySlice []int = myArray2[:5]
	for _, value := range mySlice {
		fmt.Printf("%d\t", value)
	}
	fmt.Println()
	fmt.Println("切片第一个元素：", mySlice[0], "\t切片长度：", len(mySlice))
	fmt.Println()

	//基于数组所有元素创建数组切片
	mySlice = myArray2[:]
	fmt.Println(mySlice)

	//基于数组后面的元素创建数组切片
	mySlice = myArray2[5:] //下标5(包括)开始
	fmt.Println(mySlice)
}

func ArraySlice2() {
	//直接创建
	// 内置函数 make() 可以用于灵活地创建数组切片

	// 创建一个初始元素个数为5的数组切片，元素初始值为0
	mySlice := make([]int, 5)
	fmt.Println(mySlice) //[0 0 0 0 0]

	// 创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
	mySlice2 := make([]int, 5, 10)
	fmt.Println(mySlice2, len(mySlice2)) //[0 0 0 0 0]  5
	// 注意此处len(mySlice2)输出5 表示当前数组所存储的元素个数  而不是预留的10个元素个数

	// 直接创建并初始化包含5个元素的数组切片
	mySlice3 := []int{11, 22, 33, 44, 55} //事实上还会有一个匿名数组被创建出来
	fmt.Println(mySlice3)
}

func ArraySlice3() {
	//动态增减元素
	// 可动态增减元素是数组切片比数组更为强大的功能。与数组相比，数组切片多了一个存储能力（capacity）的概念，即元素个数和分配的空间可以是两个不同的值。
	// cap() 函数返回的是数组切片分配的空间大小，而 len() 函数返回的是数组切片中当前所存储的元素个数。
	mySlice := make([]int, 5, 10)
	fmt.Println(&mySlice[0])  //0xc04203a0a0
	fmt.Println(mySlice)      //[0 0 0 0 0]
	fmt.Println(len(mySlice)) //5
	fmt.Println(cap(mySlice)) //10
	fmt.Println()

	// 继续新增元素，可以使用 append() 函数
	// mySlice.append(111)    error 用法错误
	// mySlice.append(mySlice, 111, 333)  error 用法错误

	// 此处生成一个新的数组切片
	// mySlice = append(mySlice, 111, 333)  //OK
	mySlice1 := append(mySlice, 111, 333)
	fmt.Println(&mySlice1[0])  //0xc04203a0a0
	fmt.Println(mySlice1)      //[0 0 0 0 0 111 333]
	fmt.Println(len(mySlice1)) //7
	fmt.Println(cap(mySlice1)) //10
	fmt.Println()

	// 函数 append() 的第二个参数其实是一个不定参数，我们可以按自己需求添加若干个元素
	// 甚至直接将一个数组切片追加到另一个数组切片的末尾
	mySlice2 := []int{8, 9, 10}
	// 因为 mySlice 中的元素类型为 int ，所以直接传递 mySlice2 是行不通的
	// mySlice = append(mySlice, mySlice2)  //error

	mySlice = append(mySlice, mySlice2...) //ok
	//等价于
	// mySlice = append(mySlice, 8, 9, 10)
	//加上省略号相当于把 mySlice2 包含的所有元素打散后传入
	fmt.Println(mySlice) //[0 0 0 0 0 8 9 10]

	// mySlice = append(mySlice,[1 2 3 4 5]...)  //error 错误的用法
	// mySlice = append(mySlice,[1,2,3,4,5]...)  //error 错误的用法
	// fmt.Println(map([1, 2, 3, 4, 5]))  //error 错误的用法

	// 数组切片会自动处理存储空间不足的问题。如果追加的内容长度超过当前已分配的存储空间
	fmt.Println(cap(mySlice)) //10
	mySlice = append(mySlice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 6)
	fmt.Println(mySlice)      //[0 0 0 0 0 8 9 10 1 2 3 4 5 6 7 8 9 1 1 1 2 2 3 3 4 4 5 6]
	fmt.Println(cap(mySlice)) //28

}

func ArraySlice4() {
	// 基于数组切片创建数组切片
	//类似于数组切片可以基于一个数组创建，数组切片也可以基于另一个数组切片创建
	oldSlice := []int{11, 22, 33, 44, 55}
	newSlice := oldSlice[:3] // 基于oldSlice的前3个元素构建新数组切片
	// newSlice2 := oldSlice[:20]  // newSlice 中超出oldSlice 出错

	//下面三种输出结果相同
	newSlice3 := oldSlice[:5]
	newSlice4 := oldSlice[:len(oldSlice)]
	newSlice5 := oldSlice[:]

	fmt.Println(oldSlice)
	fmt.Println(newSlice)
	fmt.Println(newSlice3)
	fmt.Println(newSlice4)
	fmt.Println(newSlice5)

	//通过切片索引   修改切片中的元素
	oldSlice[0] = 11111
	fmt.Println(oldSlice, "\n")
}

func ArraySlice5() {
	// 内容复制
	// 内置函数 copy() ，用于将内容从一个数组切片复制到另一个数组切片。如果加入的两个数组切片不一样大，就会按其中较小的那个数组切片的元素个数进行复制。
	slice1 := []int{11, 22, 33, 44, 55}
	slice2 := []int{2, 3, 4}
	// copy(slice2,slice1)  // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice1, *&slice1[0], &slice1[0]) //*&slice1[0]根据一个元素的地址取出元素    &slice1[0]取出第一个元素的地址
	// fmt.Println(*(&slice1[0]+4))  //Golang不支持指针操作
	fmt.Println(slice2, &slice2[0])

	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1, &slice1[0])
	fmt.Println(slice2, &slice2[0])
}
