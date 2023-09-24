package main

import "fmt"
import "time"

func main() {
	// Time1()
	// Time2()
	// Time3()
	// Time4()
	// Time5()
	// Time6()
	// Time7()
	Time8()
	Time9()
	fmt.Println("time package")
}

/*
组成
	time.Duration（时长，耗时）
	time.Time（时间点）
	time.C（放时间点的管道）[ Time.C:=make(chan time.Time) ]
	time包里有2个东西，一个是时间点，另一个是时长
	时间点的意思就是“某一刻”，比如 2000年1月1日1点1分1秒 那一刻（后台记录的是unix时间，从1970年开始计算）
	时长就是某一刻与另一刻的差，也就是耗时
*/
func Time1() {
	/*
	   函数
	   	Sleep函数
	   	time.Sleep(time.Duration)
	   	表示睡多少时间，睡觉时，是阻塞状态
	*/

	fmt.Println("start slepping......")
	//time.Sleep(time.Second)
	time.Sleep(time.Second * 3) //使用乘号可以实现延长时间3秒钟

	// time.Sleep(time.Second + 30)//无法实现延长时间
	//time.Sleep(3) //无法实现暂停3秒钟
	fmt.Println("end sleep....")
	//【结果】打印start sleeping后，等了正好1秒后，打印了end sleep
	//会阻塞，Sleep时，什么事情都不会做
}

func Time2() {
	/*
	   After函数
	   time.After(time.Duration)
	   和Sleep差不多，意思是多少时间之后，但在取出管道内容前不阻塞
	*/
	fmt.Println("the 1")
	tc := time.After(time.Second) //返回一个time.C这个管道，1秒(time.Second)后会在此管道中放入一个时间点(time.Now())
	//时间点记录的是放入管道那一刻的时间值
	fmt.Println(tc) //0xc042012240

	fmt.Println("the 2")
	fmt.Println("the 3")
	<-tc //阻塞中，直到取出tc管道里的数据
	fmt.Println("the 4")
	//【结果】立即打印123，等了1秒不到一点点的时间，打印了4，结束
	//打印the 1后，获得了一个空管道，这个管道1秒后会有数据进来
	//打印the 2，（这里可以做更多事情）
	//打印the 3
	//等待，直到可以取出管道的数据（取出数据的时间与获得tc管道的时间正好差1秒钟）
	//打印the 4
}

func Time3() {
	fmt.Println("the 1")
	tc := time.After(time.Second * 2) //返回一个time.C这个管道，2秒(time.Second × 2)后会在此管道中放入
	//一个时间点(time.Now())，时间点记录的是放入管道那一刻的时间值
	fmt.Println("the 2")
	fmt.Println("the 3")
	// time.Sleep(time.Second*2)
	time.Sleep(time.Second) //这里是假设这个Println动作执行了半秒钟

	fmt.Println("the 4")
	// time.Sleep(time.Second*4)
	time.Sleep(time.Second) //这里是假设这个Println动作执行了半秒钟

	fmt.Println("the 5")
	fmt.Println("the 6")
	fmt.Println("the 7")

	<-tc //阻塞中，直到取出tc管道里的数据
	fmt.Println("the 8")
	//【结果】立即打印1和2和3，等了1秒打印了4，打完后又等了1秒打印了5，然后又立即打印了678，结束
	//这里的<-tc是立即能获得数据的
	//因为早在执行差不多Print 5的时候，管道内已经有数据了
	//当gorotine线把数据丢到管道中后，它自己阻塞了（具体请了解goroutine）
}

func Time4() {
	/*
	   AfterFunc函数

	   time.AfterFunc(time.Duration,func())
	   和After差不多，意思是多少时间之后在goroutine line执行函数
	*/
	f := func() {
		fmt.Println("Time out func(*)....")
	}
	f1 := func() {
		fmt.Println("This is func f1().....")
	}
	fmt.Println("Start run f().....")
	//下面两方法的定时执行是异步的
	time.AfterFunc(5*time.Second, f)  //从开始等待5秒钟后就执行
	time.AfterFunc(1*time.Second, f1) //从开始等待1秒钟后就执行  不受函数f的影响
	fmt.Println("End f()....")
	time.Sleep(12 * time.Second) //此处的目的 要保证主线比子线“死的晚”，否则主线死了，子线也等于死了
}

func Time5() {
	// 由于f函数不是在Main Line执行的，而是注册在goroutine Line里执行的
	// 所以一旦后悔的话，需要使用Stop命令来停止即将开始的执行，如果已经开始执行就来不及了
	houhui := true
	f := func() {
		fmt.Println("this is func f()....")
	}
	ta := time.AfterFunc(5*time.Second, f)
	time.Sleep(3 * time.Second)
	if houhui {
		ta.Stop()
	}
	time.Sleep(8 * time.Second)
	//【结果】运行了8秒多一点点后，程序退出，什么都不打印
	//注册了个f函数，打算5秒后执行
	//过了3秒后，后悔了，停掉（Stop）它
}

func Time6() {
	// Tick函数
	// time.Tick(time.Duration)
	// 和After差不多，意思是每隔多少时间后，其他与After一致
	fmt.Println("the 1")
	tc := time.Tick(time.Second) //返回一个time.C这个管道，1秒(time.Second)后会在此管道中放入一个时间点，
	//1秒后再放一个，一直反复，时间点记录的是放入管道那一刻的时间

	time.Sleep(5 * time.Second)
	for i := 1; i <= 2; i++ {
		<-tc
		fmt.Println("hello Tick")
	}
}

func Time7() {
	// time.Time的方法（time.Time自己独有的函数）
	// Before & After方法
	// 判断一个时间点是否在另一个时间点的前面（后面），返回true或false
	t1 := time.Now()
	time.Sleep(time.Second)
	t2 := time.Now()

	a := t2.After(t1) //t2的记录时间是否在t1记录时间的**后面**呢，是的话，a就是true
	fmt.Println(a)

	b := t2.Before(t1) //t2的记录时间是否在t1记录时间的**前面**呢，是的话，b就是true
	fmt.Println(b)
}

func Time8() {
	// Sub方法
	// 两个时间点相减，获得时间差（Duration）
	t1 := time.Now()
	time.Sleep(time.Second)
	t2 := time.Now()
	d := t2.Sub(t1) //时间2减去时间1
	fmt.Println(d)  //打印结果差不多为1.000123几秒，因为Sleep无法做到精确的睡1秒
	// 后发生的时间  减去   先发生时间，是正数
}

func Time9() {
	// Add方法
	// 拿一个时间点，add一个时长，获得另一个时间点
	t1 := time.Now()
	fmt.Println(t1)
	d := t1.Add(time.Hour)
	fmt.Println(d)
}
