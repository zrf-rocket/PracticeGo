package main

import "fmt"

func dictType() {
	var d = map[string]string{}
	d["dict"] = "this is Golang dict"
	fmt.Println(d) //map[dict:this is Golang dict]

	d["two"] = "this is other content"
	fmt.Println(d) //map[dict:this is Golang dict two:this is other content]

	d["author"] = "SteveRocket"
	fmt.Println(d) // map[author:SteveRocket dict:this is Golang dict two:this is other content]

	d["blog"] = "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"
	fmt.Println(d)      // map[author:SteveRocket blog:https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q dict:this is Golang dict two:this is other content]
	fmt.Println(len(d)) // 4

	d["author"] = "Cramer"
	fmt.Println(d)      // map[author:Cramer blog:https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q dict:this is Golang dict two:this is other content]
	fmt.Println(len(d)) // 4
}

func main() {

}
