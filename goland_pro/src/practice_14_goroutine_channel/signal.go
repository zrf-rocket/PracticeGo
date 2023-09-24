package main

import "fmt"
import "os"

/*
希望Go程序能够智能地处理Unix信号。 例如，可能希望服务器在接收到SIGTERM时正常关闭，或者在收到SIGINT时使用命令行工具停止处理输入。
*/
func main() {
	defer fmt.Println("!defer") //当使用os.Exit时，defers不会运行。
	os.Exit(3)                  //使用os.Exit立即退出并返回给定状态。
}
