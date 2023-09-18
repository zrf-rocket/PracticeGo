// 命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	var cnts, sep string
	for i := 1; i < len(os.Args); i++ {
		cnts += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(cnts)
}
