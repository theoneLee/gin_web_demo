package playgroundmain

import "testing"

func TestHappendBefore3(t *testing.T) {
	main3()
}

var c3 = make(chan int) //TestHappendBefore 是缓冲chan，这里是无缓冲
var a3 string

func f3() {
	a3 = "hello, world"
	c3 <- 0
}

func main3() {
	go f3()
	<-c3
	print(a3)
}
