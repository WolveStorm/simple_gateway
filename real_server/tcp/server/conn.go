package server

import (
	"context"
	"net"
	"runtime"
)

type tcpKeepAliveListener struct {
	*net.TCPListener
}

type conn struct {
	server     *TCPServer
	rwc        net.Conn
	remoteAddr string
}

func (c *conn) serve(ctx context.Context) {
	c.remoteAddr = c.rwc.RemoteAddr().String()
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
		}
	}()

	if c.server.Handler == nil {
		panic("handler empty")
	}
	c.server.Handler.ServeTCP(ctx, c.rwc)
}
