package service

import (
	"context"
	"gin-blog/playerground/grpc/proto"
	"io"
)

// 基于服务端的HelloServiceServer接口可以重新实现HelloService服务：
// HelloServiceServer就是hello.pb.go的接口
type HelloServiceImpl struct {
}

// 该方法定义可以看hello.pb.go的UnimplementedHelloServiceServer结构体默认实现HelloServiceServer接口的方法定义
func (p *HelloServiceImpl) Hello(ctx context.Context, args *proto.String) (*proto.String, error) {
	reply := &proto.String{Value: "hello" + args.GetValue()}
	return reply, nil
}

/*
https://tutorial.e-learn.cn/read/advanced-go-programming-book/ch4-rpc-ch4-04-grpc.md
RPC是远程函数调用，因此每次调用的函数参数和返回值不能太大，否则将严重影响每次调用的响应时间。
因此传统的RPC方法调用对于上传和下载较大数据量场景并不适合。
同时传统RPC模式也不适用于对时间不确定的订阅和发布模式。为此，gRPC框架针对服务器端和客户端分别提供了流特性。
*/

//关键字stream指定启用流特性，参数部分是接收客户端参数的流，返回值是返回给客户端的流。
//在服务端的Channel方法参数是一个新的HelloService_ChannelServer类型的参数，可以用于和客户端双向通信。
// 客户端的Channel方法返回一个HelloService_ChannelClient类型的返回值，可以用于和服务端进行双向通信。
func (p *HelloServiceImpl) Channel(srv proto.HelloService_ChannelServer) error {
	for {
		args, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &proto.String{Value: "hello:" + args.GetValue()}

		err = srv.Send(reply)

	}
}
