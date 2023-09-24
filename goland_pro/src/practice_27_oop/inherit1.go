package main

import "fmt"

type ClassA struct {
	num1, num2, num3 int
}

//CLassB继承ClassA
type ClassB struct {
	CLassA
	d string
}

func (classb ClassB) print() {
	fmt.Println(classb.d)
}

func main() {

}
