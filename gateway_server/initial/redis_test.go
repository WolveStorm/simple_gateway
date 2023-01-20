package initial

import (
	"context"
	"fmt"
	"gateway_server/global"
	"testing"
)

func TestInitRedis(t *testing.T) {
	InitConfigByPath(true, "D:\\simple_gateway\\gateway_server\\conf\\dev\\base.yaml")
	InitRedis()
	res, _ := global.RedisClient.Ping(context.Background()).Result()
	fmt.Println(res)
}
