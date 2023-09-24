package main

import "log"
import "os"
import "fmt"
import "encoding/json"

/*
JSON处理
	编码为json格式
	解码json数据
	解码未知结构的json数据
	json的流式读写

JSON （JavaScript Object Notation）是一种比XML更轻量级的数据交换格式，在易于人们阅
读和编写的同时，也易于程序解析和生成。尽管JSON是JavaScript的一个子集，但JSON采用完全
独立于编程语言的文本格式，且表现为键/值对集合的文本描述形式（类似一些编程语言中的字
典结构） ，这使它成为较为理想的、跨平台、跨语言的数据交换语言。

开发者可以用 JSON 传输简单的字符串、数字、布尔值，也可以传输一个数组，或者一个更
复杂的复合结构。在 Web 开发领域中，JSON被广泛应用于 Web 服务端程序和客户端之间的数据
通信，但也不仅仅局限于此，其应用范围非常广阔，比如作为Web Services API输出的标准格式，
又或是用作程序网络通信中的远程过程调用（RPC）等。

Go语言内建对JSON的支持。使用Go语言内置的 encoding/json 标准库，开发者可以轻松
使用Go程序生成和解析JSON格式的数据。在Go语言实现JSON的编码和解码时，遵循RFC4627
协议标准。
*/

// 编码为JSON格式
// 使用 json.Marshal() 函数可以对一组数据进行JSON格式的编码
// json.Marshal() 函数的声明
// func Marshal(v interface{})([]byte, error)

// Book 类型的结构体
type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float32
}

// Book 类型的实例对象

func JsonEncode() {
	// gobook := Book{"Go语言编程",{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"},"ituring.com.cn",true,9.99}

	// 使用 json.Marshal()  函数将 gobook 实例生成一段JSON格式的文本
	// b, err := json.Marshal(gobook)
	// 如果编码成功， err 将赋于零值 nil ，变量 b  将会是一个进行JSON格式化之后的 []byte类型
	// fmt.Println(b)
	// 	当我们调用 json.Marshal(gobook) 语句时， 会递归遍历 gobook 对象， 如果发现 gobook 这个
	// 数据结构实现了 json.Marshaler 接口且包含有效的值， Marshal() 就会调用其 MarshalJSON()
	// 方法将该数据结构生成 JSON 格式的文本。

	// 	Go语言的大多数数据类型都可以转化为有效的JSON文本， 但channel、 complex和函数这几种类型除外。
	// 如果转化前的数据结构中出现指针， 那么将会转化指针所指向的值， 如果指针指向的是零值，那么 null 将作为转化后的结果输出。
}

// 在Go中，JSON转化前后的数据类型映射如下。
// 		布尔值转化为JSON后还是布尔类型。
// 		浮点数和整型会被转化为JSON里边的常规数字。
// 		字符串将以UTF-8编码转化输出为Unicode字符集的字符串， 特殊字符比如 < 将会被转义为\u003c 。
// 		数组和切片会转化为JSON里边的数组，但 []byte 类型的值将会被转化为 Base64 编码后的字符串， slice 类型的零值会被转化为  null 。
// 		结构体会转化为JSON对象，并且只有结构体里边以大写字母开头的可被导出的字段才会被转化输出，而这些可导出的字段会作为JSON对象的字符串索引。
// 		转化一个 map 类型的数据结构时，该数据的类型必须是  map[string]T （ T 可以是encoding/json 包支持的任意数据类型） 。

// 解码JSON数据
// 可以使用 json.Unmarshal() 函数将JSON格式的文本解码为Go里边预期的数据结构。
// json.Unmarshal() 函数的原型如下
// func Unmarshal(data []byte, v interface{}) error

// 该函数的第一个参数是输入，即JSON格式的文本（比特序列） ，第二个参数表示目标输出容器，用于存放解码后的值。
// 要解码一段JSON数据，首先需要在Go中创建一个目标类型的实例对象，用于存放解码后的值：  var book Book
// 然后调用  json.Unmarshal() 函数，将 []byte  类型的JSON数据作为第一个参数传入，将 book实例变量的指针作为第二个参数传入
// err := json.Unmarshal(b, &book)
// 如果  b 是一个有效的JSON数据并能和  book 结构对应起来，那么JSON解码后的值将会一一存放到 book 结构体中。

// Go是如何将JSON数据解码后的值一一准确无误地关联到一个数据结构中的相应字段呢？
// 实际上， json.Unmarshal() 函数会根据一个约定的顺序查找目标结构中的字段，如果找到
// 一个即发生匹配。假设一个JSON对象有个名为 "Foo" 的索引，要将 "Foo" 所对应的值填充到目标
// 结构体的目标字段上， json.Unmarshal() 将会遵循如下顺序进行查找匹配：
// (1) 一个包含 Foo 标签的字段；
// (2) 一个名为 Foo 的字段；
// (3) 一个名为 Foo 或者 Foo 或者除了首字母其他字母不区分大小写的名为 Foo 的字段。
// 这些字段在类型声明中必须都是以大写字母开头、可被导出的字段。

// 但是当JSON数据里边的结构和Go里边的目标类型的结构对不上时，会发生什么呢？

// 如果JSON中的字段在Go目标类型中不存在， json.Unmarshal() 函数在解码过程中会丢弃
// 该字段。在上面的示例代码中，由于 Sales 字段并没有在 Book 类型中定义，所以会被忽略，只有
// Title 这个字段的值才会被填充到 gobook.Title 中。
func JsonDecode() {
	b := []byte(`{"Title": "Go语言编程", "Sales": 1000000}`)
	var gobook Book
	err := json.Unmarshal(b, &gobook)
	fmt.Println(err)
	fmt.Println(b)
}

// 这个特性让我们可以从同一段JSON数据中筛选指定的值填充到多个Go语言类型中。当然，
// 前提是已知JSON数据的字段结构。这也同样意味着，目标类型中不可被导出的私有字段（非首
// 字母大写）将不会受到解码转化的影响。

// 但如果JSON的数据结构是未知的，应该如何处理呢？
// 解码未知结构的JSON数据
// 接口是一组预定义方法的组合，任何一个
// 类型均可通过实现接口预定义的方法来实现，且无需显示声明，所以没有任何方法的空接口可以
// 代表任何类型。换句话说，每一个类型其实都至少实现了一个空接口。

// Go内建这样灵活的类型系统，向我们传达了一个很有价值的信息：空接口是通用类型。如
// 果要解码一段未知结构的JSON，只需将这段JSON数据解码输出到一个空接口即可。在解码JSON
// 数据的过程中，JSON数据里边的元素类型将做如下转换：
//JSON中的布尔值将会转换为Go中的 bool 类型；
//数值会被转换为Go中的 float64 类型；
//字符串转换后还是 string 类型；
//JSON数组会转换为 []interface{} 类型；
//JSON对象会转换为 map[string]interface{} 类型；
//null 值会转换为 nil 。

// 在Go的标准库 encoding/json 包中，允许使用 map[string]interface{} 和 []interface{}类型的值来分别存放未知结构的JSON对象或数组
func decodeType() {
	b := []byte(`{
	"Title": "Go语言编程",
	"Authors": ["XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan",
	"XuDaoli"],
	"Publisher": "ituring.com.cn",
	"IsPublished": true,
	"Price": 9.99,
	"Sales": 1000000
	}`)
	var r interface{}            //r 被定义为一个空接口。
	err := json.Unmarshal(b, &r) // json.Unmarshal() 函数将一个JSON对象解码到空接口 r 中，最终 r 将会是一个键值对的  map[string]interface{} 结构
	_ = err

	/*map[string]interface{}{
	"Title": "Go语言编程",
	"Authors": ["XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan",
	"XuDaoli"],
	"Publisher": "ituring.com.cn",
	"IsPublished": true,
	"Price": 9.99,
	"Sales": 1000000
	}
	要访问解码后的数据结构，需要先判断目标结构是否为预期的数据类型
	*/
	gobook, ok := r.(map[string]interface{})
	// 通过 for 循环搭配 range 语句一一访问解码后的目标数据
	if ok {
		for k, v := range gobook {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, " is string", v2)
			case int:
				fmt.Println(k, " is int", v2)
			case bool:
				fmt.Println(k, " is bool", v2)
			case []interface{}:
				fmt.Println(k, " is an array", v2)
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println(k, "is another type not handle yet")
			}
		}
	}
}

// 虽然有些烦琐，但的确是一种解码未知结构的JSON数据的安全方式

// JSON的流式读写
// Go内建的 encoding/json 包还提供 Decoder 和 Encoder 两个类型，用于支持JSON数据的流式读写，并提供 NewDecoder() 和 NewEncoder() 两个函数来便于具体实现
// func NewDecoder(r io.Reader) *Decoder
// func NewEncoder(w io.Writer) *Encoder

// 从标准输入流中读取JSON数据，然后将其解码，但只保留 Title 字段（书名），再写入到标准输出流中
func DecoderAndEncoder() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}

		for k := range v {
			if k != "Title" {
				v[k] = nil //, false
			}
		}

		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}

// 使用 Decoder 和 Encoder 对数据流进行处理可以应用得更为广泛些，比如读写 HTTP 连接、
// WebSocket 或文件等，Go的标准库 net/rpc/jsonrpc 就是一个应用了 Decoder 和 Encoder 的实际例子。

func main() {
	// JsonEncode()
	// JsonDecode()
	// decodeType()
	DecoderAndEncoder()
}
