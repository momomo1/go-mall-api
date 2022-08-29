package serviceGrpc

import (
	"context"
	"go-mall-api/proto"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Reply: "go " + in.Name}, nil
}
