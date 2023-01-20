package grpc_impl

import (
	"context"
	"fmt"
	"gateway_server/global"
	"gateway_server/protoc"
	"google.golang.org/grpc"
	"testing"
)

func TestGrpc(t *testing.T) {
	dial, err := grpc.Dial(":8848", grpc.WithInsecure())
	t.Run("for service", func(t *testing.T) {
		if err != nil {
			panic(err)
		}
		client := protoc.NewFlowCountClient(dial)
		count, err := client.GetServiceFlowCount(context.Background(), &protoc.FlowCountRequest{ServiceName: "test_http_service"})
		if err != nil {
			panic(err)
		}
		fmt.Println(count)
	})
	t.Run("for all", func(t *testing.T) {
		if err != nil {
			panic(err)
		}
		client := protoc.NewFlowCountClient(dial)
		count, err := client.GetServiceFlowCount(context.Background(), &protoc.FlowCountRequest{ServiceName: global.TotalKey})
		if err != nil {
			panic(err)
		}
		fmt.Println(count)
	})
	t.Run("for user", func(t *testing.T) {
		if err != nil {
			panic(err)
		}
		client := protoc.NewFlowCountClient(dial)
		count, err := client.GetUserFlowCount(context.Background(), &protoc.FlowCountRequest{ServiceName: "8d7b11ec9be0e59a36b52f32366c09cc"})
		if err != nil {
			panic(err)
		}
		fmt.Println(count)
	})
}
