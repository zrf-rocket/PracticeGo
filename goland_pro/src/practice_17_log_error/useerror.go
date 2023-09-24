package main

import (
	"errors"
	"fmt"
)

//error类型是一个接口类型
func main() {
	SqrtTest()

}

//可以在编码中通过实现 error 接口类型来生成错误信息
//函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math:square root of negative number")
	}
	return f, errors.New("success......")
}

func SqrtTest() {
	fmt.Println(Sqrt(1.1))
	fmt.Println(Sqrt(-1))

	//接收方法返回的多个值
	res, status := Sqrt(1.2)
	fmt.Println(res)
	fmt.Println(status)

	_, err := Sqrt(-1)
	if err != nil {
		fmt.Println("输出结果：", err)
	}
}

//定义一个结构
type DivideError struct {
	divide  int
	divider int
}

//实现 error 接口
func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero
	dividee: %d
	divider: 0
	`
	return fmt.Sprintf(strFormat, de.divide)
}
