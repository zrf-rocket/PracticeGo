package main

import "fmt"

// 变量名不能使用如下关键字
/*
break
default
func
interface
select
case
defer
go
map
struct
chan
else
goto
package
switch
const
fallthrough
if
range
type
continue
for
import
return
var
*/

// 在 Go 语言中，采用的是后置类型的声明方式
func variable() {
	// 变量声明
	var name string
	var age int

	// 变量赋值
	name = "SteveRocket"
	age = 22

	// 表示声明一个名为 email 的整数变量，并且附上初始值为 rocket_2014@126.com
	var email string = "rocket_2014@126.com"

	// 变量初始化 简短声明方式
	address := "广东省深圳市"
	// 声明 year 为 2023 的 64 位整数
	year := int64(2023)

	// 如果有多个变量同时声明，我们可以采用 var 加括号批量声明的方式
	var (
		num1, num2 int     // 0
		salary     float32 //0
	)

	fmt.Println(name, age, address, email, num1+num2, salary, year)
}
