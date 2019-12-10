package playgroundmain

import "testing"

func TestHappendBefore2(t *testing.T) {
	main2()
}

var c2 = make(chan int, 1)
var a2 string

func f2() {
	a2 = "hello, world"
	<-c2
}

func main2() {
	go f2()
	c2 <- 0   //todo 一般做法会和TestHappendBefore一致，main goroutine 负责channel的接受，其他goroutine负责发送；则无论chan是否有缓冲都保证hello world 打印
	print(a2) //不会打印hello word ，因为从无缓冲信道进行的接收，要发生在对该信道进行的发送完成之前。（参考TestHappendBefore1）
}
