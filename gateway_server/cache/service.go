package cache

import (
	"context"
	"encoding/json"
	"gateway_server/cache/model"
	"gateway_server/global"
)

func GetAllService() []*model.ServiceDetail {
	list := make([]*model.ServiceDetail, 0)
	client := global.RedisClient
	result, err := client.HGetAll(context.Background(), global.HashServiceInfoKey).Result()
	if err != nil {
		return nil
	}
	for _, v := range result {
		detail := &model.ServiceDetail{}
		err := json.Unmarshal([]byte(v), detail)
		if err != nil {
			continue
		}
		list = append(list, detail)
	}
	return list
}

//func GetWeightMap() map[string]string {
//	m := make(map[string]string)
//	client := global.RedisClient
//	result, err := client.HGetAll(context.Background(), global.HashServiceInfoKey).Result()
//	if err != nil {
//		return nil
//	}
//	for _, v := range result {
//		detail := &model.ServiceDetail{}
//		err := json.Unmarshal([]byte(v), detail)
//		if err != nil {
//			continue
//		}
//		ipList := strings.Split(detail.LoadBalance.IpList, ",")
//		weightList := strings.Split(detail.LoadBalance.WeightList, ",")
//		for i := 0; i < len(strings.Split(detail.LoadBalance.IpList, ",")); i++ {
//			m[ipList[i]] = weightList[i]
//		}
//	}
//	return m
//}

func GetAllHTTPService() []*model.ServiceDetail {
	list := make([]*model.ServiceDetail, 0)
	client := global.RedisClient
	result, err := client.HGetAll(context.Background(), global.HashServiceInfoKey).Result()
	if err != nil {
		return nil
	}
	for _, v := range result {
		detail := &model.ServiceDetail{}
		err := json.Unmarshal([]byte(v), detail)
		if err != nil {
			continue
		}
		if detail.ServiceInfo.LoadType != global.HTTPMode {
			continue
		}
		list = append(list, detail)
	}
	return list
}

func GetAllTCPService() []*model.ServiceDetail {
	list := make([]*model.ServiceDetail, 0)
	client := global.RedisClient
	result, err := client.HGetAll(context.Background(), global.HashServiceInfoKey).Result()
	if err != nil {
		return nil
	}
	for _, v := range result {
		detail := &model.ServiceDetail{}
		err := json.Unmarshal([]byte(v), detail)
		if err != nil {
			continue
		}
		if detail.ServiceInfo.LoadType != global.TCPMode {
			continue
		}
		list = append(list, detail)
	}
	return list
}

func GetAllGRPCService() []*model.ServiceDetail {
	list := make([]*model.ServiceDetail, 0)
	client := global.RedisClient
	result, err := client.HGetAll(context.Background(), global.HashServiceInfoKey).Result()
	if err != nil {
		return nil
	}
	for _, v := range result {
		detail := &model.ServiceDetail{}
		err := json.Unmarshal([]byte(v), detail)
		if err != nil {
			continue
		}
		if detail.ServiceInfo.LoadType != global.GRPcMode {
			continue
		}
		list = append(list, detail)
	}
	return list
}
