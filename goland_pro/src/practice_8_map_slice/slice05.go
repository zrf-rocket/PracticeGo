package main
import "fmt"

func main(){
	Slice1()
	Slice2()
	Slice3()
	Slice4()
	Slice5()
	Slice6()
	Slice7()
}



//切片
func Slice1(){
// 内部机制和基础
	// slice 是一种可以动态数组，可以按我们的希望增长和收缩。它的增长操作很容易使用，因为有内建的 append 方法。
	// 我们也可以通过 relice 操作化简 slice。因为 slice 的底层内存是连续分配的，所以 slice 的索引，迭代和垃圾回收性能都很好。

	// slice 是对底层数组的抽象和控制。它包含 Go 需要对底层数组管理的三种元数据，分别是：
	// 1.指向底层数组的指针
	// 2.slice 中元素的长度
	// 3.slice 的容量(可供增长的最大值)

// 创建和初始化
	// 第一个方法是使用内建的函数 make。当我们使用 make 创建时，一个选项是可以指定 slice 的长度
	slice1 := make([]int, 5) //指定长度
	fmt.Println(slice1,len(slice1),slice1[0])  //[0 0 0 0 0] 5 0



	// 如果只指定了长度，那么容量默认等于长度。
	//指定长度和容量
	slice2 := make([]int, 3, 5)
// 可以访问3个元素，但是底层数组有5个元素。两个与长度不相干的元素可以被 slice 来用。
	// 新创建的 slice 同样可以共享底层数组和已存在的容量。

	fmt.Println(slice1)  //[0 0 0 0 0]
	fmt.Println(slice2)  //[0 0 0]
	// slice2.append({1,3,6,5})
	// append(slice2,{1,3,5,7,9})
// 当我们分别指定了长度和容量，我们创建的 slice 就可以拥有一开始并没有访问的底层数组的容量

	//根据切片索引  修改切片的值
	slice2[1] = 222
	slice2[len(slice2) - 1] = 999
	fmt.Println(slice2,len(slice2))  //[0 222 999] 3

	// slice2[4] = 90  //error  超过了切片的长度
	// slice2[3] = 90   //error  超过了切片的长度
	slice2[2] = 90  //切片的最后一个元素
	fmt.Println(slice2[:])  //[0 222 90]



// 不允许创建长度大于容量的 slice
	// slice3 := make([]int, 5, 3)  //error   len larger than cap in make([]int)



// 惯用的创建 slice 的方法是使用 slice 字面量。跟创建数组很类似，不过不用指定 []里的值。初始的长度和容量依赖于元素的个数
	// 创建一个字符串 slice    长度和容量都是 5
	slice4 := []string{"golang","scala","redis","memcache","mongo"}
	fmt.Println(slice4,"\t",len(slice4),"\t",slice4[0])  //[golang scala redis memcache mongo]      5       golang


	// 在使用 slice 字面量创建 slice 时有一种方法可以初始化长度和容量，那就是初始化索引
	// 创建一个字符串 slice   初始化一个有100个元素的空的字符串 slice
	slice5 := []string{99 : ""}
	slice6 := []string{99 : "golang"}
	fmt.Println(slice5, "\t", len(slice5), "\t", slice5[len(slice5)-1])
	fmt.Println(slice6, "\t", len(slice6), "\t", slice6[len(slice6)-1])
	// [                                                                                                   ]    100
	// [                                                                                                   golang]      100
}

// nil 和 empty slice
func Slice2(){
	// 创建一个 nil slice，创建一个 nil slice 的方法是声明它但不初始化它
	var slice []int //声明slice
	var slice1 = []int{1, 3, 5, 7}  //初始化
	// 使用 slice 字面值创建
	slice2 := []int{} //初始化     空的slice
	// 使用 make 创建
	slice3 := make([]int,0)


	fmt.Println(slice)  //[]
	fmt.Println(len(slice))  //0
	// fmt.Println(slice[0])  //panic: runtime error: index out of range

	// if not slice {   //error
	// if slice {  		//error
	if slice == nil {
		fmt.Println("切片Slice是空的")
	}else{
		fmt.Println("切片Slice是的内容：",slice)
	}
	if slice1 == nil {
		fmt.Println("切片Slice1是空的")
	}else{
		fmt.Println("切片Slice1是的内容：",slice1)
	}
	if slice2 == nil || 0 == len(slice2){    //!len(slice2)行不通
		fmt.Println("切片Slice2是空的")
	}else{
		fmt.Println("切片Slice2是的内容：", slice2)
	}
	if slice3 == nil || 0 == len(slice3){    //!len(slice2)行不通
		fmt.Println("切片Slice3是空的")
	}else{
		fmt.Println("切片Slice3是的内容：", slice3)
	}


	fmt.Println(slice2, len(slice2))
	// _ = slice2[0]  //error  panic: runtime error: index out of range
	fmt.Println(slice3, len(slice3))
	// _ = slice3[0]  //error  panic: runtime error: index out of range
// empty slice 包含0个元素并且底层数组没有分配存储空间。当我们想要表示一个空集合时它很有用处，比如一个数据库查询返回0个结果。
// 不管我们用 nil slice 还是 empty slice，内建函数 append，len和cap的工作方式完全相同。
}

//操作slice
func Slice3(){
	// 为一个指定索引值的 slice 赋值跟之前数组赋值的做法完全相同。改变单个元素的值使用 [] 操作符
	slice := []int{11,22,33,44,55,66}
	fmt.Println(slice,len(slice))   //[11 22 33 44 55 66] 6

	// 可以在底层数组上对一部分数据进行 slice 操作，来创建一个新的 slice
	// 长度为2，容量为4
	newSlice := slice[1:3]
	fmt.Println(newSlice,len(newSlice))  //[22 33] 2
	fmt.Println(&slice[0],&newSlice[0])  //0xc04203fd70 0xc04203fd78
// 在 slice 操作之后我们得到了两个 slice，它们共享底层数组。但是它们能访问底层数组的范围却不同，newSlice 不能访问它头指针前面的值。
// 必须再次明确一下现在是两个 slice 共享底层数组，因此只要有一个 slice 改变了底层数组的值，那么另一个也会随之改变
	slice[0] = 9999
	slice[1] = 1111
	fmt.Println(slice)   //[9999 1111 33 44 55 66]
	fmt.Println(newSlice)  //[1111 33]    newSlice 不能访问它头指针前面的值

	newSlice[0] = 0   //改变newSlice的值  slice的值也随之被改变
	fmt.Println(slice)  //[9999 0 33 44 55 66]

	// 计算任意 new slice 的长度和容量
	// 对于 slice[i:j] 和底层容量为 k 的数组
	// 长度：j - i
	// 容量：k - i



// 一个 slice 只能访问它长度范围内的索引，试图访问超出长度范围的索引会产生一个运行时错误。
	// 容量只可以用来增长，它只有被合并到长度才可以被访问
	slice2 := []int{11, 22, 33, 44, 55}
	newSlice2 := slice2[1:3]
	// newSlice2[3] = 45  //panic: runtime error: index out of range
	fmt.Println(slice2)
	fmt.Println(newSlice2)
}


// 容量可以被合并到长度里，通过内建的 append 函数。
func Slice4(){
	// slice 增长
// slice 比 数组的优势就在于它可以按照我们的需要增长，我们只需要使用 append 方法，然后 Go 会为我们做好一切。
// 使用 append 方法时我们需要一个源 slice 和需要附加到它里面的值。当 append 方法返回时，它返回一个新的 slice，
// append 方法总是增长 slice 的长度，另一方面，如果源 slice 的容量足够，那么底层数组不会发生改变，否则会重新分配内存空间。

	// 创建一个长度和容量都为5的 slice
	slice := []int{11, 22, 33, 44, 55}
	fmt.Println(len(slice))   //5
	fmt.Println(cap(slice))   //5
	fmt.Println(&slice[1])   //0xc04203fd78

// 创建一个新的 slice
	newSlice := slice[1:3]
	
// 为新的 slice append 一个值   此处的newSlice还会是原来的slice
	newSlice = append(newSlice,66)
	// 因为 newSlice 有可用的容量，所以在 append 操作之后 slice 索引为 3 的值也变成了 60，
	// 之前说过这是因为 slice 和 newSlice 共享同样的底层数组。
	fmt.Println(slice,len(slice))
	fmt.Println(newSlice,len(newSlice))
	fmt.Println(&newSlice[0])   //0xc04203fd78
	fmt.Println()
	

// 如果没有足够可用的容量，append 函数会创建一个新的底层数组，拷贝已存在的值和将要被附加的新值
	slice2 := []int{11,22,33,44}
	fmt.Println(slice2,&slice2[0])
// 附加一个新值到 slice，因为超出了容量，所以会创建新的底层数组
	newSlice2 := append(slice2,55)
	// append 函数重新创建底层数组时，容量会是现有元素的两倍(前提是元素个数小于1000)，
	// 如果元素个数超过1000，那么容量会以 1.25 倍来增长。
	fmt.Println(slice2,&slice2[0])
	//此时的newSlice2将会是一个新的slice
	fmt.Println(newSlice2,&newSlice2[0])   //[11 22 33 44 55] 0xc042044100
	fmt.Println()


// slice 的第三个索引参数
// slice 还可以有第三个索引参数来限定容量，它的目的不是为了增加容量，
	// 而是提供了对底层数组的一个保护机制，以方便我们更好的控制 append 操作
	source := []string{"apple", "orange", "plum", "banana", "grape"}
	fmt.Println(&source[0])

// 接着我们在源 slice 之上创建一个新的 slice
	slice3 := source[2:3:4]
	// 新创建的 slice 长度为 1，容量为 2，可以看出长度和容量的计算公式也很简单
	// 对于 slice[i:j:k]  或者 [2:3:4]
	// 长度： j - i       或者   3 - 2
	// 容量： k - i       或者   4 - 2
	fmt.Println(source, &source[0])   //[apple orange plum banana grape] 0xc04206a050
	fmt.Println(slice3, &slice3[0])  //[plum] 0xc04206a070

	// 同数组一样，另外两个内建函数 len 和 cap 分别返回 slice 的长度和容量
	fmt.Println(len(slice3),cap(slice3))  //1 2
	fmt.Println()
	// 如果我们试图设置比可用容量更大的容量，会得到一个运行时错误
	// slice4 := source[2:3:6]  //panic: runtime error: slice bounds out of range

	// 限定容量最大的用处是我们在创建新的 slice 时候限定容量与长度相同，
	// 这样以后再给新的 slice 增加元素时就会分配新的底层数组，而不会影响原有 slice 的值

	// 接着我们在源 slice 之上创建一个新的 slice
// 并且设置长度和容量相同
	slice4 := source[2:3:3]
	fmt.Println(slice4,&slice4[0])
	slice4 = append(slice4,"wiki")  //[plum] 0xc04206a070
	fmt.Println(slice4,&slice4[0])  //[plum wiki] 0xc0420425a0
	fmt.Println(source)  //[apple orange plum banana grape]
	// 如果没有第三个索引参数限定，添加 kiwi 这个元素时就会覆盖掉 banana
}


//append()   迭代 slice
func Slice5(){
// 内建函数 append 是一个变参函数，意思就是你可以一次添加多个元素
	s1 := []int{11, 22}
	s2 := []int{33, 44}
	fmt.Printf("%v\n",append(s1,s2...))  //[11 22 33 44]


// slice 也是一种集合，所以可以被迭代，用 for 配合 range 来迭代
	slice := []int{11,22,33,44,55}
	for index, value := range slice{
	// for _, value := range slice{      // 如果不需要索引值，可以使用 _ 操作符来忽略它
		fmt.Printf("Index:%d  Value:%d\n", index, value)
		fmt.Printf("Value:%d  Value-Addr:%X   ElemAddr:%X\n", value, &value, &slice[index])
		// value 变量的地址总是相同的因为它只是包含一个拷贝。如果想得到每个元素的真是地址可以使用 &slice[index]。
	}
	// 当迭代时 range 关键字会返回两个值，第一个是索引值，第二个是索引位置值的拷贝。
	// 注意：返回的是值的拷贝而不是引用，如果我们把值的地址作为指针使用，会得到一个错误


// range 总是从开始一次遍历，如果你想控制遍历的step，就用传统的 for 循环
	// for index := 2; index < len(slice); ++index{  //syntax error: unexpected ++, expecting expression
	// for index := 2; index < len(slice); index++{ //从指定的位置起始依次向后遍历  步长为1
	for index := 2; index < len(slice); index*=2{   //从指定的位置起始依次向后遍历  步长为2
		fmt.Printf("Index:%d Value:%d\n", index, slice[index])
	}


}


// 多维 slice
func Slice6(){
// 也是同数组一样，slice 可以组合为多维的 slice
	slice := [][]int{{11},{33,44,55},{77},{}}
	fmt.Println(slice,len(slice),cap(slice))  //[[11] [33 44 55] [77] []] 4 4
	fmt.Println(&slice[0][0], &slice[1][0])   //0xc04203c200 0xc0420423e0
// 需要注意的是使用 append 方法时的行为，比如我们现在对 slice[0] 增加一个元素
	slice[0] = append(slice[0],12)

// 那么只有 slice[0] 会重新创建底层数组，slice[1] 则不会
	fmt.Println(slice)
	fmt.Println(&slice[0][0], &slice[1][0])  //0xc04203c250 0xc0420423e0
}


// 在函数间传递 slice
func Slice7(){
// 在函数间传递 slice 是很廉价的，因为 slice 相当于是指向底层数组的指针，让我们创建一个很大的 slice 然后传递给函数调用它
	slice := make([]int,1e6)
	slice = foo(slice)
	fmt.Println(&slice[0])
}

func foo(slice []int) []int{
	return slice
}