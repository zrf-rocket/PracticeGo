package main

import "fmt"

func main() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for index, value := range pow {
		fmt.Printf("2**%d = %d\n", index, value) //此处的%d匹配支持C风格 Printf
	}

	mapValue := map[string]string{
		"name": "SteveRocket",
		"blog": "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q",
	}
	for key, value := range mapValue {
		fmt.Println(key, value)
	}
}
