package playerground

import (
	"fmt"
	"time"
)

func producer(list []int) chan int {
	//signChan:= make(chan struct{})
	c := make(chan int)
	go func() {
		defer close(c)

		for i := 0; i < len(list); i++ {
			c <- list[i]
			time.Sleep(1 * time.Second)
		}
		//close(c)
		//signChan<- struct{}{}
	}()
	//<-signChan

	return c
}

func processor(inputChan <-chan int) chan int {
	c := make(chan int)
	//signChan := make(chan struct{})
	go func() {
		//for {
		//	select {
		//	case i,ok := <-inputChan:
		//		if ok {
		//			c<-i*2
		//		}else {
		//			signChan<-struct{}{}
		//			close(c)
		//			break
		//		}
		//	default:
		//		fmt.Println("")
		//		time.Sleep(500*time.Millisecond)
		//	}
		//	if  {
		//
		//	}
		//}
		defer close(c)

		//foreach 写法
		for i := range inputChan {
			c <- i * 2
		}
		//close(c)
	}()
	return c
}

func consumer(inputChan <-chan int) {
	for i := range inputChan {
		fmt.Println(i)
	}

}

func Test() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c1 := producer(list)
	c2 := processor(c1)
	consumer(c2)
}

func producer2(list []int) chan int {
	c := make(chan int, 3)
	go func() {
		defer close(c)

		for i := 0; i < len(list); i++ {
			c <- list[i]
			//fmt.Println("sleep:",i,1)
			//c <- list[i]
			//fmt.Println("sleep:",i,2)
			//c <- list[i]
			//fmt.Println("sleep:",i,3)
			//c <- list[i]
			//fmt.Println("sleep:",i,4)
			//time.Sleep(1*time.Second)
		}
		//close(c)
	}()

	return c
}

func processor2(inputChan <-chan int, outPutChan chan int) { //
	defer close(outPutChan) // todo
	for i := range inputChan {
		time.Sleep(1 * time.Second)
		outPutChan <- i * 2
	}
}

func consumer2(outPutChan ...chan int) {
OUT:
	for {
		select {
		case i := <-outPutChan[0]:
			fmt.Println("outPutChan[0]:", i)
		case i := <-outPutChan[1]:
			fmt.Println("outPutChan[1]:", i)
		case i := <-outPutChan[2]:
			fmt.Println("outPutChan[2]:", i)
		case <-time.After(12 * time.Second):
			fmt.Println("time out")
			goto OUT
		}

	}
}
