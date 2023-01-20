package grpc_middlewares

import (
	"errors"
	"gateway_server/cache/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

func GrpcHeaderTransfor(detail *model.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		transfor := detail.GrpcRule.HeaderTransfor
		md, ok := metadata.FromIncomingContext(ss.Context())
		metadataCopy := md.Copy()
		if !ok {
			return errors.New("miss metadata")
		}
		for _, item := range strings.Split(transfor, ",") {
			split := strings.Split(item, " ")
			if len(split) != 3 {
				continue
			}
			if split[0] == "add" || split[0] == "edit" {
				metadataCopy.Set(split[1], split[2])
			} else if split[0] == "del" {
				//因为md是map结构
				delete(metadataCopy, split[1])
			}
		}
		context := metadata.NewIncomingContext(ss.Context(), metadataCopy)

		if err := handler(srv, wrapServerStream(context, ss)); err != nil {
			return err
		}
		return nil
	}
}
