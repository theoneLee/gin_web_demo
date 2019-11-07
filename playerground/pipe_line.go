package playerground

import (
	"fmt"
	"sync"
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
	//defer close(outPutChan) // todo 如果这里关闭chan，导致consumer2不会阻塞在select，每个outPutChan会一直拿到零值，导致time.After极大可能不会执行，从而程序不会关闭
	//todo 但实际场景下，生产者，处理者，消费者都不会关闭，也不需要time.After分支跳出consumer2的执行
	for i := range inputChan {
		time.Sleep(1 * time.Second)
		outPutChan <- i * 2
	}
}

func consumer2(outPutChan ...chan int) {
	//signChan := make(chan struct{})

	for {
		select {
		case i := <-outPutChan[0]:
			fmt.Println("outPutChan[0]:", i)
		case i := <-outPutChan[1]:
			fmt.Println("outPutChan[1]:", i)
		case i, ok := <-outPutChan[2]: //ok 为false时，表示该通道已关闭
			if ok {
				fmt.Println("outPutChan[1]:", i)
			}
		case <-time.After(12 * time.Second):
			fmt.Println("time out")
			goto OUT
		}
	}
OUT:
}

func Test2() {
	outPutChan1 := make(chan int)
	outPutChan2 := make(chan int)
	outPutChan3 := make(chan int)

	c := producer([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	go processor2(c, outPutChan1)
	go processor2(c, outPutChan2)
	go processor2(c, outPutChan3)
	consumer2(outPutChan1, outPutChan2, outPutChan3)
}

func producer3(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func processor3(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
		}
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	collect := func(in <-chan int) {
		defer wg.Done()
		for n := range in {
			out <- n
		}
	}

	wg.Add(len(cs))
	// FAN-IN
	for _, c := range cs {
		go collect(c)
	}

	// 错误方式：直接等待是bug，死锁，因为merge写了out，main却没有读
	// wg.Wait()
	// close(out)

	// 正确方式
	go func() {
		wg.Wait()  //该goroutine会阻塞在这里，只有wg.Done()执行，使得信号量为0时，取消阻塞。
		close(out) //关闭该chan ，使得Test3的for ret := range merge(c1, c2, c3) 取消阻塞，跳出循环
	}()

	return out
}

func Test3() {
	in := producer3(1, 2, 3, 4)

	// FAN-OUT
	c1 := processor3(in)
	c2 := processor3(in)
	c3 := processor3(in)

	// consumer
	for ret := range merge(c1, c2, c3) {
		fmt.Printf("%3d ", ret)
	}
	fmt.Println()
}
