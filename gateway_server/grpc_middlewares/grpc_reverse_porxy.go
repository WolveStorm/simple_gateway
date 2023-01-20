package grpc_middlewares

import (
	"context"
	"gateway_server/reverse_proxy/load_balance"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func GRPCReverseProxyWithLoadBalance(lb load_balance.LoadBalance) func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
	return func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		outCtx := metadata.NewOutgoingContext(ctx, md.Copy())
		addr := lb.Get("")
		cc, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			return outCtx, nil, err
		}
		return outCtx, cc, nil
	}
}
