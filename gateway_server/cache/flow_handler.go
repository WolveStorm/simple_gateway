package cache

import (
	"sync"
	"time"
)

var FlowManager *RedisFlowManager

// 将计时器存入内存，读取流量的时候更加迅速。
type RedisFlowManager struct {
	RedisFlowMap   map[string]*RedisFlowCounter
	RedisFlowSlice []*RedisFlowCounter
	mutex          sync.Mutex
}

func init() {
	FlowManager = NewRedisFlowManager()
}
func NewRedisFlowManager() *RedisFlowManager {
	return &RedisFlowManager{
		RedisFlowMap:   make(map[string]*RedisFlowCounter),
		RedisFlowSlice: make([]*RedisFlowCounter, 0),
		mutex:          sync.Mutex{},
	}
}

func (r *RedisFlowManager) GetFlowCounter(serviceName string) *RedisFlowCounter {
	if counter, ok := r.RedisFlowMap[serviceName]; ok {
		return counter
	} else {
		flowCounter := NewRedisFlowCounter(serviceName, 1*time.Second)
		r.RedisFlowSlice = append(r.RedisFlowSlice, flowCounter)
		r.mutex.Lock()
		r.RedisFlowMap[serviceName] = flowCounter
		r.mutex.Unlock()
		return flowCounter
	}
}
