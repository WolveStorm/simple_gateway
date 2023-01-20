package server

import (
	"context"
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type TCPHandler interface {
	ServeTCP(ctx context.Context, conn net.Conn)
}

type TCPServer struct {
	// Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	// See net.Dial for details of the address format.
	Addr string

	Handler TCPHandler // handler to invoke, http.DefaultServeMux if nil

	ReadTimeout time.Duration

	ReadHeaderTimeout time.Duration

	WriteTimeout time.Duration

	IdleTimeout time.Duration

	MaxHeaderBytes int

	inShutdown int64 // true when server is in shutdown\

	mu sync.Mutex

	doneChan chan struct{}

	l *onceCloseListener

	BaseCtx context.Context
}

func ListenAndServe(addr string, handler TCPHandler, ctx context.Context) (*TCPServer, error) {
	server := &TCPServer{Addr: addr, Handler: handler, BaseCtx: ctx}
	err := server.ListenAndServe()
	return server, err
}

func (srv *TCPServer) ListenAndServe() error {
	if srv.shuttingDown() {
		return errors.New("已经关闭服务器")
	}
	addr := srv.Addr
	if addr == "" {
		return errors.New("无addr")
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(&tcpKeepAliveListener{ln.(*net.TCPListener)})
}
func (srv *TCPServer) Serve(l net.Listener) error {
	ln := &onceCloseListener{Listener: l}
	srv.l = ln
	if srv.BaseCtx == nil {
		srv.BaseCtx = context.Background()
	}
	defer ln.close()
	for {
		rw, err := l.Accept()
		if err != nil {
			select {
			case <-srv.getDoneChan():
				return errors.New("close")
			default:
			}
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				continue
			}
			return err
		}
		c := srv.newConn(rw)
		go c.serve(srv.BaseCtx)
	}
}

func (srv *TCPServer) newConn(con net.Conn) *conn {
	c := &conn{
		server:     srv,
		rwc:        con,
		remoteAddr: "",
	}
	return c
}

func (srv *TCPServer) shuttingDown() bool {
	return srv.inShutdown == 1
}

func (srv *TCPServer) Close() error {
	atomic.StoreInt64(&srv.inShutdown, 1)
	srv.mu.Lock()
	defer srv.mu.Unlock()
	srv.closeDoneChanLocked()
	err := srv.closeListenersLocked()
	return err
}

func (s *TCPServer) getDoneChan() <-chan struct{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.getDoneChanLocked()
}

func (s *TCPServer) closeDoneChanLocked() {
	ch := s.getDoneChanLocked()
	select {
	case <-ch:
		// Already closed. Don't close again.
	default:
		// Safe to close here. We're the only closer, guarded
		// by s.mu.
		close(ch)
	}
}

func (s *TCPServer) getDoneChanLocked() chan struct{} {
	if s.doneChan == nil {
		s.doneChan = make(chan struct{})
	}
	return s.doneChan
}

func (s *TCPServer) closeListenersLocked() error {
	return s.l.Listener.Close()
}

type onceCloseListener struct {
	net.Listener
	once     sync.Once
	closeErr error
}

func (oc *onceCloseListener) Close() error {
	oc.once.Do(oc.close)
	return oc.closeErr
}

func (oc *onceCloseListener) close() { oc.closeErr = oc.Listener.Close() }
