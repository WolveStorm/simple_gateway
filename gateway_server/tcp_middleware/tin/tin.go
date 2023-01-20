package tin

import (
	"context"
	"gateway_server/tcp/server"
	"math"
	"net"
)

var abortIndex int8 = math.MaxInt8 / 2 //最多 63 个中间件

// 仿造gin的洋葱圈结构，打造一个轻量的TCP中间件
// Tin
type TinHandler func(c *TinContext)

type TinRouter struct {
	groups []*TinRouterGroup
}

// 构造 router
func NewTinRouter() *TinRouter {
	return &TinRouter{}
}

type TinRouterGroup struct {
	path     string
	handlers []TinHandler
	*TinRouter
}

type TinContext struct {
	Conn  net.Conn
	Ctx   context.Context
	index int8
	*TinRouterGroup
}

type TinRouterHandler struct {
	coreFunc func(tinContext *TinContext) server.TCPHandler
	router   *TinRouter
}

func NewTinSliceRouterHandler(coreFunc func(tinContext *TinContext) server.TCPHandler, router *TinRouter) *TinRouterHandler {
	return &TinRouterHandler{
		coreFunc: coreFunc,
		router:   router,
	}
}

func (c *TinContext) Get(key interface{}) interface{} {
	return c.Ctx.Value(key)
}

func (c *TinContext) Set(key, val interface{}) {
	c.Ctx = context.WithValue(c.Ctx, key, val)
}

func newTcpSliceRouterContext(conn net.Conn, r *TinRouter, ctx context.Context) *TinContext {
	newTcpSliceGroup := &TinRouterGroup{}
	*newTcpSliceGroup = *r.groups[0] //浅拷贝数组指针,只会使用第一个分组
	c := &TinContext{Conn: conn, TinRouterGroup: newTcpSliceGroup, Ctx: ctx}
	c.Reset()
	return c
}
func (w *TinRouterHandler) ServeTCP(ctx context.Context, conn net.Conn) {
	c := newTcpSliceRouterContext(conn, w.router, ctx)
	c.handlers = append(c.handlers, func(c *TinContext) {
		w.coreFunc(c).ServeTCP(ctx, conn)
	})
	c.Reset()
	c.Next()
}

// 创建 Group
func (g *TinRouter) Group(path string) *TinRouterGroup {
	if path != "/" {
		panic("only accept path=/")
	}
	return &TinRouterGroup{
		TinRouter: g,
		path:      path,
	}
}

// 构造回调方法
func (g *TinRouterGroup) Use(middlewares ...TinHandler) *TinRouterGroup {
	g.handlers = append(g.handlers, middlewares...)
	existsFlag := false
	for _, oldGroup := range g.TinRouter.groups {
		if oldGroup == g {
			existsFlag = true
		}
	}
	if !existsFlag {
		g.TinRouter.groups = append(g.TinRouter.groups, g)
	}
	return g
}

// 从最先加入中间件开始回调
func (c *TinContext) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// 是否跳过了回调
func (c *TinContext) IsAborted() bool {
	return c.index >= abortIndex
}

func (c *TinContext) Abort() {
	c.index = abortIndex
}

// 重置回调
func (c *TinContext) Reset() {
	c.index = -1
}
