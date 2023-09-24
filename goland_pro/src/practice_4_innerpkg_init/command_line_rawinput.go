package main

import (
	"bufio"
	"log"
	"os"
)

/*
	演示命令行输入
*/
func main() {
	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		log.Print("please input command content:")
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "stop" {
			running = false
		}
		log.Println("command contents =>>", command)
	}
}
