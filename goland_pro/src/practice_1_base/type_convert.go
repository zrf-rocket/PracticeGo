package main

import (
	"fmt"
	"math"
	"strconv"
)

func varConvert() {
	var num1 int = 28
	var num2 float64 = 23.456
	// 不同类型的变量不能拿来运算
	//sum := num2 + num1  // invalid operation: num2 + num1 (mismatched types float64 and int)
	sum := num1 + int(num2)
	sum2 := float64(num1) + num2
	fmt.Println(sum, sum2) // 51 51.456
	// 不同类型的变量不能互相比较
	//fmt.Println(sum == sum2) // invalid operation: sum == sum2 (mismatched types int and float64)
}

func strongConvert() {
	var age int = 28
	// 强制类型转换（实际应用场景上也不是一个靠谱的方法）
	var height float64 = float64(age)
	fmt.Println(age, height) // 28 28

	var num int = 10
	var byteValue byte = byte(num)                       //int转为byte
	newNum := int(byteValue)                             //byte转为int
	fmt.Printf("%T %T %T", byte(num), byteValue, newNum) // uint8 uint8 int
}

func varConvert02() {
	//高精度向低精度转换，数字很小时这种转换没问题
	var uValue uint64 = 1
	i8Value := int8(uValue)
	fmt.Printf("i8Value=%d\n", i8Value) // i8Value=1

	//最高位的1变成了符号位
	uValue = uint64(math.MaxUint64)
	i64Value := int64(uValue)
	fmt.Printf("i64Value=%d\n", i64Value) // i64Value=-1

	//位数丢失
	ui32Value := uint32(uValue)
	fmt.Printf("ui32Value=%d\n", ui32Value) // ui32Value=4294967295

	//单个字符可以转为int
	var num int = int('a')
	//输出字母a对应的ASCII码
	fmt.Printf("num=%d\n", num) // num=97

	//bool和int不能相互转换

	//byte和int可以互相转换
	var itob byte = byte(num)
	btoi := int(itob)
	fmt.Printf("itob=%d  %T\n", itob, itob) // itob=97  uint8
	fmt.Printf("btoi=%d  %T\n", btoi, btoi) // btoi=97  int

	//float和int可以互相转换，小数位会丢失
	var float32Value float32 = float32(12345678)
	float32ToInt := int(float32Value)
	fmt.Printf("float32Value=%d  %T\n", float32Value, float32Value) // float32Value=%!d(float32=1.2345678e+07)  float32
	fmt.Printf("float32ToInt=%d  %T\n", float32ToInt, float32ToInt) // float32ToInt=12345678  int
}

func strconvFunc() {
	//string和其他数据类型互转
	var err error
	var i int = 8
	var i64 int64 = int64(i)
	//int转string
	var s string = strconv.Itoa(i) //内部调用FormatInt
	s = strconv.FormatInt(i64, 10)
	//string转int
	i, err = strconv.Atoi(s)
	//string转int64
	i64, err = strconv.ParseInt(s, 10, 64)

	//float转string
	var f float64 = 8.123456789
	s = strconv.FormatFloat(f, 'f', 2, 64) //保留2位小数
	fmt.Println(s)
	//string转float
	f, err = strconv.ParseFloat(s, 64)

	//string<-->[]byte
	var arr []byte = []byte(s)
	s = string(arr)

	//string<-->[]rune
	var brr []rune = []rune(s)
	s = string(brr)

	fmt.Printf("err %v\n", err)
}
func main() {
	//varConvert()
	//strongConvert()
	varConvert02()
	//strconvFunc()
}
