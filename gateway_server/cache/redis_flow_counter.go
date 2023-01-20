package cache

import (
	"context"
	"gateway_server/global"
	"go.uber.org/zap"
	"strconv"
	"sync/atomic"
	"time"
)

type RedisFlowCounter struct {
	ServiceName string        `json:"service_name"`
	QPS         int           `json:"qps"`
	LastTicket  int64         `json:"last_ticket"`
	AddTicket   int64         `json:"add_ticket"`
	Interval    time.Duration `json:"interval"`
	Unix        int64         `json:"unix"`
}

func NewRedisFlowCounter(serviceName string, interval time.Duration) *RedisFlowCounter {
	counter := &RedisFlowCounter{
		ServiceName: serviceName,
		QPS:         0,
		LastTicket:  0,
		AddTicket:   0,
		Interval:    interval,
		Unix:        0,
	}
	ticker := time.NewTicker(interval)
	go func() {
		for {
			<-ticker.C
			// 得到上一秒加的ticket
			addTicket := atomic.LoadInt64(&counter.AddTicket)
			// 累加流量到redis
			_, err := global.RedisClient.IncrBy(context.Background(), GetDayKey(serviceName, time.Now()), addTicket).Result()
			if err != nil {
				zap.S().Warnf("day flow count error,serviceName:%s", serviceName)
			}
			// 累加流量到redis
			_, err = global.RedisClient.IncrBy(context.Background(), GetHourKey(serviceName, time.Now()), addTicket).Result()
			if err != nil {
				zap.S().Warnf("hour flow count error,serviceName:%s", serviceName)
			}
			// 将流量归0
			atomic.StoreInt64(&counter.AddTicket, 0)
			// 这一轮不计算qps，下一轮计算
			if counter.Unix == 0 {
				counter.Unix = time.Now().Unix()
				lastTotalStr, _ := global.RedisClient.Get(context.Background(), GetDayKey(serviceName, time.Now())).Result()
				lastTotal, _ := strconv.ParseInt(lastTotalStr, 10, 64)
				atomic.StoreInt64(&counter.LastTicket, lastTotal)
				continue
			}
			unix := time.Now().Unix()
			lastTotalStr, _ := global.RedisClient.Get(context.Background(), GetDayKey(serviceName, time.Now())).Result()
			currentTotal, _ := strconv.ParseInt(lastTotalStr, 10, 64)
			changeTicket := currentTotal - counter.LastTicket
			if unix > counter.Unix {
				counter.QPS = int((changeTicket) / (unix - counter.Unix))
				counter.Unix = unix
				atomic.StoreInt64(&counter.LastTicket, currentTotal)
			}
		}
	}()
	return counter
}

func GetHourKey(serviceName string, t time.Time) string {
	now := t.Format("2006-01-02 15")
	return global.KVServiceFlowCount + serviceName + ":" + now
}

func GetDayKey(serviceName string, t time.Time) string {
	now := t.Format("2006-01-02")
	return global.KVServiceFlowCount + serviceName + ":" + now
}

func GetHourData(serviceName string, t time.Time) int32 {
	key := GetHourKey(serviceName, t)
	data, err := global.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return 0
	}
	v, _ := strconv.Atoi(data)
	return int32(v)
}

func GetDayData(serviceName string, t time.Time) int32 {
	key := GetDayKey(serviceName, t)
	data, err := global.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return 0
	}
	v, _ := strconv.Atoi(data)
	return int32(v)
}
