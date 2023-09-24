package main

import (
	"fmt"
	"strconv"
	"strings"
)

var information = "公众号：CTO Plus，blog：https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q"

func stringsFunc() {
	fmt.Println(len(information))                                         // 79
	fmt.Printf("\"s\" length=%d\n", len("A"))                             // "s" length=1
	fmt.Printf("%d\n", len("中"))                                          // 3
	fmt.Println(strings.Contains(information, "CTO"))                     //true
	fmt.Println(strings.Index(information, "http"))                       // 30
	fmt.Println(strings.LastIndex(information, "o"))                      // 52
	fmt.Println(strings.Replace(information, "s", "S", 3))                // 公众号：CTO PluS，blog：httpS://mp.weixin.qq.com/S/0yqGBPbOI6QxHqK17WxU8Q
	fmt.Println(strings.ReplaceAll(information, "q", "Q"))                // 公众号：CTO Plus，blog：https://mp.weixin.QQ.com/s/0yQGBPbOI6QxHQK17WxU8Q
	fmt.Println(strings.Split(information, "/"))                          // [公众号：CTO Plus，blog：https:  mp.weixin.qq.com s 0yqGBPbOI6QxHqK17WxU8Q]
	fmt.Println(strings.SplitN(information, "/", 2))                      //[公众号：CTO Plus，blog：https: /mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q]
	fmt.Println(strings.SplitAfter(information, "/"))                     // [公众号：CTO Plus，blog：https:/ / mp.weixin.qq.com/ s/ 0yqGBPbOI6QxHqK17WxU8Q]
	fmt.Println(strings.SplitAfterN(information, "/", 2))                 // [公众号：CTO Plus，blog：https:/ /mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q]
	fmt.Println(strings.ToLower(information))                             // 公众号：cto plus，blog：https://mp.weixin.qq.com/s/0yqgbpboi6qxhqk17wxu8q
	fmt.Println(strings.ToUpper(information))                             // 公众号：CTO PLUS，BLOG：HTTPS://MP.WEIXIN.QQ.COM/S/0YQGBPBOI6QXHQK17WXU8Q
	fmt.Println(strings.HasSuffix(information, "0yqGBPbOI6QxHqK17WxU8Q")) // 以指定字符串结尾  true
	fmt.Println(strings.HasPrefix(information, "公众号"))                    // 以指定字符串开头  true

	//fmt.Println(strings.Count(information, 'o'))  // 错误用法，第二个参数必须是字符串，cannot use 'o' (type untyped rune) as type string in argument to strings.Count
	fmt.Println(strings.Count(information, "o")) // 2

	fmt.Println(strings.IndexAny("SteveRocket", "steve"))    // 1
	fmt.Println(strings.IndexAny("SteveRocket", "CTO Plus")) // -1
	fmt.Println(strings.IndexByte("SteveRocket", 'S'))       // 0
	fmt.Println(strings.IndexByte("SteveRocket", 'X'))       // -1

	fmt.Println(strings.IndexFunc("SteveRocket", func(r rune) bool {
		return r == 'R'
	})) // 5

	fmt.Println(strings.IndexRune("SteveRocket", 'S')) // 0
	fmt.Println(strings.IndexRune("SteveRocket", 'X')) // -1

	fmt.Println(strings.Compare("SteveRocket", "SteveRocket")) // 0
	fmt.Println(strings.Compare("SteveRocket", "steverocket")) // -1
	fmt.Println(strings.Compare("SteveRocket", "CTO Plus"))    // 1
	fmt.Println(strings.Compare("SteveRocket", "steve"))       // -1

	fmt.Println(strings.ContainsAny("SteveRocket", "rocket")) // true
	fmt.Println(strings.ContainsAny("SteveRocket", "plm"))    // false

	fmt.Println(strings.ContainsRune("SteveRocket", 'r')) // false
	fmt.Println(strings.ContainsRune("SteveRocket", 'X')) // false

	fmt.Println(strings.EqualFold("SteveRocket", "rocket"))      // false
	fmt.Println(strings.EqualFold("SteveRocket", "STEVEROCKET")) // true
	fmt.Println(strings.EqualFold("SteveRocket", "seven"))       // false

	fmt.Println(strings.Fields("SteveRocket"))           // [SteveRocket]
	fmt.Println(strings.Fields("   Steve   Rocket    ")) // [Steve Rocket]
	fmt.Println(strings.FieldsFunc("Steve Rocket", func(r rune) bool {
		return r == 'R'
	})) // [Steve  ocket]
}

func strconvFunc() {
	//fmt.Println(strconv.Atoi(information))
	//0 strconv.Atoi: parsing "公众号：CTO Plus，blog：https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q": invalid syntax

	//fmt.Println(strconv.ParseFloat(information, 1))
	//0 strconv.ParseFloat: parsing "公众号：CTO Plus，blog：https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q": invalid syntax

	fmt.Println(strconv.Itoa(78))                         // 78
	fmt.Println(strconv.FormatFloat(1.23456, 'f', 5, 64)) // 1.23456
	fmt.Println(strconv.FormatInt(123456, 2))             // 11110001001000000

	//strconv.Atoi()函数用于将字符串转换为整数类型。它接受一个字符串参数，返回转换后的整数值和一个错误。如果转换失败，将返回错误。
	fmt.Println(strconv.Atoi("123456")) // 123456 <nil>

	//strconv.ParseFloat()函数用于将字符串转换为浮点数类型。它接受一个字符串参数和一个位数参数，返回转换后的浮点数值和一个错误。如果转换失败，将返回错误。
	fmt.Println(strconv.ParseFloat("1.233455666", 6)) // 1.233455666 <nil>

	//strconv.Itoa()函数用于将整数转换为字符串类型。它接受一个整数参数，返回转换后的字符串。
	fmt.Println(strconv.Itoa(12345678)) // 12345678

	//strconv.FormatFloat()函数用于将浮点数转换为字符串类型。它接受一个浮点数参数、格式参数和位数参数，返回转换后的字符串。
	fmt.Println(strconv.FormatFloat(3.14, 'f', 2, 64)) // 3.14

	//strconv.FormatInt()函数用于将整数转换为字符串类型。它接受一个整数参数、基数参数和位数参数，返回转换后的字符串。
	fmt.Println(strconv.FormatInt(123, 10)) // 123

	//strconv.ParseBool()函数用于将字符串解析为布尔值。它接受一个字符串参数，返回解析后的布尔值和一个错误。如果字符串是"true"或"1"，则返回true；如果字符串是"false"或"0"，则返回false；否则返回错误。
	boolean, err := strconv.ParseBool("123")

	fmt.Println(strconv.ParseBool("true"))  // true <nil>
	fmt.Println(strconv.ParseBool("false")) // false <nil>
	if err != nil {
		fmt.Println(boolean, err) // false strconv.ParseBool: parsing "123": invalid syntax
	}

	// strconv.ParseInt()函数用于将字符串解析为整数。它接受一个字符串参数、基数参数和位数参数，返回解析后的整数值和一个错误。基数参数指定了字符串表示的数值的进制，位数参数指定了结果的位数，可以是0、8、10或16。如果解析失败，则返回错误。
	fmt.Println(strconv.ParseInt("123456", 10, 64)) // 123456 <nil>

	//strconv.ParseComplex()函数用于将字符串解析为复数。它接受一个字符串参数，返回解析后的复数值和一个错误。字符串的格式可以是"real+imagi"或"real+imagj"，其中real表示实部，imag表示虚部。如果解析失败，则返回错误。
	fmt.Println(strconv.ParseComplex("3+4i", 64)) // (3+4i) <nil>

	//strconv.ParseUint()函数用于将字符串解析为无符号整数。它接受一个字符串参数、基数参数和位数参数，返回解析后的无符号整数值和一个错误。基数参数指定了字符串表示的数值的进制，位数参数指定了结果的位数，可以是0、8、10或16。如果解析失败，则返回错误。
	fmt.Println(strconv.ParseUint("123", 10, 64)) // 123 <nil>
}

func strConvert() {
	s1 := "公众号："
	s2 := "CTO Plus"
	s3 := "，blog："
	s4 := "https://"
	s5 := "mp.weixin.qq.com"
	s6 := "/s/"
	s7 := "0yqGBPbOI6QxHqK17WxU8Q"
	merged := strings.Join([]string{s1, s2, s3, s4, s5, s6, s7}, "")
	fmt.Println(merged)                // 公众号：CTO Plus，blog：https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q
	fmt.Println(merged == information) // true

	//把多个字符串拼接成一个长的字符串有多种方式。
	//当有大量的string需要拼接时，用strings.Builder效率最高
	strBuilder := strings.Builder{}
	strBuilder.WriteString(s1)
	strBuilder.WriteString(s2)
	strBuilder.WriteString(s3)
	strBuilder.WriteString(s4)
	strBuilder.WriteString(s5)
	strBuilder.WriteString(s6)
	strBuilder.WriteString(s7)
	newInformation := strBuilder.String()
	fmt.Println(newInformation)                // 公众号：CTO Plus，blog：https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q
	fmt.Println(newInformation == information) // true

	arrInformation := []byte(newInformation)
	runeInformation := []rune(newInformation)
	// string可以转换为[]byte或[]rune类型
	fmt.Printf("byte %d\n", arrInformation[len(arrInformation)-1]) // byte 81
	//byte或rune可以转为string
	fmt.Printf("byte %c\n", arrInformation[len(arrInformation)-1]) // byte Q

	fmt.Printf("rune %d\n", runeInformation[len(runeInformation)-1]) // rune 81
	fmt.Printf("rune %c\n", runeInformation[len(runeInformation)-1]) // rune Q

	for i := 0; i < len(runeInformation); i++ {
		fmt.Printf("%c", arrInformation[i])
	}
}

func strMap() {
	fmt.Println(information) // 公众号：CTO Plus，blog：https://mp.weixin.qq.com/s/0yqGBPbOI6QxHqK17WxU8Q
	result := strings.Map(func(r rune) rune {
		if r == 'q' {
			return 'Q'
		}
		return r
	}, information)
	fmt.Println(result) // 公众号：CTO Plus，blog：https://mp.weixin.QQ.com/s/0yQGBPbOI6QxHQK17WxU8Q
}

func main() {
	//stringsFunc()
	//strconvFunc()
	//strConvert()
	strMap()
}
