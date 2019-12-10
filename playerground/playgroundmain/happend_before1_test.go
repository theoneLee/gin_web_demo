package playgroundmain

import "testing"

func TestHappendBefore1(t *testing.T) {
	main1()
}

var c1 = make(chan int)
var a1 string

func f1() {
	a1 = "hello, world"
	<-c1
}

/*
从无缓冲信道进行的接收，要发生在对该信道进行的发送完成之前。

此程序（与上面的相同，但交换了发送和接收语句的位置，且使用无缓冲信道）:
也可保证打印出 "hello, world"。该程序首先对 a 进行写入， 然后从 c 中接收信号，随后向 c 发送对应的信号，最后执行 print 函数。

若该信道为带缓冲的（例如，c = make(chan int, 1)）， 则该程序将无法保证打印出 "hello, world"。（它可能会打印出空字符串， 崩溃，或做些别的事情。）
*/
func main1() {
	go f1()
	c1 <- 0
	print(a1)
}
