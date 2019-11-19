package service

import (
	"context"
	"fmt"
	"gin-blog/playerground/grpc/proto"
	"github.com/moby/moby/pkg/pubsub"
	"strings"
	"time"
)

//基于gRPC简单实现了一个跨网络的发布和订阅服务。

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{pub: pubsub.NewPublisher(100*time.Millisecond, 10)}
}

func (p *PubsubService) Publish(ctx context.Context, req *proto.String) (*proto.String, error) {
	fmt.Println("publish:", req, p.pub)
	p.pub.Publish(req.GetValue())
	return &proto.String{}, nil
}

func (p *PubsubService) Subscribe(req *proto.String, stream proto.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, req.GetValue()) {
				return true
			}
		}
		return false
	})
	for v := range ch {
		if err := stream.Send(&proto.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}
