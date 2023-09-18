package main

import (
	"fmt"
	"os"
)

// 输出命令行参数

func main() {
	cnts, sep := "", ""
	for _, arg := range os.Args[1:] {
		cnts += sep + arg
		sep = " "
	}
	fmt.Println(cnts)
}
