package main

type error interface {
	Error() string
}

//可以在编码中通过实现 error 接口类型来生成错误信息
//函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息
func Sqrt(f float64) (float64, error) {

}

func main() {

}
