package main
import "fmt"

func main(){
	Map()
	Map2()
	Map3()
	Map4()
	Map5()
	Map6()
	Map7()
	Map8()
	Map9()
	Map10()
	Map11()
	Map12()
	Map13()
	Map14()
}


//使用 map 不需要引入任何库
// map 是一堆键值对的未排序集合。
type PersonInfo struct{
	ID string
	Age int
	Name string
	Address string
}
func Map(){
	var personDb map[string] PersonInfo
	fmt.Println(personDb)
	personDb = make(map[string] PersonInfo)
	fmt.Println(personDb)

	//向map中插入数据
	personDb["1234"] = PersonInfo{"1234", 23, "Golang","room 1234,..."}
	personDb["1"] = PersonInfo{"1", 22, "Scala", "room 1,....."}
	fmt.Println(personDb)
	fmt.Println(len(personDb))//获取字典中元素的个数2    没有cap方法 cap(personDb))

	//查找指定key的信息
	person1, ok1 := personDb["123"]
	person2, ok2 := personDb["1"]

	fmt.Println(person1,ok1)  //{ 0  } false
	fmt.Println(person2,ok2) //{1 22 Scala room 1,.....} true

	// ok是一个返回的bool型，返回true表示找到了对应的数据
	if ok1 {
		fmt.Println(person1.Address)
	}else{
		fmt.Println("Did not find person with ID 123")
	}


	//取出值
	fmt.Println(personDb[person2.ID].Name)
	//拆分成如下两条
	fmt.Println(person2.ID)
	fmt.Println(personDb["1"].Name)

}


func Map2(){
// 1. 变量声明    map 的声明基本上没有多余的元素
	var myMap map[string] PersonInfo
	 // myMap 是声明的 map 变量名， string 是键的类型， PersonInfo 则是其中所存放的值类型

// 2. 创建    使用Go语言内置的函数 make() 来创建一个新 map
	myMap = make(map[string] PersonInfo)  //创建了一个键类型为 string 、值类型为 PersonInfo 的 map
	// 也可以选择是否在创建时指定该 map 的初始存储能力
	fmt.Println(myMap, len(myMap))  //map[] 0


	// 创建了一个初始存储能力为100的 map
	myMap = make(map[string] PersonInfo, 100)
	fmt.Println(myMap, len(myMap)) //map没有cap(myMap)方法


	// 创建并初始化 map 
	myMap = map[string] PersonInfo{
		"id" : PersonInfo{"Map", 23,"这是map","room 1234,..."},
	}
	fmt.Println(myMap, len(myMap))  //map[id:{Map 23 这是map room 1234,...}] 1


// 3. 元素赋值
	myMap["001"] = PersonInfo{"Map1",24,"another map","address...."}
	fmt.Println(myMap, len(myMap))  //map[id:{Map 23 这是map room 1234,...} 001:{Map1 24 another map address....}] 2


// 4. 元素删除    Go语言提供了一个内置函数 delete() ，用于删除容器内的元素
	delete(myMap, "id")
 	// stat := delete(myMap, "id")  //delete(myMap, "id") used as value
 	// fmt.Println(stat)
 	// 如果要删除的这个键不存在，那么这个调用将什么都不发生，也不会有什么副作用。但是如果传入的 map 变量的值是 nil ，该调用将导致程序抛出异常（panic）
	fmt.Println(myMap, len(myMap))   //map[001:{Map1 24 another map address....}] 1


// 5. 元素查找
	// 判断能否从 map 中获取一个值的常规做法是
	// (1) 声明并初始化一个变量为空；
	// (2) 试图从 map 中获取相应键的值到该变量中；
	// (3) 判断该变量是否依旧为空，如果为空则表示 map 中没有包含该变量。

// 判断变量是否为空这条语句并不能真正表意（是否成功取到对应的值），从而影响代码的可读性和可维护性。有些库甚至会设计为因为一个键不存在而抛出异常，
// 让开发者用起来胆战心惊，不得不一层层嵌套 try-catch 语句，这更是不人性化的设计。在Go语言中，要从 map 中查找一个特定的键，可以通过下面的代码来实现
	value, ok := myMap["id"]
	if ok{
		fmt.Println(ok, "存在要找的value：",value)
	}else{
		fmt.Println(ok, "不存在要找的value：",value) //false 不存在要找的value： { 0  }
	}

	value, ok = myMap["001"]
	if ok{
		fmt.Println(ok, "存在要找的value：",value) //true 存在要找的value： {Map1 24 another map address....}
	}else{
		fmt.Println(ok, "不存在要找的value：",value)
	}
// 判断是否成功找到特定的键，不需要检查取到的值是否为 nil ，只需查看第二个返回值 ok ，
// 这让表意清晰很多。配合 := 操作符，让你的代码没有多余成分，看起来非常清晰易懂。
}



func Map3(){


	
}


func Map4(){


	
}


func Map5(){



}


func Map6(){



}


func Map7(){



}


func Map8(){



}


func Map9(){



}


func Map10(){



}


func Map11(){



}


// 在函数间传递 map
func Map12(){
// 在函数间传递 map 不是传递 map 的拷贝。所以如果我们在函数中改变了 map，那么所有引用 map 的地方都会改变
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
	     "Coral":       "#ff7F50",
	     "DarkGray":    "#a9a9a9",
	     "ForestGreen": "#228b22",
	}
	fmt.Println(colors)  //map[AliceBlue:#f0f8ff Coral:#ff7F50 DarkGray:#a9a9a9 ForestGreen:#228b22]
	removeColor(colors,"Coral")
	fmt.Println(colors)  //map[AliceBlue:#f0f8ff 			   DarkGray:#a9a9a9 ForestGreen:#228b22]
	// 可以看出来传递 map 也是十分廉价的，类似 slice
}
func removeColor(col map[string]string,key string){
	delete(col,key)
}


func Map13(){
// 测试 map 的键是否存在是 map 操作的重要部分，因为它可以让我们判断是否可以执行一个操作，或者是往 map 里缓存一个值。
	// 它也可以被用来比较两个 map 的键值对是否匹配或者缺失。

// 从 map 里检索一个值有两种选择，我们可以同时检索值并且判断键是否存在
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Red": "#da1337", 
		"Orange": "#e95a22",
		"ForestGreen": "#228b22",
	}
	_ = map[string]string{"Red": "#da1337", "Orange": "#e95a22","":""}
	//不能出现重复的key
	// colors1 := map[string]string{"Red": "#da1337", "Orange": "#e95a22", "Orange": ""} //duplicate key "Orange" in map literal
	//不存在键  则exists输出false   value为空
	value, exists := colors["Blue"]  // false
	fmt.Println(value, exists)
	if exists{
		fmt.Println(value)
	}

	value1, exists1 := colors["Red"]
	//存在指定的key   则exists1输出true   并输出对应的value
	fmt.Println(value1, exists1)  //#da1337 true

	value2 := colors["Yellow"]
	// 当索引一个 map 取值时它总是会返回一个值，即使键不存在。此处就返回了对应类型的零值。
	fmt.Println("value1:", value2)
	if value2 != ""{
		fmt.Println(value2)
	}


// 迭代一个 map 和迭代数组和 slice 是一样的，使用 range 关键字，
	// 不过在迭代 map 时我们不使用 index/value 而使用 key/value 结构
	//由于字典map的元素序列是不固定的  所以遍历输出的结果也是无序的
	for key,value := range colors{
		fmt.Printf("key:%s  value:%s\n", key, value)
	}


	// 如果我们想要从 map 中移除一个键值对，使用内建函数 delete。要是也能返回移除是否成功就好了？
	fmt.Println(colors)
	// status := delete(colors,"Coral")  // delete(colors, "Coral") used as value
	delete(colors,"Coral") //不存在要移除的指定的key   也不报错
	fmt.Println(colors)
	delete(colors,"Red")
	fmt.Println(colors)  //从输出结果来看  键为Red的字段已被移除
}


// 内部机制
func Map14(){
// map 是一种无序的键值对的集合。map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
// map 是一种集合，所以我们可以像迭代数组和 slice 那样迭代它。不过，map 是无序的，我们无法决定它的返回顺序，
// 这是因为 map 是使用 hash 表来实现的。
// map 的 hash 表包含了一个桶集合(collection of buckets)。当我们存储，移除或者查找键值对(key/value pair)时，
// 都会从选择一个桶开始。在映射(map)操作过程中，我们会把指定的键值(key)传递给 hash 函数(又称散列函数)。
// hash 函数的作用是生成索引，索引均匀的分布在所有可用的桶上。


 // 创建和初始化
 // 可以使用内建函数 make 也可以使用 map 字面值
// 通过 make 来创建
	dict := make(map[string]int)
	fmt.Println(dict)  //map[]
// 通过字面值创建
	// dict2 := map[string]string{"Name":"Golang","Age":22}  //cannot use 22 (type int) as type string in map value

	dict2 := map[string]string{"Name":"Golang","Age":"22"}
	fmt.Println(dict2)   //map[Name:Golang Age:22]
	fmt.Println(len(dict2))  //cap(dict2)没有cap方法  invalid argument dict2 (type map[string]string) for cap
	// fmt.Println(dict2[0])  //cannot use 0 (type int) as type string in map index
	fmt.Println(dict2["Name"])   //Golang

// 使用字面值是创建 map 惯用的方法(为什么不使用make)。初始化 map 的长度依赖于键值对的数量。

// map 的键可以是任意内建类型或者是 struct 类型，map 的值可以是使用 ==操作符的表达式。
	// slice，function 和 包含 slice 的 struct 类型不可以作为 map 的键，否则会编译错误
	// dict3 := map[[]string]int{}   //invalid map key type []string


// 给 map 赋值就是指定合法类型的键，然后把值赋给键
	colors := map[string]string{}  //此处是一个 empty map
	fmt.Println(colors)
	colors["Red"] = "#da1337"
	fmt.Println(colors)

// 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对，否则会报运行时错误
	var colors2 map[string]string
	if colors2 == nil{
		fmt.Println("this is nil map")
	}
	// colors2["Red"] = "#da1337"  //panic: assignment to entry in nil map
}	




