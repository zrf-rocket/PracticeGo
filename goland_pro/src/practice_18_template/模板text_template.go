package main

import (
	"fmt"
	"os"
	"text/template"
)

/*
go模板的使用
模板就是将一组文本嵌入另一组文本里
*/
func main() {
	// 传入string--最简单的替换
	// TemplateString()

	// 传入struct
	// TemplateStruct()

	// 多模板，介绍New，Name，Lookup
	//MultiTemplate()

	// 文件模板，介绍ParseFiles
	// FileTemplateParseFiles()

	// FileTemplateParseGlob()

	// 模板的输出，介绍ExecuteTemplate和Execute
	//TemplateEcho()

	// 模板的复用
	//TemplateReuse()

	// 模板的回车
	TemplateReturn()
}

// 传入string--最简单的替换
func TemplateString() {
	name := "GoLang Template replace"

	// templ, err := template.New("test").Parse("hello,{{.}}")  //建立一个模板，内容是"hello, {{.}}"
	//等价
	templates := "hello,{{.}}"
	templ, err := template.New("test").Parse(templates)

	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	err = templ.Execute(os.Stdout, name) //将string与模板合成，变量name的内容会替换掉{{.}}
	fmt.Println(err)
	//合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}
	fmt.Println(templ)
}

// 传入struct
// 模板合成那句，第2个参数是interface{}，所以可以传入任何类型，现在传入struct看看 要取得struct的值，只要使用成员名字即可
type Inventory struct {
	Material string
	Count    uint
}

func TemplateStruct() {
	sweaters := Inventory{"Struct", 25}
	//取出结构体中的成员值
	fmt.Println(sweaters)
	fmt.Println(sweaters.Material)
	fmt.Println(sweaters.Count)

	//建立一个模板
	muban := "{{.}} {{.Count}} items are made of {{.Material}}"
	templ, err := template.New("test").Parse(muban)

	if err != nil {
		panic(err)
	}
	err = templ.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

// 多模板，介绍New，Name，Lookup
func MultiTemplate() {
	/*
			sweaters := Inventory{"Template", 24}
		//一个模板可以有多种，以Name来区分
			muban_eng := "{{.Count}} items are made of {{.Material}}"
			muban_chn := "{{.Material}} 模板有多种， {{.Count}}个模板"

			//建立一个模板的名称是china，模板的内容是muban_chn字符串
			tmpl, err := template.New("china")
			tmpl, err = templ.Parse(muban_chn)
			//建立一个模板的名称是english，模板的内容是muban_eng字符串
			tmpl, err = template.New("english")
			tmpl, err = tmpl.Parse(muban_eng)

			//将struct与模板合成，用名字是china的模板进行合成，结果放到os.Stdout里
			err = tmpl.ExecuteTemplate(os.Stdout, "china", sweaters)
			//将struct与模板合成，用名字是english的模板进行合成，结果放到os.Stdout里
			err = tmpl.ExecuteTemplate(os.Stdout, "english", sweaters)


			tmpl, err = template.New("english")
			fmt.Println(tmpl.Name())
			tmpl, err = template.New("china")
			fmt.Println(tmpl.Name())

			tmpl = tmpl.Lookup("english") //必须要有返回，否则不生效
			fmt.Println(tmpl.Name())
	*/
}

// 文件模板，介绍ParseFiles
func FileTemplateParseFiles() {
	structs := Inventory{"templates", 22}

	// 模板可以是一行或者多行
	muban1 := "{{.Count}}  items are made of {{.Material}}"
	muban2 := `items number is {{.Count}}
	there made of {{.Material}}
	`

	templ, err := template.New("test").Parse(muban1)
	if err != nil {
		panic(err)
	}
	templ.Execute(os.Stdout, structs)

	fmt.Println()
	temp, err1 := template.New("test").Parse(muban2)
	if err1 != nil {
		panic(err1)
	}
	temp.Execute(os.Stdout, structs)

	fmt.Println()
	// 把模板的内容发在一个文本文件里，用的时候将文本文件里的所有内容赋值给muban这个变量即可
	temp2, err2 := template.ParseFiles("template.txt")
	temp2.Execute(os.Stdout, structs)
	if err2 != nil {
		panic(err2)
	}
}

// ParseFiles接受一个字符串，字符串的内容是一个模板文件的路径（绝对路径or相对路径）
// ParseGlob也差不多，是用正则的方式匹配多个文件
func FileTemplateParseGlob() {
	structs := Inventory{"<<templates>>", 22}
	templ, err := template.ParseGlob("*.txt")
	templ.Execute(os.Stdout, structs)
	if err != nil {
		panic(err)
	}
}

// 模板的输出，介绍ExecuteTemplate和Execute
func TemplateEcho() {
	// 模板下有多套模板，其中有一套模板是当前模板
	// 可以使用Name的方式查看当前模板
	//sweaters := Inventory{"Templates", 33}

	//err = templ.ExecuteTemplate(os.Stdout, "english", sweaters) //指定模板名，这次为english

	//err = templ.Execute(os.Stdout, sweaters)//模板名省略，打印的是当前模板
}

// 模板的复用
// 模板里可以套模板，以达到复用目的，用template关键字
func TemplateReuse() {
	muban1 := `hi, {{template "M2"}},
	hello, {{template "M3"}}
	`
	muban2 := `这是模板2  {{template "M3"}}`
	muban3 := "i'm templates 3"

	templ, err := template.New("M1").Parse(muban1)
	if err != nil {
		panic(err)
	}

	templ.New("M2").Parse(muban2)
	if err != nil {
		panic(err)
	}

	templ.New("M3").Parse(muban3)
	if err != nil {
		panic(err)
	}

	err = templ.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
}

// 模板的回车
// 模板文件里的回车也是模板的一部分，如果对回车位置控制不好，合成出来的文章会走样
// 标准库 https://gowalker.org/text/template#Template
func TemplateReturn() {
	const letter = `Dear {{.Name}},
		{{if .Attended}}It was a pleasure to see you at the wedding.
		如果Attended是true的话，这句是第二行{{else}}It is a shame you couldn't make it to the wedding.
		如果Attended是false的话，这句是第二行{{end}}
		{{with .Gift}}Thank you for the lovely {{.}}.
		{{end}}
		Best wishes,
		Josie
	`
	fmt.Println(letter)

}

/*
正确的正文排版如下
如果正文就一行，要把true和false的所有内容都写在一行
比如{{if .Attended}}true line,hello true{{else}}false line,hi false{{end}}
如果正文有多行，就等于把一行拆成多行
会发现true的最后一行和false的第一行是在同一行
{{if .Attended}}和ture的第一行在同一行
{{end}}和false的最后一行在同一行
如下：
{{if .Attended}}true line
hello true{{else}}false line
hi false{{end}}


关于{{with .Gift}},意思是如果Gift不是为空的话，就打印整行，如果为空，就不打印
只有这样写法，with对应的end要写在第2行，才会把“Thank you”这句后面带一个回车进去，这样写法，就像“Thank you”这句是插在正文下面的
只有这样写，不管有没有“Thank you”，正文和Best wishes,之间始终只有1行空白
*/
