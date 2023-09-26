package main

type MyStruct struct {
	Name string
	Age  int
}

type MyStruct02 struct {
	Name, Blog string
	Age        int
}

func main() {
	var myInfo MyStruct02
	myInfo = MyStruct02{Name: "SteveRocket", Age: 28, Blog: "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"}

	var myInfo02 MyStruct02
	myInfo02 = MyStruct02{Name: "SteveRocket", Age: 28, Blog: "https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"}

}
