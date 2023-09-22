package main

import (
	"fmt"
	"math"
)

//内置模块math

func mathMax() {

	var maxValue = math.MaxFloat32
	var maxValue2 = math.MaxFloat64

	fmt.Println(maxValue, maxValue2) // 3.4028234663852886e+38 1.7976931348623157e+308

	maxIntValue := math.MaxInt
	maxInt8Value := math.MaxInt8
	maxInt16Value := math.MaxInt16
	maxInt32Value := math.MaxInt32
	maxInt64Value := math.MaxInt64
	maxUint8Value := math.MaxUint8
	maxUint16Value := math.MaxUint16
	maxUint32Value := math.MaxUint32
	fmt.Println(maxIntValue, maxInt8Value, maxInt16Value, maxInt32Value, maxInt64Value, maxUint8Value, maxUint16Value, maxUint32Value)
	// 9223372036854775807 127 32767 2147483647 9223372036854775807 255 65535 4294967295

	//maxUint64Value := math.MaxUint64 // 1<<64-1
	//cannot use math.MaxUint64 (untyped int constant 18446744073709551615) as int value in assignment (overflows)
	//fmt.Println(maxUint64Value)
}
func mathFunc() {
	//absValue  := math.Abs(-123)
	//maxValue := math.Max([11,22,33, 44, 55, 66])
	//math.Acos()
}
func main() {
	mathMax()
	//mathFunc()
}
