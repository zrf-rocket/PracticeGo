package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

// 使用ICMP协议向在线的主机发送一个问候，并等待主机返回
/*
go build icmptest.go
./icmptest www.baidu.com
*/
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "host")
		os.Exit(0)
	}

	service := os.Args[1]
	fmt.Println(service) //输出程序的第一个参数

	conn, err := net.Dial("ip6:icmp", service)
	checkError(err)

	var msg [512]byte
	msg[0] = 8  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum
	msg[3] = 0  // checksum
	msg[4] = 0  // identifier[0]
	msg[5] = 13 //identifier[1]
	msg[6] = 0  // sequence[0]
	msg[7] = 37 // sequence[1]
	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[0:len])
	checkError(err)

	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("Got response")

	if msg[5] == 13 {
		fmt.Println("Identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}
	os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	sum := 0
	// 先假设为偶数
	for i := 1; i < len(msg)-1; i += 2 {
		sum += int(msg[i])*256 + int(msg[i+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
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
