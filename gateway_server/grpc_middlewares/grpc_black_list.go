package grpc_middlewares

import (
	"errors"
	"gateway_server/cache/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"strings"
)

func GRPCBlackList(detail *model.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if detail.AccessControl.BlackList != "" && detail.AccessControl.OpenAuth == 1 {
			blackList := strings.Split(detail.AccessControl.BlackList, "\n")
			context, ok := peer.FromContext(ss.Context())
			if !ok {
				return errors.New("获得peer失败")
			}
			split := strings.Split(context.Addr.String(), ":")
			clientIp := ""
			if len(split) == 2 {
				clientIp = split[0]
			}
			var match bool
			for _, ip := range blackList {
				if ip == clientIp {
					match = true
					break
				}
			}
			if match {
				return errors.New("您在黑名单内，请联系管理员从黑名单中移除")
			}
		}
		err := handler(srv, ss)
		if err != nil {
			return err
		}
		return nil
	}
}
