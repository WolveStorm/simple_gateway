package initial

import (
	"fmt"
	"gateway_server/global"
	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	conf := global.DebugFullConfig.RedisConfig
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password, // no password set
		DB:       0,             // use default DB
	})
}
