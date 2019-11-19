package service

import (
	"context"
	"gin-blog/playerground/grpc/proto"
)

// 基于服务端的HelloServiceServer接口可以重新实现HelloService服务：
// HelloServiceServer就是
type HelloServiceImpl struct {
}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *proto.String) (*proto.String, error) {
	reply := &proto.String{Value: "hello" + args.GetValue()}
	return reply, nil
}
