package main

var a string

func main() {
	a = "SteveRocket"
	print(a) // SteveRocket
	f1()
}
func f1() {
	a := "Cramer"
	print(a) // Cramer
	f2()
}
func f2() {
	print(a) // SteveRocket
}
