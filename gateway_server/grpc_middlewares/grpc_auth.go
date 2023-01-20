package grpc_middlewares

import (
	"context"
	"errors"
	"gateway_server/cache"
	"gateway_server/cache/model"
	"gateway_server/util"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
	"strings"
)

var AuthKey = "authorization"

type serverStream struct {
	grpc.ServerStream
	ctx context.Context

	receivedMessageID int
	sentMessageID     int
}

func (w *serverStream) Context() context.Context {
	return w.ctx
}

func (w *serverStream) RecvMsg(m interface{}) error {
	err := w.ServerStream.RecvMsg(m)

	if err == nil {
		w.receivedMessageID++
	}

	return err
}

func (w *serverStream) SendMsg(m interface{}) error {
	err := w.ServerStream.SendMsg(m)

	w.sentMessageID++

	return err
}
func wrapServerStream(ctx context.Context, ss grpc.ServerStream) *serverStream {
	return &serverStream{
		ServerStream: ss,
		ctx:          ctx,
	}
}

func GRPCJWTAuth(detail *model.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		var match bool
		var ctx context.Context
		md, ok := metadata.FromIncomingContext(ss.Context())
		if !ok {
			return errors.New("请先登录 1")
		}
		metadataCopy := md.Copy()
		arr := md.Get(AuthKey)
		var auth string
		if len(arr) > 0 {
			auth = arr[0]
		}
		zap.L().Info(auth)
		if auth != "" {
			split := strings.Split(auth, " ")
			if len(split) != 2 {
				return errors.New("请先登录  2")
			}
			token := split[1]
			if token != "" {
				users := cache.GetAllUSer()
				claim, err := util.VerifyToken(token)
				if err != nil {
					return errors.New("请先登录 3")
				}
				for _, u := range users {
					if u.AppID == claim.AppId {
						match = true
						metadataCopy.Set("userid", string(u.AppID))
						metadataCopy.Set("qpd", strconv.FormatInt(u.Qpd, 10))
						metadataCopy.Set("qps", strconv.FormatInt(u.Qps, 10))
						ctx = metadata.NewIncomingContext(ss.Context(), metadataCopy)
						break
					}
				}
			}
		}
		// 如果网关开启了验证，且用户并没有输入token就需要进行拦截
		if detail.AccessControl.OpenAuth == 1 && !match {
			return errors.New("请先登录 4")
		}
		_ = handler(srv, wrapServerStream(ctx, ss))
		return nil
	}
}
