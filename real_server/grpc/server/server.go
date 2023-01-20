package main

import (
	"context"
	"fmt"
	capi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"net"
	"net/http"
	"os"
	"real_server/grpc/proto"
	"strconv"
	"strings"
)

type Server struct{}

func (s *Server) UnaryEcho(ctx context.Context, in *proto.EchoRequest) (*proto.EchoResponse, error) {
	fmt.Println("----- into UnaryEcho -----")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("miss metadata from context")
	}
	fmt.Printf("metadata is %v\n", md)
	fmt.Printf("recieve from request %v\n,sending echo", in)
	return &proto.EchoResponse{Message: in.Message}, nil
}

func (s *Server) ServerStreamingEcho(in *proto.EchoRequest, stream proto.Echo_ServerStreamingEchoServer) error {
	fmt.Println("----- into ServerStreamingEcho -----\n")
	fmt.Printf("request received: %v\n", in)
	for i := 0; i < 10; i++ {
		fmt.Printf("echo message %v\n", in.Message)
		err := stream.Send(&proto.EchoResponse{Message: in.Message})
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *Server) ClientStreamingEcho(stream proto.Echo_ClientStreamingEchoServer) error {
	fmt.Println("----- into ClientStreamingEcho -----\n")
	var message string
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("echo last received message exit!")
			return stream.SendAndClose(&proto.EchoResponse{Message: message})
		}
		message = recv.Message
		fmt.Printf("request recived:%v\n", message)
		if err != nil {
			return nil
		}
	}
}
func (s *Server) BidirectionalStreamingEcho(stream proto.Echo_BidirectionalStreamingEchoServer) error {
	fmt.Printf("--- BidirectionalStreamingEcho ---\n")
	// Read requests and send responses.
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("request received %v, sending echo\n", in)
		if err := stream.Send(&proto.EchoResponse{Message: in.Message}); err != nil {
			return err
		}
	}
}

func main() {
	reg := &capi.AgentServiceRegistration{
		Name:    "real_server",
		ID:      "real_server",
		Tags:    []string{"real_server"},
		Address: "127.0.0.1:7777",
	}
	client, _ := capi.NewClient(defaultConfig(nil, cleanhttp.DefaultPooledTransport, "127.0.0.1:8500"))
	agent := client.Agent()
	if err := agent.ServiceRegister(reg); err != nil {
		fmt.Println(err)
	}
	listen, err := net.Listen("tcp", ":7777")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen at port :7777")
	server := grpc.NewServer()
	//调用pb.go的注册方法
	proto.RegisterEchoServer(server, &Server{})
	server.Serve(listen)
}
func defaultConfig(logger hclog.Logger, transportFn func() *http.Transport, addr string) *capi.Config {
	if logger == nil {
		logger = hclog.New(&hclog.LoggerOptions{
			Name: "consul-api",
		})
	}

	config := &capi.Config{
		Address:   addr,
		Scheme:    "http",
		Transport: transportFn(),
	}

	if addr := os.Getenv(capi.HTTPAddrEnvName); addr != "" {
		config.Address = addr
	}

	if tokenFile := os.Getenv(capi.HTTPTokenFileEnvName); tokenFile != "" {
		config.TokenFile = tokenFile
	}

	if token := os.Getenv(capi.HTTPTokenEnvName); token != "" {
		config.Token = token
	}

	if auth := os.Getenv(capi.HTTPAuthEnvName); auth != "" {
		var username, password string
		if strings.Contains(auth, ":") {
			split := strings.SplitN(auth, ":", 2)
			username = split[0]
			password = split[1]
		} else {
			username = auth
		}

		config.HttpAuth = &capi.HttpBasicAuth{
			Username: username,
			Password: password,
		}
	}

	if ssl := os.Getenv(capi.HTTPSSLEnvName); ssl != "" {
		enabled, err := strconv.ParseBool(ssl)
		if err != nil {
			logger.Warn(fmt.Sprintf("could not parse %s", capi.HTTPSSLEnvName), "error", err)
		}

		if enabled {
			config.Scheme = "https"
		}
	}

	if v := os.Getenv(capi.HTTPTLSServerName); v != "" {
		config.TLSConfig.Address = v
	}
	if v := os.Getenv(capi.HTTPCAFile); v != "" {
		config.TLSConfig.CAFile = v
	}
	if v := os.Getenv(capi.HTTPCAPath); v != "" {
		config.TLSConfig.CAPath = v
	}
	if v := os.Getenv(capi.HTTPClientCert); v != "" {
		config.TLSConfig.CertFile = v
	}
	if v := os.Getenv(capi.HTTPClientKey); v != "" {
		config.TLSConfig.KeyFile = v
	}
	if v := os.Getenv(capi.HTTPSSLVerifyEnvName); v != "" {
		doVerify, err := strconv.ParseBool(v)
		if err != nil {
			logger.Warn(fmt.Sprintf("could not parse %s", capi.HTTPSSLVerifyEnvName), "error", err)
		}
		if !doVerify {
			config.TLSConfig.InsecureSkipVerify = true
		}
	}

	if v := os.Getenv(capi.HTTPNamespaceEnvName); v != "" {
		config.Namespace = v
	}

	if v := os.Getenv(capi.HTTPPartitionEnvName); v != "" {
		config.Partition = v
	}

	return config
}
