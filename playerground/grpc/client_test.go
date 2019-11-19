package grpc

import (
	"context"
	"gin-blog/playerground/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestClientMain(t *testing.T) {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &proto.String{Value: "lee"})

	if err != nil {
		log.Fatalln(err)

	}

	//fmt.Println(reply.GetValue())
	t.Log(reply.GetValue())
}
