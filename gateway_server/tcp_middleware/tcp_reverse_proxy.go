package tcp_middleware

import (
	"context"
	"gateway_server/reverse_proxy/load_balance"
	"gateway_server/tcp/server"
	"gateway_server/tcp_middleware/tin"
	"io"
	"log"
	"net"
	"time"
)

func TCPReverseProxyWithLoadBalance(c *tin.TinContext, lb load_balance.LoadBalance) server.TCPHandler {
	return func() *TcpReverseProxy {
		addr := lb.Get("")
		return &TcpReverseProxy{
			ctx:             c.Ctx,
			Addr:            addr,
			KeepAlivePeriod: time.Second,
			DialTimeout:     time.Second,
		}
	}()
}

//TCP反向代理
type TcpReverseProxy struct {
	ctx             context.Context //单次请求单独设置
	Addr            string
	KeepAlivePeriod time.Duration //设置
	DialTimeout     time.Duration //设置超时时间
	OnDialError     func(src net.Conn, dstDialErr error)
}

func (dp *TcpReverseProxy) onDialError() func(src net.Conn, dstDialErr error) {
	if dp.OnDialError != nil {
		return dp.OnDialError
	}
	return func(src net.Conn, dstDialErr error) {
		log.Printf("tcpproxy: for incoming conn %v, error dialing %q: %v", src.RemoteAddr().String(), dp.Addr, dstDialErr)
		src.Close()
	}
}
func (dp *TcpReverseProxy) keepAlivePeriod() time.Duration {
	if dp.KeepAlivePeriod != 0 {
		return dp.KeepAlivePeriod
	}
	return time.Minute
}
func (dp *TcpReverseProxy) ServeTCP(ctx context.Context, src net.Conn) {
	if dp.DialTimeout > 0 {
		dst, err := net.DialTimeout("tcp", dp.Addr, dp.DialTimeout)
		if err != nil {
			dp.onDialError()(dst, err)
			return
		}
		defer dst.Close()
		// 设置keepAlive
		if ka := dp.keepAlivePeriod(); ka > 0 {
			if c, ok := dst.(*net.TCPConn); ok {
				c.SetKeepAlive(true)
				c.SetKeepAlivePeriod(ka)
			}
		}
		errChan := make(chan error)
		go execProxy(dst, src, errChan)
		go execProxy(src, dst, errChan)
		// 直到出错
		<-errChan
	} else {
		dst, err := net.Dial("tcp", dp.Addr)
		if err != nil {
			dp.onDialError()(dst, err)
			return
		}
		defer dst.Close()
		// 设置keepAlive
		if ka := dp.keepAlivePeriod(); ka > 0 {
			if c, ok := dst.(*net.TCPConn); ok {
				c.SetKeepAlive(true)
				c.SetKeepAlivePeriod(ka)
			}
		}
		errChan := make(chan error)
		go execProxy(dst, src, errChan)
		go execProxy(src, dst, errChan)
		// 直到出错
		<-errChan
	}
}

func execProxy(dst, src net.Conn, errChan chan<- error) {
	_, err := io.Copy(dst, src)
	errChan <- err
}
