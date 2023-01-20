package model

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"simple_gateway/global"
	"sync"
)

var (
	// hget hash:gateway:service_info test_grpc_service
	HashServiceInfoKey = "hash:gateway:service_info"
	// hget hash:gateway:app_info tianjia
	HashAppInfoKey = "hash:gateway:app_info"
)

type RedisSyncClient struct {
	client *redis.Client
	once   sync.Once
	mutex  sync.Mutex
}

func NewRedisSyncClient(client *redis.Client) *RedisSyncClient {
	return &RedisSyncClient{
		client: client,
		once:   sync.Once{},
		mutex:  sync.Mutex{},
	}
}

func SyncToRedis() {
	sClient := NewRedisSyncClient(global.RedisClient)
	// 只做一次，系统启动的时候，替换系统当前的配置到redis
	sClient.once.Do(func() {
		redisClient := sClient.client
		user := &User{}
		all, _ := user.SelectAll()
		userMap := make(map[string]string)
		sClient.mutex.Lock()
		for _, v := range all {
			userInfo, _ := json.Marshal(v)
			userMap[v.AppID] = string(userInfo)
		}
		sClient.mutex.Unlock()
		_, err := redisClient.HMSet(context.Background(), HashAppInfoKey, userMap).Result()
		if err != nil {
			zap.S().Error("app信息载入到内存出错！")
		}
		info := &ServiceInfo{}
		services, _ := info.SelectAll()
		details := make([]*ServiceDetail, 0)
		for _, v := range services {
			serviceDetail, err := GetServiceDetail(v)
			if err != nil {
				zap.S().Errorf("serviceDetail查询失败，id:", v.ID)
				continue
			}
			details = append(details, serviceDetail)
		}
		serviceMap := make(map[string]string)
		sClient.mutex.Lock()
		for _, v := range details {
			serviceInfo, _ := json.Marshal(v)
			serviceMap[v.ServiceInfo.ServiceName] = string(serviceInfo)
		}
		sClient.mutex.Unlock()
		_, err = redisClient.HMSet(context.Background(), HashServiceInfoKey, serviceMap).Result()
		if err != nil {
			zap.S().Error("service信息载入到内存出错！")
		}
	})
}

func ServiceSync(detail ServiceDetail) {
	sClient := NewRedisSyncClient(global.RedisClient)
	redisClient := sClient.client
	serviceInfo, _ := json.Marshal(detail)
	_, err := redisClient.HSet(context.Background(), HashServiceInfoKey, detail.ServiceInfo.ServiceName, string(serviceInfo)).Result()
	if err != nil {
		zap.S().Errorf("sync service error,id:%s", detail.Id)
	}
}

func DeleteServiceSync(serviceName string) {
	sClient := NewRedisSyncClient(global.RedisClient)
	redisClient := sClient.client
	_, err := redisClient.HDel(context.Background(), HashServiceInfoKey, serviceName).Result()
	if err != nil {
		zap.S().Errorf("delete sync service error,serviceName:%s", serviceName)
	}
}

func UserSync(app User) {
	sClient := NewRedisSyncClient(global.RedisClient)
	redisClient := sClient.client
	serviceInfo, _ := json.Marshal(app)
	_, err := redisClient.HSet(context.Background(), HashAppInfoKey, app.AppID, string(serviceInfo)).Result()
	if err != nil {
		zap.S().Errorf("sync user error,id:%s", app.AppID)
	}
}

func DeleteUserSync(appId string) {
	sClient := NewRedisSyncClient(global.RedisClient)
	redisClient := sClient.client
	_, err := redisClient.HDel(context.Background(), HashAppInfoKey, appId).Result()
	if err != nil {
		zap.S().Errorf("delete sync service error,appId:%s", appId)
	}
}
