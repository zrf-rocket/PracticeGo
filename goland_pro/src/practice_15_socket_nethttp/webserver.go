package main

import (
	"io" //使用 io 包而不是 fmt 包来输出字符串，这样源文件编译成可执行文件后，体积要小很多，运行起来也更省资源。
	"log"
	// "fmt"
	"net/http" //主要用于提供Web服务，响应并处理客户端（浏览器）的HTTP请求。
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is GoLang first Web application")
}

/*
helloHandler() 方法是 http.HandlerFunc 类型的实例，并传入http.ResponseWriter 和 http.Request 作为其必要的两个参数。
http.ResponseWriter 类型的对象用于包装处理HTTP服务端的响应信息。

将字符串 "This is GoLang first Web application" 写入类型为http.ResponseWriter 的 w 实例中，即可将该字符串数据发送到HTTP客户端。
第二个参数r *http.Request 表示的是此次HTTP请求的一个数据结构体，即代表一个客户端，不过该示例中我们尚未用到它。
*/
func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is GoLang index page....")
}

func main() {
	// http.HandleFunc() 方法接受两个参数，第一个参数是HTTP请求的目标路径 "/hello" ，该参数值可以是字符串，也可以是字符串形式的正则表达式，第二个参数指定
	// 具体的回调方法，比如 helloHandler 。当我们的程序运行起来后，访问http://localhost:8080/hello ，程序就会去调用 helloHandler() 方法中的业务逻辑程序。
	http.HandleFunc("/hello", helloHandler) //该方法用于分发请求，即针对某一路径请求将其映射到指定的业务逻辑处理方法中。
	http.HandleFunc("/", indexHandler)      //可以将其形象地理解为提供类似URL路由或者URL映射之类的功能。

	err := http.ListenAndServe(":8080", nil)
	// http.ListenAndServe() ，该方法用于在示例中监
	// 听 8080 端口，接受并调用内部程序来处理连接到此端口的请求。如果端口监听失败，会调用
	// log.Fatal() 方法输出异常出错信息。
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
