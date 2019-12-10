package playgroundmain

import "testing"

func TestHappendBefore(t *testing.T) {
	main()
}

var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	c <- 0
}

/*
信道上的发送操作总在对应的接收操作完成前发生。

可保证打印出 "hello, world"。该程序首先对 a 进行写入， 然后在 c 上发送信号，随后从 c 接收对应的信号，最后执行 print 函数。

若在信道关闭后从中接收数据，接收者就会收到该信道返回的零值。

在这个例子中，用 close(c) 代替 c <- 0 仍能保证该程序产生相同的行为。
*/
func main() {
	go f()
	<-c
	print(a)
}
