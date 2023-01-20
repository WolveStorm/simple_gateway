package server

import (
	"context"
	"net"
	"testing"
)

type TcpHandler struct {
}

func (*TcpHandler) ServeTCP(ctx context.Context, conn net.Conn) {
	conn.Write([]byte("hello,this is tcp"))
}

// telnet 127.0.0.1 7777
func TestListenAndServe(t *testing.T) {
	ListenAndServe(":7777", &TcpHandler{}, context.Background())
}
