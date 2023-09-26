package main

import "os"

func main() {
	buf := make([]byte, 1024)
	// 下划线意思是忽略这个变量
	//f, err := os.Open("operator.go")
	//如果此时不需要知道返回的错误值, 则可以忽略掉err
	f, _ := os.Open("operator.go") // os.Open，返回值为*os.File，error
	defer f.Close()
	for {
		// 这种情况的下划线是起到一个占位的作用，方法返回两个结果，而只想要一个结果。
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
