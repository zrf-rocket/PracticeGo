package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

// 建立TCP链接来实现初步的HTTP协议，通过向网络主机发送HTTP  Head 请求，读取网络主机返回的信息
/*
go run "5网络编程 tcphttp.go" wwww.baidu.com:80
*/
func main() {
	var arg = os.Args
	if len(arg) != 2 {
		fmt.Println(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}

	service := arg[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
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

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
