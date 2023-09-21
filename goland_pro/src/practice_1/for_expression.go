package main

import (
	"fmt"
	"time"
)

func for_expreesion1() {
	index := 0
	// 死循环
	for {
		fmt.Println("run loop.")
		index += 1
		for index > 5 {
			fmt.Println("large number")
			break // 没有break 则此处死循环
		}
		time.Sleep(time.Second)
	}
}

func for_expreesion2() {
	map := {"name":"SteveRocket"}
	for key, value := range map{
		fmt.Println(key, value)
	}
}

func main() {
	//for_expreesion1()
	for_expreesion2()
}
