package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Create("SteveRocket.txt") // 如果文件已存在，会将文件清空
	fmt.Println(fp, err)                    // &{0xc000120780} <nil>
}
