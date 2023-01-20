package initial

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"simple_gateway/global"
)

func InitRedis() {
	conf := global.DebugFullConfig.RedisConfig
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password, // no password set
		DB:       0,             // use default DB
	})
}
