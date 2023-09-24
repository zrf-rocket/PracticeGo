package main

// package singleton
import "fmt"

/*
单例模式(Singleton)：表示一个类只会生成唯一的一个对象。
A．这些类只能有一个实例；
B．这些能够自动实例化；
C．这个类对整个系统可见，即必须向整个系统提供这个实例。
*/

var _instance *object

type object struct {
	name string
}

func Instance() *object {
	if _instance == nil {
		_instance = new(object)
	}
	return _instance
}

func (p *object) SetName(name string) {
	p.name = name
}

func (p *object) Say() {
	fmt.Println(p.name)
}

func main() {
	// Say("rocket")
	// object.name = "rocket"
	// Instance(object)
}
