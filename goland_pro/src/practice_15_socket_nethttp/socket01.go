package main

import "fmt"

// import "net"
// import "io"
import "net/http"

//支持基于IP层、TCP/UDP层及更高层面（如HTTP、FTP、SMTP）的网络操作，其中用于IP层的称为RawSocket。

/*
Socket编程
在Go语言中编写网络程序时，我们将看不到传统的编码形式。以前我们使用Socket编程时，会按照如下步骤展开
	(1) 建立Socket：使用 socket() 函数。
	(2) 绑定Socket：使用 bind() 函数。
	(3) 监听：使用 listen() 函数。或者连接：使用 connect() 函数。
	(4) 接受连接：使用 accept() 函数。
	(5) 接收：使用 receive() 函数。或者发送：使用 send() 函数。
Go语言标准库对此过程进行了抽象和封装。无论我们期望使用什么协议建立什么形式的连接，都只需要调用 net.Dial() 即可。

	网络通信
	Dial()函数
	func Dial(net, addr string)(Conn, error)
其中 net 参数是网络协议的名字， addr 参数是IP地址或域名，而端口号以“:”的形式跟随在地址或域名的后面，端口号可选。如果连接成功，返回连接对象，否则返回 error 。

	ICMP 链接（使用协议名称）
	conn, err := net.Dial("ip4:icmp", "www.baidu.com")

	ICMP链接（使用协议编号）
	conn, err := net.Dial("ip4:1", "10.0.0.3")

	TCP
	conn, err := net.Dial("tcp", "localhost:2100")

	UDP
	conn, err := net.Dial("udp", "localhost:975")

 	Dial() 函数支持如下几种网络协议:
 	tcp tcp4（仅限IPv4） tcp6（仅限IPv6）
 	udp udp4（仅限IPv4） udp6（仅限IPv6）
 	ip  ip4（仅限IPv4）  ip6（仅限IPv6）
	发送数据时，使用 conn 的 Write()成员方法，接收数据时使用 Read() 方法。
	实际上， Dial() 函数是对 DialTCP() 、 DialUDP() 、 DialIP() 和 DialUnix() 的封装。
	可以直接调用这些函数，它们的功能是一致的。这些函数的原型如下
		func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err error)
		func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err error)
		func DialIP(netProto string, laddr, raddr *IPAddr) (*IPConn, error)
		func DialUnix(net string, laddr, raddr *UnixAddr) (c *UnixConn, err error)

 	net 包中还包含了一系列的工具函数，合理地使用这些函数可以更好地保障程序的质量。
	验证IP地址有效性的代码如下：
		func net.ParseIP()
	创建子网掩码的代码如下：
		func IPv4Mask(a, b, c, d byte) IPMask
	获取默认子网掩码的代码如下：
		func (ip IP) DefaultMask() IPMask
	根据域名查找IP的代码如下：
		func ResolveIPAddr(net, addr string) (*IPAddr, error)
		func LookupHost(name string) (cname string, addrs []string, err error)；

HTTP编程
	HTTP（HyperText Transfer Protocol，超文本传输协议）是互联网上应用最为广泛的一种网络协议，定义了客户端和服务端之间请求与响应的传输标准。
	Go语言标准库内建提供了 net/http 包，涵盖了HTTP客户端和服务端的具体实现。使用net/http 包，我们可以很方便地编写HTTP客户端或服务端的程序。

	Go内置的 net/http 包提供了最简洁的HTTP客户端实现，我们无需借助第三方网络通信库（比如 libcurl ）就可以直接使用HTTP中用得最多的GET和POST方式请求数据。
	net/http 包的 Client 类型提供了如下几个方法，让我们可以用最简洁的方式实现 HTTP请求
		func (c *Client) Get(url string) (r *Response, err error)
		func (c *Client) Post(url string, bodyType string, body io.Reader) (r *Response, err error)
		func (c *Client) PostForm(url string, data url.Values) (r *Response, err error)
		func (c *Client) Head(url string) (r *Response, err error)
		func (c *Client) Do(req *Request) (resp *Response, err error)


	Server
	如何处理HTTP请求和HTTPS请求
处理HTTP请求
	使用 net/http  包提供的  http.ListenAndServe() 方法，可以在指定的地址进行监听，
	开启一个HTTP，服务端该方法的原型如下
		func ListenAndServe(addr string, handler Handler) error

	该方法用于在指定的 TCP 网络地址 addr 进行监听，然后调用服务端处理程序来处理传入的连
接请求。该方法有两个参数：第一个参数 addr  即监听地址；第二个参数表示服务端处理程序，
通常为空，这意味着服务端调用  http.DefaultServeMux  进行处理，而服务端编写的业务逻
辑处理程序  http.Handle() 或  http.HandleFunc() 默认注入  http.DefaultServeMux 中，
具体代码如下
http.Handle("/foo", fooHandler)
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})
log.Fatal(http.ListenAndServe(":8080", nil))

如果想更多地控制服务端的行为，可以自定义  http.Server
s := &http.Server{
	Addr: ":8080",
	Handler: myHandler,
	ReadTimeout: 10 * time.Second,
	WriteTimeout: 10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())


处理HTTPS请求
	net/http 包还提供  http.ListenAndServeTLS() 方法，用于处理 HTTPS 连接请求
	func ListenAndServeTLS(addr string, certFile string, keyFile string, handler Handler)
		error

	ListenAndServeTLS() 和 ListenAndServe() 的行为一致，区别在于只处理HTTPS请求。
	此外，服务器上必须存在包含证书和与之匹配的私钥的相关文件，比如 certFile 对应SSL证书
	文件存放路径， keyFile 对应证书私钥文件路径。如果证书是由证书颁发机构签署的， certFile
	参数指定的路径必须是存放在服务器上的经由CA认证过的SSL证书。
		开启 SSL 监听服务也很简单，如下列代码所示
			http.Handle("/foo", fooHandler)
			http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
			})
			log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil))

			或者是：
			ss := &http.Server{
				Addr: ":10443",
				Handler: myHandler,
				ReadTimeout: 10 * time.Second,
				WriteTimeout: 10 * time.Second,
				MaxHeaderBytes: 1 << 20,
			}
			log.Fatal(ss.ListenAndServeTLS("cert.pem", "key.pem"))





	Client
	设计优雅的 HTTP Client
	Go语言标准库提供的 HTTP Client 是相当优雅的。一方面提供了极其简单的使用方式，另一方面又具备极大的灵活性。

	Go语言标准库提供的HTTP Client 被设计成上下两层结构。一层是上述提到的 http.Client
	类及其封装的基础方法，我们不妨将其称为“业务层”。之所以称为业务层，是因为调用方通常
	只需要关心请求的业务逻辑本身，而无需关心非业务相关的技术细节，这些细节包括：
	  	HTTP 底层传输细节
	  HTTP 代理
	  gzip 压缩
	  连接池及其管理
	  认证（SSL或其他认证方式）
	之所以HTTP Client可以做到这么好的封装性，是因为HTTP Client 在底层抽象了
	http.RoundTripper 接口，而 http.Transport 实现了该接口，从而能够处理更多的细节，我
	们不妨将其称为“传输层”。HTTP Client 在业务层初始化 HTTP Method、目标URL、请求参数、
	请求内容等重要信息后，经过“传输层”，“传输层”在业务层处理的基础上补充其他细节，然后
	再发起 HTTP 请求，接收服务端返回的 HTTP 响应。

=======================================
自定义 http.Client
	使用的 http.Get() 、 http.Post() 、 http.PostForm() 和 http.Head() 方法其实都是在 http.DefaultClient 的基础上进行调用的，
	比如 http.Get() 等价于 http.DefaultClient.Get() ，依次类推

	http.DefaultClient 在字面上就向我们传达了一个信息，既然存在默认的 Client，那么HTTP Client 大概是可以自定义的。
	实际上确实如此，在 net/http 包中，的确提供了 Client 类型。
	 http.Client 类型的结构
		type Client struct {
			// Transport用于确定HTTP请求的创建机制。
			// 如果为空，将会使用DefaultTransport
			Transport RoundTripper
			// CheckRedirect定义重定向策略。
			// 如果CheckRedirect不为空，客户端将在跟踪HTTP重定向前调用该函数。
			// 两个参数req和via分别为即将发起的请求和已经发起的所有请求，最早的
			// 已发起请求在最前面。
			// 如果CheckRedirect返回错误，客户端将直接返回错误，不会再发起该请求。
			// 如果CheckRedirect为空，Client将采用一种确认策略，将在10个连续
			// 请求后终止
			CheckRedirect func(req *Request, via []*Request) error
			// 如果Jar为空，Cookie将不会在请求中发送，并会
			// 在响应中被忽略
			Jar CookieJar
		}
	在Go语言标准库中， http.Client 类型包含了3个公开数据成员：
	Transport RoundTripper
	CheckRedirect func(req *Request, via []*Request) error
	Jar CookieJar

 Transport 类型必须实现 http.RoundTripper 接口。 Transport 指定了执行一个HTTP 请求的运行机制，倘若不指定具体的 Transport ，默认会使用 http.DefaultTransport ，
这意味着 http.Transport 也是可以自定义的。 net/http 包中的 http.Transport 类型实现了http.RoundTripper 接口。

CheckRedirect 函数指定处理重定向的策略。当使用 HTTP Client 的 Get() 或者是 Head()方法发送 HTTP 请求时，若响应返回的状态码为
30x （比如 301 / 302 / 303 / 307），HTTP Client 会在遵循跳转规则之前先调用这个 CheckRedirect 函数。

Jar 可用于在 HTTP Client 中设定 Cookie， Jar 的类型必须实现了  http.CookieJar 接口，该接口预定义了  SetCookies() 和 Cookies() 两个方法。
如果 HTTP Client 中没有设定  Jar ，Cookie将被忽略而不会发送到客户端。实际上，我们一般都用  http.SetCookie() 方法来设定Cookie。


使用自定义的 http.Client 及其 Do() 方法，我们可以非常灵活地控制 HTTP 请求，比如发送自定义 HTTP Header 或是改写重定向策略等。
创建自定义的 HTTP Client 非常简单，具体代码如下
client := &http.Client {
	CheckRedirect: redirectPolicyFunc,
}
resp, err := client.Get("http://example.com")
// ...
req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("User-Agent", "Our Custom User-Agent")
req.Header.Add("If-None-Match", `W/"TheFileEtag"`)
resp, err := client.Do(req)
// ...

===========================================
自定义 http.Transport
在 http.Client 类型的结构定义中，我们看到的第一个数据成员就是一个  http.Transport对象，该对象指定执行一个 HTTP 请求时的运行规则。
http.Transport  类型的具体结构：
	type Transport struct {
		// Proxy指定用于针对特定请求返回代理的函数。
		// 如果该函数返回一个非空的错误，请求将终止并返回该错误。
		// 如果Proxy为空或者返回一个空的URL指针，将不使用代理
		Proxy func(*Request) (*url.URL, error)   //Proxy 指定了一个代理方法，该方法接受一个 *Request 类型的请求实例作为参数并返回
		一个最终的 HTTP 代理。如果  Proxy 未指定或者返回的 *URL 为零值，将不会有代理被启用。

		// Dial指定用于创建TCP连接的dail()函数。
		// 如果Dial为空，将默认使用net.Dial()函数
		Dial func(net, addr string) (c net.Conn, err error) //Dial 指定具体的 dial() 方法来创建 TCP 连接。如果不指定，默认将使用  net.Dial() 方法。


		// TLSClientConfig指定用于tls.Client的TLS配置。
		// 如果为空则使用默认配置
		TLSClientConfig *tls.Config   	//SSL连接专用， TLSClientConfig 指定 tls.Client 所用的 TLS 配置信息，如果不指定，也会使用默认的配置。
		DisableKeepAlives bool    		//是否取消长连接，默认值为 false ，即启用长连接。
		DisableCompression bool  		//是否取消压缩（GZip），默认值为  false ，即启用压缩。
		// 如果MaxIdleConnsPerHost为非零值，它用于控制每个host所需要
		// 保持的最大空闲连接数。如果该值为空，则使用DefaultMaxIdleConnsPerHost
		MaxIdleConnsPerHost int 		//指定与每个请求的目标主机之间的最大非活跃连接（keep-alive）数量。如果不指定，默认使
用  DefaultMaxIdleConnsPerHost 的常量值。
	}
定义了 http.Transport 类型中的公开数据成员

除了 http.Transport 类型中定义的公开数据成员以外，它同时还提供了几个公开的成员方法
	func(t *Transport) CloseIdleConnections() 。该方法用于关闭所有非活跃的连接。

	func(t *Transport) RegisterProtocol(scheme string, rt RoundTripper) 。
	该方法可用于注册并启用一个新的传输协议，比如 WebSocket 的传输协议标准（ws），或者 FTP、File 协议等。

	func(t *Transport) RoundTrip(req *Request) (resp *Response, err error) 。
	用于实现  http.RoundTripper 接口。


自定义 http.Transport 也很简单，如下列代码所示
tr := &http.Transport{
	TLSClientConfig: &tls.Config{RootCAs: pool},
	DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")

Client 和 Transport 在执行多个 goroutine 的并发过程中都是安全的，但出于性能考虑，应当创建一次后反复使用



=========================================
灵活的 http.RoundTripper 接口
HTTP Client是可以自定义的，而 http.Client 定义的第一个公开成员就是一个 http.Transport 类型的实例，且该成员所对应的类型必须实现http.RoundTripper 接口。
http.RoundTripper 接口的具体定义：
	type RoundTripper interface {
		// RoundTrip执行一个单一的HTTP事务，返回相应的响应信息。
		// RoundTrip函数的实现不应试图去理解响应的内容。如果RoundTrip得到一个响应，
		// 无论该响应的HTTP状态码如何，都应将返回的err设置为nil。非空的err
		// 只意味着没有成功获取到响应。
		// 类似地，RoundTrip也不应试图处理更高级别的协议，比如重定向、认证和
		// Cookie等。
		//
		// RoundTrip不应修改请求内容, 除非了是为了理解Body内容。每一个请求
		// 的URL和Header域都应被正确初始化
		RoundTrip(*Request) (*Response, error)
	}

http.RoundTripper 接口很简单，只定义了一个名为 RoundTrip的方法。任何实现了  RoundTrip()  方法的类型即可实现 http.RoundTripper 接口。
http.Transport 类型正是实现了  RoundTrip() 方法继而实现了该接口

http.RoundTripper 接口定义的  RoundTrip() 方法用于执行一个独立的 HTTP 事务，接
受传入的  \*Request  请求值作为参数并返回对应的  \*Response 响应值，以及一个 error 值。
在实现具体的  RoundTrip() 方法时，不应该试图在该函数里边解析 HTTP 响应信息。若响应成
功， error 的值必须为 nil ，而与返回的 HTTP 状态码无关。若不能成功得到服务端的响应， error
必须为非零值。类似地，也不应该试图在  RoundTrip() 中处理协议层面的相关细节，比如重定
向、认证或是 cookie 等。

非必要情况下，不应该在  RoundTrip() 中改写传入的请求体（ \*Request ），请求体的内
容（比如 URL 和 Header 等）必须在传入  RoundTrip() 之前就已组织好并完成初始化。

可以在默认的 http.Transport 之上包一层  Transport 并实现 RoundTrip()方法
*/
type OurCustomTransport struct {
	Transport http.RoundTripper
}

func (t *OurCustomTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *OurCustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// 处理一些事情 ...
	// 发起HTTP请求
	// 添加一些域到req.Header中
	return t.transport().RoundTrip(req)
}

func (t *OurCustomTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

// 因为实现了 http.RoundTripper 接口的代码通常需要在多个 goroutine中并发执行，因此我们必须确保实现代码的线程安全性。
func main() {
	var url = "http://10.11.115.82:8080/agent/state/bb7537c4b0e3ee7882ddc46133aa2991"
	t := &OurCustomTransport{
		// fmt.Println("OurCustomTransport.......")
	}

	c := t.Client()
	resp, err := c.Get(url)
	fmt.Println(resp)
	fmt.Println(err)

	// fmt.Println(httpGet(url))

	// fmt.Println()
	// httpHead(url)
}

// http.Get()
// 要请求一个资源，只需调用 http.Get() 方法（等价于 http.DefaultClient.Get() ）即可
// 请求一个网站首页，并将其网页内容打印到标准输出流中。
func httpGet(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Request bad:", err)
		return 404
	}
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)
	fmt.Println(resp.Body)
	fmt.Println(resp)
	// fmt.Println(resp.Code)
	fmt.Println(resp.StatusCode) //请求成功返回  200
	fmt.Println(http.StatusOK)   //请求成功返回  200
	return 201
}

// http.Post()
// 要以POST的方式发送数据，也很简单，只需调用 http.Post() 方法并依次传递下面的3个参数即可
// 	 请求的目标 URL
// 将要 POST 数据的资源类型(MIMEType)
// 数据的比特流([]byte 形式)
func httpPost(url string) {
	// resp, err := http.Post(url, "image/jpeg", &imageDataBuf)
	// if err != nil {
	// 	//错误处理
	// 	return
	// }
	// if resp.StatusCode != http.StatusOK{
	// 	//错误处理
	// 	return
	// }
}

// http.PostForm()
// http.PostForm() 方法实现了标准编码格式为 application/x-www-form-urlencoded的表单提交。下面的示例代码模拟HTML表单提交一篇新文章
func httpPostForm(url string) {
	// resp, err := http.PostForm(url, url.Values{"title":{"article title"}, "content": {"article body"}})
	// if err != nil{
	// 	//错误处理
	// 	return
	// }
}

// http.Head()
// HTTP 中的 Head 请求方式表明只请求目标 URL 的头部信息，即 HTTP Header 而不返回 HTTPBody。Go 内置的  net/http 包同样也提供了  http.Head() 方法，
// 该方法同  http.Get() 方法一样，只需传入目标 URL 一个参数即可。下面的示例代码请求一个网站首页的 HTTP Header信息
func httpHead(url string) {
	resp, err := http.Head(url)
	if err != nil {
		return
	}
	fmt.Println(resp)
	fmt.Println(resp.StatusCode)
}

// (*http.Client).Do()
// 在多数情况下， http.Get() 和 http.PostForm() 就可以满足需求，但是如果我们发起的HTTP 请求需要更多的定制信息，
// 我们希望设定一些自定义的 Http Header 字段，比如：
// 		设定自定义的 "User-Agent" ，而不是默认的 "Go http package"
// 		传递 Cookie
// 此时可以使用 net/http 包 http.Client 对象的 Do() 方法来实现
func httpClientDo(url string) {
	// req, err := http.NewRequest("GET", url, nil)
	// req.Header.Add("User-Agent", "Gobook Custom User-Agent")
	// client := &http.Client{
	// 	//...
	// }
	// 	resp, err := client.Do(req)
}
