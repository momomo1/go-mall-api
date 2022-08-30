package proto

import (
	"context"
	"go-mall-api/api/hello/v1"
)

type Server struct {
	v1.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloResponse, error) {
	return &v1.HelloResponse{Reply: "go " + in.Name}, nil
}
