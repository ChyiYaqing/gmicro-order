package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/chyiyaqing/gmicro-order/config"
	"github.com/chyiyaqing/gmicro-order/internal/ports"
	"github.com/chyiyaqing/gmicro-proto/golang/order"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api                            ports.APIPort // API 接口
	port                           int           // 服务端端口
	server                         *grpc.Server  // gRPC服务器实例
	order.UnimplementedOrderServer               // 未实现的订单服务器基类
}

func NewAdaptor(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a *Adapter) Run() {
	var err error
	// 1. 创建 TCP 监听器
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	// 2. 创建 gRPC 服务器，配置OpenTelemetry拦截器用于监控
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	// 3. 保存服务器实例
	a.server = grpcServer

	// 4. 注册订单服务
	order.RegisterOrderServer(grpcServer, a)

	// 5. 在开发环境下启动gRPC反射服务 -- 便于调试
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	// 6. 启动服务
	log.Printf("starting order service on port %d ...", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}
}

func (a *Adapter) Stop() {
	a.server.Stop()
}
