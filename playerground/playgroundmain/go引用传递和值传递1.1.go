package main

import (
	"fmt"
	"time"
)

/*
 todo 仔细运行这三个main，分析引用传递和值传递。
 对于这块内容，要多测试下是否符合预期。
 毕竟如下面例子中，即使外层的map默认是指针类型，但map的value是Count的结构体，结构体你不显式指定引用传递他还是会作为值传递的，因此其他地方的修改无法影响到这里。
 如果一个结构体里面包含类型为切片的字段，拿更应该注意检查影响
*/

type Counter1 struct {
	count int
}

func main() {

	//var mapChan = make(chan map[string]Counter,1)
	var mapChan = make(chan map[string]Counter1, 1) // map是指针类型，Counter1是值类型 =》外部修改不会影响原本值

	syncChan := make(chan struct{}, 2)
	go func() {
		//receive
		for {
			if elem, ok := <-mapChan; ok {
				c := elem["count"]
				c.count++
			} else {
				break
			}

		}
		fmt.Println("Stopped. [receive]")
		syncChan <- struct{}{}
	}()

	go func() {
		// send
		countMap := make(map[string]Counter1)
		countMap["count"] = Counter1{} // 值
		//countMap := map[string]*Counter{
		// "count":&Counter{},
		//}
		for i := 0; i < 5; i++ {
			mapChan <- countMap          //在receiver 通道 拿到数据作出的修改，并不会影响原本的值
			time.Sleep(time.Millisecond) // 保证mapchannel接受到该值
			fmt.Printf("The count map. %v. [sender]\n", countMap)

		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan
}

func (c *Counter1) String() string {
	return fmt.Sprintf("{count: %d}", c.count)
}
