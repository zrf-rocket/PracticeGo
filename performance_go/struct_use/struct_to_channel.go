package main

import "fmt"

func worker(ch chan struct{}){
	<-ch
	fmt.Println("微信公众号：CTO Plus")
	close(ch)
}

func main(){
	fmt.Println("Author:SteveRocket")
	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}
}
