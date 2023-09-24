package main

import "errors"

// import "fmt"
// import "net/rpc"
// import "encoding/gob"

/*
RPC编程
	Go中的rpc支持与处理
	Gob
	设计优雅的RPC接口


RPC（Remote Procedure Call，远程过程调用）是一种通过网络从远程计算机程序上请求服
务， 而不需要了解底层网络细节的应用程序通信协议。 RPC协议构建于TCP或UDP， 或者是 HTTP
之上，允许开发者直接调用另一台计算机上的程序，而开发者无需额外地为这个调用过程编写网
络通信相关代码，使得开发包括网络分布式程序在内的应用程序更加容易。

RPC 采用客户端—服务器（Client/Server）的工作模式。请求程序就是一个客户端（Client） ，
而服务提供程序就是一个服务器（Server） 。当执行一个远程过程调用时，客户端程序首先发送一
个带有参数的调用信息到服务端，然后等待服务端响应。在服务端，服务进程保持睡眠状态直到
客户端的调用信息到达为止。当一个调用信息到达时，服务端获得进程参数，计算出结果，并向
客户端发送应答信息，然后等待下一个调用。最后，客户端接收来自服务端的应答信息，获得进
程结果，然后调用执行并继续进行。

在Go中，标准库提供的 net/rpc 包实现了 RPC 协议需要的相关细节，开发者可以很方便地
使用该包编写 RPC 的服务端和客户端程序， 这使得用 Go 语言开发的多个进程之间的通信变得非
常简单。

net/rpc 包允许 RPC 客户端程序通过网络或是其他 I/O 连接调用一个远端对象的公开方法
（必须是大写字母开头、可外部调用的） 。在 RPC 服务端，可将一个对象注册为可访问的服务，
之后该对象的公开方法就能够以远程的方式提供访问。一个 RPC 服务端可以注册多个不同类型
的对象，但不允许注册同一类型的多个对象。

一个对象中只有满足如下这些条件的方法，才能被 RPC 服务端设置为可供远程访问：
	必须是在对象外部可公开调用的方法（首字母大写）；
	必须有两个参数，且参数的类型都必须是包外部可以访问的类型或者是Go内建支持的类型；
	第二个参数必须是一个指针；
	方法必须返回一个 error 类型的值。

以上4个条件，可以简单地用如下一行代码表示
func (t *T) MethodName(argType T1, replyType *T2)error
类型 T 、 T1 和  T2 默认会使用 Go 内置的 encoding/gob 包进行编码解码
该方法（ MethodName ）的第一个参数表示由 RPC 客户端传入的参数，第二个参数表示要返
回给RPC客户端的结果，该方法最后返回一个  error 类型的值

RPC 服务端可以通过调用  rpc.ServeConn 处理单个连接请求。多数情况下，通过 TCP 或
是 HTTP 在某个网络地址上进行监听来创建该服务是个不错的选择。

在 RPC 客户端，Go 的  net/rpc 包提供了便利的 rpc.Dial() 和  rpc.DialHTTP() 方法
来与指定的 RPC 服务端建立连接。在建立连接之后，Go 的  net/rpc 包允许我们使用同步或者
异步的方式接收 RPC 服务端的处理结果。调用 RPC 客户端的  Call() 方法则进行同步处理，这
时候客户端程序按顺序执行，只有接收完 RPC 服务端的处理结果之后才可以继续执行后面的程
序。当调用 RPC 客户端的  Go()  方法时，则可以进行异步处理，RPC 客户端程序无需等待服务
端的结果即可执行后面的程序，而当接收到 RPC 服务端的处理结果时，再对其进行相应的处理。
无论是调用 RPC 客户端的  Call() 或者是  Go() 方法，都必须指定要调用的服务及其方法名称，
以及一个客户端传入参数的引用，还有一个用于接收处理结果参数的指针。

如果没有明确指定 RPC 传输过程中使用何种编码解码器，默认将使用 Go 标准库提供的
encoding/gob 包进行数据传输。


Gob 是 Go 的一个序列化数据结构的编码解码工具，在 Go 标准库中内置 encoding/gob 包
以供使用。一个数据结构使用 Gob 进行序列化之后，能够用于网络传输。与 JSON 或 XML 这种
基于文本描述的数据交换语言不同，Gob 是二进制编码的数据流，并且 Gob 流是可以自解释的，
它在保证高效率的同时，也具备完整的表达能力。

作为针对 Go 的数据结构进行编码和解码的专用序列化方法，这意味着 Gob 无法跨语言使
用。在 Go 的 net/rpc 包中，传输数据所需要用到的编码解码器，默认就是 Gob。由于 Gob 仅局
限于使用 Go 语言开发的程序，这意味着我们只能用 Go 的 RPC 实现进程间通信。然而，大多数
时候，我们用 Go 编写的 RPC 服务端（或客户端） ，可能更希望它是通用的，与语言无关的，无
论是Python 、 Java 或其他编程语言实现的 RPC 客户端，均可与之通信。

*/

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

// 注册服务对象并开启该 RPC 服务
arith := new(Arith)
rpc.Register(arith)
rpc.HandleHTTP()

l, e := net.Listen("tcp", ":1234")

if e != nil{
log.Fatal("Listen error:", e)
}
go http.Serve(1, nil)
// RPC 服务端注册了一个 Arith 类型的对象及其公开方法 Arith.Multiply() 和Arith.Divide() 供 RPC 客户端调用。

// RPC 在调用服务端提供的方法之前，必须先与 RPC 服务端建立连接
client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
if err != nil{
log.Fatal("dialing:", err)
}


// 在建立连接之后，RPC 客户端可以调用服务端提供的方法。
// 同步调用程序顺序执行的方式
args := &server.Args(7, 8)
var reply int
err = client.Call("Arith.Multiply", args, &reply)
if err != nil{
log.Fatal("arith error:", err)
}
fmt.Printf("Arith:%d * %d = %d\n", args.A, args.B, reply)

// 以异步方式进行调用
quotient := new(Quotient)
divCall := client.Go("Arith.Divide", args, &quotient, nil)
replyCall := <-divCall.Done

/*
设计优雅的RPC接口
Go 的 net/rpc 很灵活，它在数据传输前后实现了编码解码器的接口定义。这意味着，开发
者可以自定义数据的传输方式以及 RPC 服务端和客户端之间的交互行为。
RPC 提供的编码解码器接口如下
*/
type ClientCodec interface {
	WriteRequest(*Request, interface{}) error
	ReadResponseHeader(*Response) error
	ReadResponseBody(interface{}) error
	Close() error
}

// 接口 ClientCodec 定义了 RPC 客户端如何在一个 RPC 会话中发送请求和读取响应。 客户端程
// 序通过  WriteRequest() 方法将一个请求写入到 RPC 连接中，并通过 ReadResponseHeader()
// 和 ReadResponseBody() 读取服务端的响应信息。当整个过程执行完毕后，再通过  Close() 方
// 法来关闭该连接。

type ServerCodec interface {
	ReadRequestHeader(*Request) error
	ReadRequestBody(interface{}) error
	WriteResponse(*Response, interface{}) error
	Close() error
}

// 接口 ServerCodec 定义了 RPC 服务端如何在一个 RPC 会话中接收请求并发送响应。服务端
// 程序通过 ReadRequestHeader() 和 ReadRequestBody() 方法从一个 RPC 连接中读取请求
// 信息，然后再通过 WriteResponse()  方法向该连接中的 RPC 客户端发送响应。当完成该过程
// 后，通过 Close()  方法来关闭连接。

// 通过实现上述接口，我们可以自定义数据传输前后的编码解码方式，而不仅仅局限于 Gob。
// 同样，可以自定义RPC 服务端和客户端的交互行为。实际上，Go 标准库提供的 net/rpc/json
// 包，就是一套实现了 rpc.ClientCodec 和 rpc.ServerCodec 接口的 JSON-RPC 模块。

func main() {

}
