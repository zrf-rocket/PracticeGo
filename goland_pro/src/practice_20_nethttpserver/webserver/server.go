package main

import (
	"fmt"
	"net/http"
	"strings"
)

/*
获取请求头和请求参数
*/
func param(w http.ResponseWriter, reuest *http.Request) {
	fmt.Fprintln(w, "第一个")
	headers := reuest.Header
	fmt.Fprintln(w, "header中的全部数据", headers)
	var arr []string = headers["Accept-Encoding"]
	split := strings.Split(arr[0], ",")
	for _, value := range split {
		fmt.Fprintln(w, strings.TrimSpace(value))
	}
	reuest.ParseForm()
	fmt.Fprintln(w, reuest.Form)
	/*
		按照请求参数名获取参数值
		根据源码,FormValue(key)=req.Form[key]
	*/
	name := reuest.FormValue("name")
	age := reuest.FormValue("age")
	fmt.Fprintln(w, name, age)
}

func main() {
	http.HandleFunc("/", param)
	// 访问方式：http://localhost:5656/?access_token=253453456
	http.ListenAndServe("0.0.0.0:5656", nil)
}
