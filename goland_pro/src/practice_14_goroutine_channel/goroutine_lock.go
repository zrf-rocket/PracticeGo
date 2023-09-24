package main

import "fmt"
import "sync"
import "runtime"

/*
并发通信
关键字 go 的引入使得在Go语言中并发编程变得简单而优雅，但我们同时也应该意识到并发编程的原生复杂性，并时刻对并发中容易出现的问题保持警惕。
并发编程的难度在于协调，而协调就要通过交流。并发单元间的通信是最大的问题。

两种最常见的并发通信模型：共享数据和消息
	共享数据是指多个并发单元分别保持对同一个数据的引用，实现对该数据的共享。被共享的
数据可能有多种形式，比如内存数据块、磁盘文件、网络数据等。在实际工程应用中最常见的无
疑是内存了，也就是常说的共享内存。

	Go语言提供的是另一种通信模型，即以消息机制而非共享内存作为通信方式。
消息机制认为每个并发单元是自包含的、独立的个体，并且都有自己的变量，但在不同并发
单元间这些变量不共享。每个并发单元的输入和输出只有一种，那就是消息。这有点类似于进程
的概念，每个进程不会被其他进程打扰，它只做好自己的工作就可以了。不同进程间靠消息来通
信，它们不会共享内存。
	Go语言提供的消息通信机制被称为channel
	不要通过共享内存来通信，而应该通过通信来共享内存。
*/

var counter int = 0

// 在10个goroutine中共享了变量 counter 。每个goroutine执行完成后，
// 将 counter 的值加1。因为10个goroutine是并发执行的，所以我们还引入了锁，也就是代码中的
// lock 变量。每次对 n 的操作，都要先将锁锁住，操作完成后，再将锁打开。
func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func main() {
	// Goroutine1()

	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	// go Goroutine1(lock)

	for {
		lock.Lock()
		// 使用 for循环来不断检查 counter 的值（同样需要加锁）。当其值达到10时，说明所有goroutine都执行完毕了，
		// 这时主函数返回，程序退出。
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}

	// go Goroutine1(lock)
}

func Goroutine1(lock *sync.Mutex) {
	lock.Lock()
	for i := 0; i < 5; i++ {
		go Add(1111, 3333)
	}
	lock.Unlock()
}

func Add(num1, num2 int) {
	fmt.Println(num1 + num2)
}
