package grpc

import (
	"gin-blog/playerground/grpc/proto"
	"gin-blog/playerground/grpc/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

func TestServerMain(t *testing.T) {
	grpcServer := grpc.NewServer()
	proto.RegisterHelloServiceServer(grpcServer, new(service.HelloServiceImpl))
	proto.RegisterPubsubServiceServer(grpcServer, service.NewPubsubService())
	list, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal(err)

	}

	grpcServer.Serve(list)
}
