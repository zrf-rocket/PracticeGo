package main

/*
1. package
package关键字用于定义一个包，一个包可以包含多个相关的Go文件。每个Go文件都必须属于一个包，包名通常与文件名相同。在一个包中，可以定义变量、函数、类型等。

2. import
import关键字用于导入一个包，以便在代码中使用该包中的函数、变量等。在Go语言中，导入的包可以是标准库中的包，也可以是自己编写的包。
*/
import (
	"errors"
	"fmt"
	"runtime"
)


var CSDN_URL = "https://blog.csdn.net/zhouruifu2015/"
var WEIXIN_URL = "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"
var GIT_URL = "https://gitee.com/SteveRocket/practice_python.git"


/*
5. const
const关键字用于定义一个常量，常量是在程序运行时不可修改的值。常量可以是数值、布尔值、字符串等。
*/

const PI = 3.1415926
const PI2 int = 3
const (
		CSDN_URL01 = "https://blog.csdn.net/zhouruifu2015/"
		WEIXIN_URL01 = "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"
		GIT_URL01 = "https://gitee.com/SteveRocket/practice_python.git"
)

func const_kw(){
	CSDN_URL = WEIXIN_URL
	fmt.Println(CSDN_URL)
	
	// 不允许对常量赋值
	//CSDN_URL01 = WEIXIN_URL //错误用法 cannot assign to CSDN_URL01
}


/*
3. func
func关键字用于定义一个函数，函数是Go语言中的基本代码块。函数可以有参数和返回值，也可以没有参数和返回值。
*/
func break_kw(){
	for i := 0; i < 10; i++ {
		if i==5 {
			fmt.Println(WEIXIN_URL)
			break
		}
		fmt.Println(i)
	}
}


func case_kw(){
	/*
	4. var
	var关键字用于定义一个变量，变量可以是基本类型、结构体、数组、切片等。变量可以在函数内部或包级别定义。
	*/
	//var URL = "WEIXIN_URL"
	URL := "DOMAIN"
	switch URL {
	case "CSDN_URL":
		fmt.Println(CSDN_URL)
	case "WEIXIN_URL":
		fmt.Println(WEIXIN_URL)
	case "GIT_URL":
		fmt.Println(GIT_URL)
	//default关键字用于在switch语句中匹配一个默认条件，并执行相应的语句块。
	default:
		fmt.Println("is not url patterns")
	}
}


func chan_kw(age int){
	ch:=make(chan int)
	go func(){
		ch <- age
	}()
	fmt.Println(<-ch)
}


func continue_kw(){
	fmt.Println(WEIXIN_URL01)
	// 遍历字符串常量
	for _, v := range WEIXIN_URL01{
		// 遇到 . 则跳过  实现过滤指定字符
		//fmt.Println(string(rune(v)))
		if (string(rune(v)) == ".") {
			//fmt.Println("跳过点符号", v)
			continue
		}
		fmt.Printf("%c",v)
	}


	const count int = 20
	for i := 0; i < count; i++ {
		if(i % 5 == 0){
			fmt.Println("整除,跳过", i)
			continue
		}
		fmt.Println("无法整除：", i)
	}
}


//fallthrough关键字用于在switch语句中，将控制流程转移到下一个case语句中，并执行相应的语句块
func fallthrough_kw(){
	URL := "CSDN_URL"
	switch URL {
	case "CSDN_URL":
		fmt.Println(CSDN_URL)
		fallthrough
	case "WEIXIN_URL":
		fmt.Println(WEIXIN_URL)
	case "GIT_URL":
		fmt.Println(GIT_URL)
	default:
		fmt.Println("is not url patterns")
	}
}


// for关键字用于实现循环结构，可以用于遍历数组、切片、映射等数据结构，也可以用于实现无限循环。
func for_kw(){
	// 遍历字符串
	for _,v := range WEIXIN_URL{
		fmt.Printf("%c", v)
	}
	fmt.Println()

	for i := 0; i < 4; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	//数组
	arr := [5]int{11, 22, 33, 44, 55}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d\t", arr[i])
	}
	fmt.Println()

	//切片
	slice := []int{1,2,3,4,5}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d\t", slice[i])
	}
	fmt.Println()

	//映射
	/*
	9. map
	map关键字用于定义一个映射，映射是一种键值对的集合。每个键必须是唯一的，值可以是任意类型。
	*/
	// 定义的同时初始化
	maps := map[string]string{"csdn":CSDN_URL, "weixin":WEIXIN_URL, "git":GIT_URL}
	for k,v := range maps {
		fmt.Println(k, v)
	}
}


/*
10. make
make关键字用于创建一个切片、映射或通道。make函数返回一个已初始化的对象，可以直接使用。
*/

func map_kw(){
	//创建切片
	//创建映射
	maps := make(map[string]string)
	// 先定义后初始化
	maps["csdn"] = CSDN_URL
	maps["weixin"] = WEIXIN_URL
	maps["git"] = GIT_URL
	fmt.Println(len(maps))
	fmt.Println(maps["weixin"])
	fmt.Println(maps["csdn"])

	//创建通道
}


func func_kw(num1, num2 int, name string) int{
	fmt.Println("Hello,", name)
	return num1 + num2
}

func go_kw1(){	
	/*
	12. go
	go关键字用于启动一个新的并发线程，以便在后台执行一个函数。使用go关键字可以实现并发编程。
	*/
	go func(){
		fmt.Println(WEIXIN_URL)
	}()
}

func go_kw2(){	
	for i := 0; i < 3; i++ {
		go func(){
			fmt.Println(WEIXIN_URL)
		}()
	}
	runtime.Gosched()
}


/*
6. type
type关键字用于定义一个类型，类型可以是基本类型、结构体、接口、函数等。类型定义可以提高代码的可读性和可维护性。
8. interface
interface关键字用于定义一个接口，接口是一种抽象类型，只包含方法的声明，不包含方法的实现。接口可以用于实现多态性。
*/
type SteveRocket interface{
	Area() float64
	Perimeter() float64
}


func select_kw(){
	c1 := make(chan string)
	c2 := make(chan int)
	go func ()  {
		c1 <- WEIXIN_URL
	}()

	go func ()  {
		c2 <- 25
	}()

	select{
	case c:= <- c1:
		fmt.Println(c)
	case c:= <- c2:
		fmt.Println(c)
	}
}


func return_kw(num1, num2 int, url string)(int, error){
	fmt.Println("this is ", url)
	if num1<0 || num2<0 {
		return 0, errors.New("invalid argument")
	}
	return num1 + num2, nil
}


/*
7. struct
struct关键字用于定义一个结构体，结构体是一种自定义的复合类型。结构体可以包含多个字段，每个字段可以是不同的类型。
*/
func struct_kw(){
	type Person struct {
		Name string
		Age int
		Blog string
	}

	p:=Person{"Steverocket", 25, WEIXIN_URL}
	fmt.Println("My name is", p.Name , ". age is", p.Age, ". this is my blog", p.Blog)
}


func type_kw(){
	//定义结构体类型
	type Person struct {
		Name string
		Age int
		Blog string
	}

	//定义基本类型
	type MyInt int
	var age MyInt = 25
	fmt.Println(age)
}


func var_kw(){
	// 定义基本类型
	var age int = 25
	var blog string = WEIXIN_URL
	fmt.Println(age, blog)

	// 定义指针
	var ptr *int = &age
	fmt.Println(*ptr)
}


/*
13. defer
defer关键字用于在函数返回之前执行一些操作，例如关闭文件、释放资源等。defer语句按照后进先出的顺序执行。
*/
func defer_kw(){
	defer fmt.Println(WEIXIN_URL)
	for _,k := range CSDN_URL{
		fmt.Printf("%c", k)
	}
	fmt.Println("\nrun over...")
}


func goto_kw(){
	times := 0
	Here:
		fmt.Println("this is ", times , WEIXIN_URL)
		times += 1
		if times < 3 {
			goto Here
		}
		fmt.Println(CSDN_URL)
}




func main() {
	//fmt.Println(CSDN_URL)

	//break_kw()

	//case_kw()

	//chan_kw(26)

	//const_kw()

	//continue_kw()

	//fallthrough_kw()

	//for_kw()

	//fmt.Println(func_kw(11,22,"steverocket"))

	/*
	go_kw1()
	go_kw2()
	//go go_kw1()
	// 使用go关键字放在函数前，这样就定义了一个goroutine
	go func(){
		fmt.Println(CSDN_URL)
	}()
	time.Sleep(time.Second * 3)
	*/

	//goto_kw()

	//fmt.Println(WEIXIN_URL)

	//map_kw()

	/*
	result, err := return_kw(0, 13, WEIXIN_URL)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(return_kw(0, 13, WEIXIN_URL))
	fmt.Println(return_kw(11, 22, WEIXIN_URL))
	*/

	//select_kw()

	//struct_kw()

	//var_kw()

	defer_kw()


}