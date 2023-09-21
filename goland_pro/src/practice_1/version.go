package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前go版本
	fmt.Printf("%s", runtime.Version()) // go1.20.4
}
