package grpc

import (
	"context"
	"fmt"
	"gin-blog/playerground/grpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"testing"
	"time"
)

func TestClientMain(t *testing.T) {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 客户端调用 服务端的Hello方法
	client := proto.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &proto.String{Value: "lee"})

	if err != nil {
		log.Fatalln(err)

	}
	//fmt.Println(reply.GetValue())
	t.Log(reply.GetValue())

	// 客户端调用 服务端的Channel方法 （grpc stream）
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 向服务端发送数据
	go func() {
		for {
			if err := stream.Send(&proto.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	// main goroutine 接受数据
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}

// 从客户端向服务器发布 主题
func TestPublishClientMain(t *testing.T) {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := proto.NewPubsubServiceClient(conn)

	// 发布
	_, err = client.Publish(
		context.Background(), &proto.String{Value: "golang: hello Go"},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(
		context.Background(), &proto.String{Value: "docker: hello Docker"},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func TestSubscribeClientMain(t *testing.T) {
	waitChan := make(chan struct{})
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := proto.NewPubsubServiceClient(conn)

	// 订阅 主题
	stream, err := client.Subscribe(
		context.Background(), &proto.String{Value: "golang:"},
	)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			reply, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			fmt.Println(reply.GetValue())
		}
	}()

	//
	stream2, err := client.Subscribe(
		context.Background(), &proto.String{Value: "doc"},
	)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			reply, err := stream2.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			fmt.Println(reply.GetValue())
		}
	}()

	<-waitChan //time.Sleep(5000)//
}
