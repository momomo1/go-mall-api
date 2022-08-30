package bootstrap

import (
	"go-mall-api/pkg/config"
	"go-mall-api/pkg/logger"
	"go-mall-api/proto"
	"google.golang.org/grpc"
	"net"
)

func SetGrpc() {
	// 监听本地的端口
	lis, err := net.Listen(config.Get("grpc.link"), config.Get("grpc.port"))
	logger.LogIf(err)
	s := grpc.NewServer()                           // 创建gRPC服务器
	proto.RegisterGreeterServer(s, &proto.Server{}) // 在gRPC服务端注册服务
	// 启动服务
	err = s.Serve(lis)
	logger.LogIf(err)
}
