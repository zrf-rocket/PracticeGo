package main

import "fmt"
import "math"

type ClassA struct {
	num1, num2, num3 int
}

func (classa *ClassA) sum() int { //传引用（指针）
	return classa.num1 + classa.num2
}

type MyFloat float64 // 可对任何类型增加方法

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	// fmt.Println(ClassA(11,22).sum())  //调用错误

	fmt.Println(MyFloat(-math.Sqrt2).Abs()) //1.4142135623730951
}
