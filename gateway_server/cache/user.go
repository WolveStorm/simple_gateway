package cache

import (
	"context"
	"encoding/json"
	"gateway_server/cache/model"
	"gateway_server/global"
)

func GetAllUSer() []*model.User {
	list := make([]*model.User, 0)
	client := global.RedisClient
	result, err := client.HGetAll(context.Background(), global.HashAppInfoKey).Result()
	if err != nil {
		return nil
	}
	for _, v := range result {
		detail := &model.User{}
		err := json.Unmarshal([]byte(v), detail)
		if err != nil {
			continue
		}
		list = append(list, detail)
	}
	return list
}
