package struct_use

import "fmt"

type SteveRocket struct {}

func (rocket SteveRocket) Open() {
	fmt.Println("open https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q")
}

func (rocket SteveRocket) Close(){
	fmt.Println("close https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q")
}
