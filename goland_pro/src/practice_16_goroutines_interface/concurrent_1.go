package main

/*
并发模式之内核
这种并发模式的内核只需要协程和通道就够了。协程负责执行代码，通道负责在协程之间传递事件。
要想编写一个良好的并发程序，我们不得不了解线程，锁，semaphore，barrier甚至CPU更新高速缓存的方式，而且他们个个都有怪脾气，处处是陷阱。

协程是轻量级的线程。在过程式编程中，当调用一个过程的时候，需要等待其执行完才返回。而调用一个协程的时候，不需要等待其执行完，会立即返回。
协程十分轻量，Go语言可以在一个进程中执行有数以十万计的协程，依旧保持高性能。而对于普通的平台，一个进程有数千个线程，其CPU会忙于上下文切换，
性能急剧下降。随意创建线程可不是一个好主意，但是我们可以大量使用的协程。




*/
func main() {

}
