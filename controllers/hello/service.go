package hello

import (
	"context"
	helloV1 "go-mall-api/api/hello/v1"
)

type GreeterService struct {
	helloV1.UnimplementedGreeterServer
}

func NewGreeterService() *GreeterService {
	return &GreeterService{}
}

func (s *GreeterService) SayHello(ctx context.Context, in *helloV1.HelloRequest) (*helloV1.HelloReply, error) {
	return &helloV1.HelloReply{Reply: "go " + in.Name}, nil
}
