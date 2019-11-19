package playgroundmain

import (
	"fmt"
	"github.com/moby/moby/pkg/pubsub"
	"strings"
	"testing"
	"time"
)

//其中pubsub.NewPublisher构造一个发布对象，p.SubscribeTopic()可以通过函数筛选感兴趣的主题进行订阅。
func TestPubSub(t *testing.T) {
	fmt.Println("111")
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})
	docker := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}
		return false
	})
	go p.Publish("hi")
	go p.Publish("golang: https://golang.org")
	go p.Publish("golang: https://golang.com")
	go p.Publish("docker: https://www.docker.com/")
	time.Sleep(1)
	go func() {
		for item := range golang {
			fmt.Println("golang topic:", item.(string))
		}
	}()
	go func() {
		fmt.Println("docker topic:", <-docker)
	}()
	time.Sleep(2000) //<-make(chan bool)
}
