package cache

import (
	"golang.org/x/time/rate"
	"sync"
)

var FlowLimitManager *FlowLimiter

// 将计时器存入内存，读取流量的时候更加迅速。
type FlowLimiter struct {
	RedisFlowMap   map[string]*Limiter
	RedisFlowSlice []*Limiter
	mutex          sync.Mutex
}

type Limiter struct {
	Limit *rate.Limiter
}

func init() {
	FlowLimitManager = NewFlowLimiterManager()
}
func NewFlowLimiterManager() *FlowLimiter {
	return &FlowLimiter{
		RedisFlowMap:   make(map[string]*Limiter),
		RedisFlowSlice: make([]*Limiter, 0),
		mutex:          sync.Mutex{},
	}
}

// 通过一秒拿到多少张令牌的数量来创建限流器。
func (l *FlowLimiter) NewLimiter(serviceName string, qps int) *Limiter {
	if limiter, ok := l.RedisFlowMap[serviceName]; ok {
		return limiter
	} else {
		limiter := rate.NewLimiter(rate.Limit(qps), 3*qps)
		l1 := &Limiter{Limit: limiter}
		l.RedisFlowSlice = append(l.RedisFlowSlice, l1)
		l.mutex.Lock()
		l.RedisFlowMap[serviceName] = l1
		l.mutex.Unlock()
		return l1
	}
}
