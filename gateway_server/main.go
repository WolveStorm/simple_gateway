package main

import (
	"gateway_server/cache/grpc_impl"
	"gateway_server/global"
	"gateway_server/grpc"
	"gateway_server/http"
	"gateway_server/https"
	"gateway_server/initial"
	"gateway_server/protoc"
	"gateway_server/tcp"
	"go.uber.org/zap"
	Grpc "google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initial.InitConfig(true)
	initial.InitAllZap()
	initial.InitRedis()
	go func() {
		http.StartHTTPServer()
	}()
	go func() {
		https.StartHTTPSServer()
	}()
	go func() {
		tcp.StartTCPServer()
	}()
	go func() {
		grpc.StartGRPCServer()
	}()
	go func() {
		// for flow count
		StartGrpcService(&grpc_impl.FlowCountService{})
	}()
	sign := make(chan os.Signal)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	<-sign
	defer func() {
		http.CloseHTTPServer()
		https.CloseHTTPSServer()
		tcp.CloseTcpServer()
		grpc.CloseGrpcServer()
		CloseGrpcService()
	}()
}

var server *Grpc.Server

func StartGrpcService(srv protoc.FlowCountServer) {
	// 开放grpc端口供管理服务器查询流量
	server = Grpc.NewServer()
	listen, _ := net.Listen("tcp", global.DebugFullConfig.GRPCServer.Host)
	RegisterService(srv)
	err := server.Serve(listen)
	if err != nil {
		zap.S().Errorf("[Navi Gateway] grpc service start error,host:%s", global.DebugFullConfig.GRPCServer.Host)
		return
	}
}

func RegisterService(srv protoc.FlowCountServer) {
	protoc.RegisterFlowCountServer(server, srv)
}

func CloseGrpcService() {
	server.Stop()
}
