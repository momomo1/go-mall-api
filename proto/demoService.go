package proto

import (
	"context"
)

type Server struct {
	UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Reply: "go " + in.Name}, nil
}
