package main

import "fmt"

func main() {
	var val interface{} = "公众号：CTO Plus"
	str, ok := val.(string)
	if ok {
		fmt.Println(str) // 公众号：CTO Plus
	} else {
		fmt.Println("Value is not a string")
	}
}
