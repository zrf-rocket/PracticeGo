package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

/*
使用Go语言处理信号
Go信号通知通过在通道上发送os.Signal值来工作。创建一个通道来接收这些通知(还会在程序退出时通知我们)。
*/

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	//Ctrl+c 即可退出程序
	fmt.Println("exiting")
}
