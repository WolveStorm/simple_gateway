package load_balance

import (
	"crypto/tls"
	"fmt"
	"gateway_server/cache/model"
	"gateway_server/global"
	"gateway_server/https/ca"
	"net"
	"net/http"
	"strings"
	"time"
)

type LoadBalance interface {
	Add(param ...string) error
	Get(addr string) string
	Update(addrs []string, weightMap map[string]string) // 观察者模式
	SetConf(LoadBalanceCof)
	SetForbid([]string)
	GetForbid() []string
}
type LoadType int

var (
	RandomLoadType LoadType = 1
	RoundLoadType  LoadType = 2
	WeightLoadType LoadType = 3
	HashConsistent LoadType = 4
)

func GenerateLoadBalance(t LoadType) LoadBalance {
	switch t {
	case RandomLoadType:
		r := &RandomLoadBalance{}
		return r
	case RoundLoadType:
		r := &RoundLoadBalance{}
		return r
	case WeightLoadType:
		r := &WeightRoundLoadBalance{}
		return r
	case HashConsistent:
		r := NewHashConsistentLoadBalance(3, nil)
		return r
	default:
		r := &RandomLoadBalance{}
		return r
	}
}

func GetLoadBalanceByService(service *model.ServiceDetail) LoadBalance {
	balance := service.LoadBalance
	loadBalance := GenerateLoadBalance(LoadType(balance.RoundType))
	loadBalance.SetForbid(strings.Split(balance.ForbidList, ","))
	weightMap := make(map[string]string)
	ipList := strings.Split(service.LoadBalance.IpList, ",")
	weightList := strings.Split(service.LoadBalance.WeightList, ",")
	for index, item := range ipList {
		weightMap[item] = weightList[index]
	}
	// 判断协议
	scheme := ""
	if service.ServiceInfo.LoadType == global.HTTPMode && service.HttpRule.NeedHttps == 1 {
		scheme = "https"
		if service.HttpRule.NeedWebsocket == 1 {
			scheme = "wss"
		}
	} else if service.ServiceInfo.LoadType == global.HTTPMode && service.HttpRule.NeedHttps == 0 {
		scheme = "http"
		if service.HttpRule.NeedWebsocket == 1 {
			scheme = "ws"
		}
	} else if service.ServiceInfo.LoadType == global.TCPMode || service.ServiceInfo.LoadType == global.GRPcMode {
		scheme = ""
	}
	var consulConf *ConsulConf
	if service.ServiceInfo.LoadType == global.HTTPMode {
		consulConf = NewConsulConf(loadBalance, fmt.Sprintf("%s://%s", scheme, "%s"), "real_server", global.DebugFullConfig.ConsulConfig.Addr, weightMap)
	} else {
		consulConf = NewConsulConf(loadBalance, fmt.Sprintf("%s%s", scheme, "%s"), "real_server", global.DebugFullConfig.ConsulConfig.Addr, weightMap)
	}
	loadBalance.SetConf(consulConf)
	consulConf.Attach(loadBalance)
	return loadBalance
}

func GetTransport(service *model.ServiceDetail) *http.Transport {
	if service.LoadBalance.UpstreamConnectTimeout == 0 {
		service.LoadBalance.UpstreamConnectTimeout = 30
	}
	if service.LoadBalance.UpstreamMaxIdle == 0 {
		service.LoadBalance.UpstreamMaxIdle = 100
	}
	if service.LoadBalance.UpstreamIdleTimeout == 0 {
		service.LoadBalance.UpstreamIdleTimeout = 90
	}
	if service.LoadBalance.UpstreamHeaderTimeout == 0 {
		service.LoadBalance.UpstreamHeaderTimeout = 30
	}
	config := &tls.Config{}
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], _ = tls.LoadX509KeyPair(ca.Path("server.crt"), ca.Path("server.key"))
	config.InsecureSkipVerify = true
	trans := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(service.LoadBalance.UpstreamConnectTimeout),
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          service.LoadBalance.UpstreamMaxIdle,
		IdleConnTimeout:       time.Duration(service.LoadBalance.UpstreamIdleTimeout),
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ResponseHeaderTimeout: time.Duration(service.LoadBalance.UpstreamHeaderTimeout) * time.Second,
		TLSClientConfig:       config,
	}
	return trans
}
