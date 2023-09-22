package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Hello", "-", "SteveRocket") // Hello-SteveRocket

	name := "SteveRocket"
	age := 28
	// 类似C语言  %s是一个占位符 使用name这个变量的值替换%s占位符
	fmt.Printf("My name is %s and i am %d year old.", name, age) //My name is SteveRocket and i am 28 year old.

	fmt.Println("\n33 + 44 =", 33+44) // 33 + 44 = 77

	infos := fmt.Sprintln("Hello", name, age)
	fmt.Println(infos) // Hello SteveRocket 28

	file, _ := os.Create("steverocket.txt")
	fmt.Fprintln(file, "Hello", name, age, "微信公众号：CTO Plus")

	file2, _ := os.Create("steverocket.txt")
	fmt.Fprintf(file2, "My name is %s and I am %d years old.", name, age)

	info := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(info) // My name is SteveRocket and I am 28 years old.

	err := fmt.Errorf("An error occurred: %s, age is %d", "file not found", age)
	fmt.Println(err) // An error occurred: file not found, age is 28

	println("Hello", "SteveRocket", 28, 123.456, true) // Hello SteveRocket 28 +1.234560e+002 true
	print("Hello", "SteveRocket", 28, 123.456, true)   // HelloSteveRocket28+1.234560e+002true
	println("args1", "args2", "args3")                 // args1 args2 args3
}
