package main

import "fmt"
import "time"
import "sync"

/*
channel
channel是Go语言在语言级别提供的goroutine间的通信方式。我们可以使用channel在两个或多个goroutine之间传递消息。channel是进程内的通信方式，
因此通过channel传递对象的过程和调用函数时的参数传递行为比较一致，比如也可以传递指针等。如果需要跨进程通信，我们建议用分布式系统的方法来解决，
比如使用Socket或者HTTP等通信协议。Go语言对于网络方面也有非常完善的支持。

channel是类型相关的。也就是说，一个channel只能传递一种类型的值，这个类型需要在声
明channel时指定。如果对Unix管道有所了解的话，就不难理解channel，可以将其认为是一种类
型安全的管道。


// 使用Go语言开发时，经常会遇到需要实现条件等待的场景，这也是channel可以发挥
// 作用的地方。对channel的熟练使用，才能真正理解和掌握Go语言并发编程。
	select
	缓冲机制
	超时机制
	channel的传递
	单向channel
	关闭channel

多核并行化
出让时间片
同步
	时间锁
	全局唯一性操作
*/

func main() {
	// Channel1()
	// Channel2()
	// Select()

	// BufferMechanism()
	// TimeoutMechanism()
	// PassTheChannel()
	// OnewayChannel()
	// CloseTheChannel()
	MulticoreParallelization()
	TransferTimeSlice()
	Synchronize()
	TimeLock()
	GlobalUniqueOperation()

}

// 用channel实现了类似锁的功能，进而保证了所有goroutine完成后主函数才返回。
func Channel1() {
	// 定义了一个包含10个channel的数组（名为 chs ），并把数组中的每个channel分配给10个不同的goroutine。
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	// 在所有的goroutine启动完成后，我们通过 <-ch 语句从10个channel中依次读取数据。在对应的channel写入数据前，这个操作也是阻塞的。
	for _, ch := range chs {
		<-ch
		fmt.Println(ch, *&ch)
	}
}

// 在每个goroutine的 Count() 函数完成后，我们通过 ch <- 1 语句向对应的channel中写入一个数据。在这个channel被读取前，这个操作是阻塞的。
func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting.......")
}

func Channel2() {
	// 声明一个传递类型为 int 的channel
	var ch chan int //与一般的变量声明不同的地方仅仅是在类型之前加了 chan 关键字
	// int指定这个channel所能传递的元素类型

	// 声明一个 map ，元素是 bool 型的channel
	var m map[string]chan bool

	// 定义一个channel也很简单，直接使用内置的函数 make() 即可
	ch2 := make(chan int) //这就声明并初始化了一个 int 型的名为 ch2 的channel
	_, _ = m, ch2

	// 在channel的用法中，最常见的包括写入和读出。
	// 将一个数据写入（发送）至channel
	ch <- 123

	// 向channel写入数据通常会导致程序阻塞，直到有其他goroutine从这个channel中读取数据。
	// 从channel中读取数据的语法是
	value := <-ch
	fmt.Println(value)
	// 如果channel之前没有写入数据，那么从channel中读取数据也会导致程序阻塞，直到channel中被写入数据为止。

}

// select
// 早在Unix时代， select 机制就已经被引入。通过调用 select() 函数来监控一系列的文件句柄，一旦其中一个文件句柄发生了IO动作，该 select() 调用就会被返回。
// 后来该机制也被用于实现高并发的Socket服务器程序。Go语言直接在语言级别支持 select 关键字，用于处理异步IO问题。
func Select() {
	// select 的用法与 switch 语言非常类似，由 select 开始一个新的选择块，每个选择条件由case 语句来描述。与 switch 语句可以选择任何可使用相等比较的条件相比，
	// select 有比较多的限制，其中最大的一条限制就是每个 case 语句里必须是一个IO操作

	var chan1 chan int
	// var chan2 int  //error   invalid operation: chan2 <- 1 (send to non-chan type int)
	var chan2 chan int
	select {
	case <-chan1: // 试图从 chan1 读取一个数据并直接忽略读到的数据
		// 如果chan1成功读到数据，则进行该case处理语句
	case chan2 <- 1: //试图向 chan2 中写入一个整型数1，如果这两者都没有成功，则到达 default 语句
		//如果成功向chan2写入数据，则进行该case处理语句
	default:
		//如果上面都没有成功，则进入default处理流程
	}

	//死循环   随机向 ch 中写入一个0或者1的过程
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("value received:", i)
	}
}

// select 不像 switch ，后面并不带判断条件，而是直接去查看 case 语句。每个case 语句都必须是一个面向channel的操作。

// 缓冲机制
func BufferMechanism() {
	// 之前我们示范创建的都是不带缓冲的channel，这种做法对于传递单个数据的场景可以接受，
	// 但对于需要持续传输大量数据的场景就有些不合适了。接下来我们介绍如何给channel带上缓冲，
	// 从而达到消息队列的效果。

	// 创建一个带缓冲的channel
	c := make(chan int, 1024) //在调用 make() 时将缓冲区大小作为第二个参数传入即可创建了一个大小
	// 为1024的 int 类型 channel ，即使没有读取方，写入方也可以一直往channel里写入，在缓冲区被
	// 填完之前都不会阻塞。

	// 从带缓冲的channel中读取数据可以使用与常规非缓冲channel完全一致的方法，但我们也可
	// 以使用 range 关键来实现更为简便的循环读取
	for i := range c {
		fmt.Println("received:", i)
	}
}

// 超时机制
func TimeoutMechanism() {
	// 在并发编程的通信过程中，最需要处理的就是超时问题，即向channel写数据时发现channel
	// 已满，或者从channel试图读取数据时发现channel为空。如果不正确处理这些情况，很可能会导
	// 致整个goroutine锁死。

	// 虽然goroutine是Go语言引入的新概念，但通信锁死问题已经存在很长时间，在之前的C/C++
	// 开发中也存在。操作系统在提供此类系统级通信函数时也会考虑入超时场景，因此这些方法通常
	// 都会带一个独立的超时参数。超过设定的时间时，仍然没有处理完任务，则该方法会立即终止并
	// 返回对应的超时信息。超时机制本身虽然也会带来一些问题，比如在运行比较快的机器或者高速
	// 的网络上运行正常的程序，到了慢速的机器或者网络上运行就会出问题，从而出现结果不一致的
	// 现象，但从根本上来说，解决死锁问题的价值要远大于所带来的问题。

	// 使用channel时需要小心，比如对于以下这个用法：
	// i := <-ch
	// 不出问题的话一切都正常运行。但如果出现了一个错误情况，即永远都没有人往 ch 里写数据，那
	// 么上述这个读取动作也将永远无法从 ch 中读取到数据，导致的结果就是整个goroutine永远阻塞并
	// 没有挽回的机会。如果channel只是被同一个开发者使用，那样出问题的可能性还低一些。但如果
	// 一旦对外公开，就必须考虑到最差的情况并对程序进行保护。

	// Go语言没有提供直接的超时处理机制，但我们可以利用 select 机制。虽然 select 机制不是
	// 专为超时而设计的，却能很方便地解决超时问题。因为 select 的特点是只要其中一个 case 已经
	// 完成，程序就会继续往下执行，而不会考虑其他 case 的情况。

	// 为channel实现超时机制
	// 首先，我们实现并执行一个匿名的超时等待函数
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) // 等待1秒钟
		timeout <- true
	}()

	var ch chan int
	// 然后我们把timeout这个channel利用起来
	select {
	case <-ch:
		// 从ch中读取到数据
	case <-timeout:
		// 一直没有从ch中读取到数据，但从timeout中读取到了数据
	}
	// 这样使用 select 机制可以避免永久等待的问题，因为程序会在 timeout 中获取到一个数据
	// 后继续执行，无论对 ch 的读取是否还处于等待状态，从而达成1秒超时的效果。
	// 这种写法看起来是一个小技巧，但却是在Go语言开发中避免channel通信超时的最有效方法。
	// 在实际的开发过程中，这种写法也需要被合理利用起来，从而有效地提高代码质量。
}

// channel的传递
func PassTheChannel() {
	// 在Go语言中channel本身也是一个原生类型，与 map 之类的类型地位一样，因此channel本身在定义后也可以通过channel来传递。

	// 可以使用这个特性来实现*nix上非常常见的管道（pipe）特性。管道也是使用非常广泛的一种设计模式，比如在处理数据时，
	// 我们可以采用管道设计，这样可以比较容易以插件的方式增加数据的处理流程。

	// 利用channel可被传递的特性来实现我们的管道 假设在管道中传递的数据只是一个整型数，在实际的应用场景中这通常会是一个数据块。

}

//基本的数据结构
type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

// 只要定义一系列 PipeData 的数据结构并一起传递给这个函数，就可以达到流式处理数据的目的
func handle(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

// 单向channel    如何控制channel只接受写或者只允许读取
func OnewayChannel() {
	// 单向channel只能用于发送或者接收数据。channel本身必然是同时支持读写的，
	// 否则根本没法用。假如一个channel真的只能读，那么肯定只会是空的，因为你没机会往里面写数
	// 据。同理，如果一个channel只允许写，即使写进去了，也没有丝毫意义，因为没有机会读取里面
	// 的数据。所谓的单向channel概念，其实只是对channel的一种使用限制。

	// 在将一个channel变量传递到一个函数时，可以通过将其指定为单向channel变量，从
	// 而限制该函数中可以对此channel的操作，比如只能往这个channel写，或者只能从这个
	// channel读

	// 单向channel变量的声明
	var ch1 chan int       // ch1是一个正常的channel，不是单向的
	var ch2 chan<- float64 // ch2是单向channel，只用于写float64数据
	var ch3 <-chan int     // ch3是单向channel，只用于读取int数据

	// 那么单向channel如何初始化呢？
	// channel是一个原生类型，因此不仅
	// 支持被传递，还支持类型转换。只有在介绍了单向channel的概念后，读者才会明白类型转换对于
	// channel的意义：就是在单向channel和双向channel之间进行转换。
	ch4 := make(chan int)
	ch5 := <-chan int(ch4) // ch5就是一个单向的读取channel
	ch6 := chan<- int(ch4) // ch6 是一个单向的写入channel
	// 基于 ch4 ，我们通过类型转换初始化了两个单向channel：单向读的 ch5 和单向写的 ch6
	// 为什么要做这样的限制呢？从设计的角度考虑，所有的代码应该都遵循“最小权限原则”，
	// 从而避免没必要地使用泛滥问题，进而导致程序失控。

	// 写过C++程序的读者肯定就会联想起 const
	// 指针的用法。非 const 指针具备 const 指针的所有功能，将一个指针设定为 const 就是明确告诉
	// 函数实现者不要试图对该指针进行修改。单向channel也是起到这样的一种契约作用。

	_, _, _, _, _ = ch1, ch2, ch3, ch5, ch6
	close(ch1)
	close(ch2)
	// close(ch3)  // invalid operation: close(ch3) (cannot close receive-only channel)
	// close(ch5)  //invalid operation: close(ch5) (non-chan type int)
	close(ch6)
}

// 单向channel的用法
func Parse(ch <-chan int) {
	for value := range ch {
		fmt.Println("Parsing value", value)
	}
}

// 除非这个函数的实现者无耻地使用了类型转换，否则这个函数就不会因为各种原因而对 ch
// 进行写，避免在 ch 中出现非期望的数据，从而很好地实践最小权限原则。

// 关闭channel
func CloseTheChannel() {
	// 关闭channel非常简单，直接使用Go语言内置的 close() 函数即可
	// 如何判断一个channel是否已经被关闭？我们可以在读取的时候使用多重返回值的方式
	// x, ok := <-ch
	// 这个用法与 map 中的按键获取 value 的过程比较类似，只需要看第二个 bool 返回值即可，如
	// 果返回值是 false 则表示 ch 已经被关闭。
}

// 多核并行化
type Vector []float64

const NCPU = 16 // 假设总共有16核
func MulticoreParallelization() {
	// 需要了解CPU核心的数量，并针对性地分解计算任务到多个goroutine中去并行运行。

	// 并行的计算任务：计算N个整型数的总和。我们可以将所有整
	// 型数分成M份，M即CPU的个数。让每个CPU开始计算分给它的那份计算任务，最后将每个CPU
	// 的计算结果再做一次累加，这样就可以得到所有N个整型数的总和
	/*
	   虽然我们确实创建了多个goroutine，并且从运行状态看这些goroutine也都在并行运行，但实际上所有这些goroutine都运行在同一个CPU核心上，
	   在一个goroutine得到时间片执行的时候，其他goroutine都会处于等待状态。从这一点可以看出，虽然goroutine简化了我们写并行代码的过程，
	   但实际上整体运行效率并不真正高于单线程程序。

	   在Go语言升级到默认支持多CPU的某个版本之前，我们可以先通过设置环境变量GOMAXPROCS 的值来控制使用多少个CPU核心。
	   具体操作方法是通过直接设置环境变量GOMAXPROCS 的值，或者在代码中启动goroutine之前先调用以下这个语句以设置使用16个CPU核心
	   runtime.GOMAXPROCS(16)

	   到底应该设置多少个CPU核心呢，其实 runtime 包中还提供了另外一个函数 NumCPU() 来获取核心数。可以看到，Go语言其实已经感知到所有的环境信息，
	   下一版本中完全可以利用这些信息将goroutine调度到所有CPU核心上，从而最大化地利用服务器的多核计算能力。抛弃GOMAXPROCS 只是个时间问题。
	*/
}

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		//v[i] += u.Op(v[i])
	}
	c <- 1 // 发信号告诉任务管理者我已经计算完成了
}

func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU) // 用于接收每个CPU的任务完成信号
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	// 等待所有CPU的任务完成
	for i := 0; i < NCPU; i++ {
		<-c // 获取到一个数据，表示一个CPU计算完成了
	}
	// 到这里表示所有计算已经结束
}

// 出让时间片
func TransferTimeSlice() {
	// 可以在每个goroutine中控制何时主动出让时间片给其他goroutine，这可以使用 runtime包中的 Gosched() 函数实现。
	// 实际上，如果要比较精细地控制goroutine的行为，就必须比较深入地了解Go语言开发包中runtime 包所提供的具体功能

}

// 同步   同步锁
func Synchronize() {
	// 倡导用通信来共享数据，而不是通过共享数据来进行通信，但考虑
	// 到即使成功地用channel来作为通信手段，还是避免不了多个goroutine之间共享数据的问题，Go
	// 语言的设计者虽然对channel有极高的期望，但也提供了妥善的资源锁方案。

	// Go语言包中的 sync 包提供了两种锁类型： sync.Mutex 和 sync.RWMutex 。 Mutex 是最简单
	// 的一种锁类型，同时也比较暴力，当一个goroutine获得了 Mutex 后，其他goroutine就只能乖乖等
	// 到这个goroutine释放该 Mutex 。 RWMutex 相对友好些，是经典的单写多读模型。在读锁占用的情
	// 况下，会阻止写，但不阻止读，也就是多个goroutine可同时获取读锁（调用 RLock() 方法；而写
	// 锁（调用 Lock() 方法）会阻止任何其他goroutine（无论读和写）进来，整个锁相当于由该goroutine
	// 独占。从 RWMutex 的实现看， RWMutex 类型其实组合了 Mutex
	type RWMutex struct {
		// w Mutex
		writerSem   uint32
		readerSem   uint32
		readerCount int32
		readerWait  int32
	}

	// 对于这两种锁类型，任何一个 Lock() 或 RLock() 均需要保证对应有 Unlock() 或 RUnlock()
	// 调用与之对应，否则可能导致等待该锁的所有goroutine处于饥饿状态，甚至可能导致死锁。锁的
	// 典型使用模式如下
	// var i sync.Mutex
	// func foo(){
	// 	i.Lock()
	// 	defer i.Unlock()
	// 	//...
	// }
}

// 时间锁
func TimeLock() {

}

var a string
var once sync.Once

func setup() {
	a = "this is golang sync.Once"
}

// 如果这段代码没有引入 Once ， setup() 将会被每一个goroutine先调用一次，这至少对于这个
// 例子是多余的。在现实中，我们也经常会遇到这样的情况。Go语言标准库为我们引入了 Once 类
// 型以解决这个问题。
func doprint() {
	once.Do(setup)
	// 	 once 的 Do() 方法可以保证在全局范围内只调用指定的函数一次（这里指
	// setup() 函数），而且所有其他goroutine在调用到此语句时，将会先被阻塞，直至全局唯一的
	// once.Do() 调用结束后才继续。
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

// 全局唯一性操作
func GlobalUniqueOperation() {
	// 对于从全局的角度只需要运行一次的代码，比如全局初始化操作，Go语言提供了一个 Once类型来保证全局的唯一性操作
	twoprint()

	doprint2()

}

// 如果没有 once.Do() ，我们很可能只能添加一个全局的 bool 变量，在函数 setup() 的最后
// 一行将该 bool 变量设置为 true 。在对 setup() 的所有调用之前，需要先判断该 bool 变量是否已
// 经被设置为 true ，如果该值仍然是 false ，则调用一次 setup() ，否则应跳过该语句。

var done bool = false

func setup2() {
	a = "\nthis is other setup2 function"
	done = true
}

// 这段代码初看起来比较合理，但是细看还是会有问题，因为 setup() 并不是一个原子性操作，
// 这种写法可能导致 setup() 函数被多次调用，从而无法达到全局只执行一次的目标。这个问题的
// 复杂性也更加体现了 Once 类型的价值。
func doprint2() {
	if !done {
		setup2()
	}
	print(a)
}

// 为了更好地控制并行中的原子性操作， sync 包中还包含一个 atomic 子包，它提供了对于一
// 些基础数据类型的原子操作函数，比如下面这个函数
// func CompareAndSwapUint64(val *uint64, old , new uint64)(swapped bool)
// 供了比较和交换两个 uint64 类型数据的操作。这让开发者无需再为这样的操作专门添加Lock 操作
