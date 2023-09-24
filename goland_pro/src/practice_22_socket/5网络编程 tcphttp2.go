package main

import (
	"fmt"
	"io/ioutil"
	"net" //net 包中还包含了一系列的工具函数，合理地使用这些函数可以更好地保障程序的质量。
	"os"
)

func main() {
	arg := os.Args
	if len(arg) != 2 {
		fmt.Println(os.Stderr, "Usage:%s  host:port", os.Args[0])
		os.Exit(1)
	}

	service := arg[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service) //net.ResolveTCPAddr() ，用于解析地址和端口号
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr) // net.DialTCP() ，用于建立链接。
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
