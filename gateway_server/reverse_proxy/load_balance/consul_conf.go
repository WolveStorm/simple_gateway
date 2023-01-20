package load_balance

import (
	"fmt"
	capi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 负载均衡的发现中心通用接口
// 该go文件实现consul
type LoadBalanceCof interface {
	Attach(observer LoadBalance)
	GetActiveIPS() []string
	UpdateIPS([]string)
}

type Observer interface {
	Update(addrs []string, weightMap map[string]string)
}

type ConsulConf struct {
	lb           LoadBalance
	ServiceName  string   // 服务名
	ConsulHost   string   // consul的地址
	Format       string   // 根据pattern返回对应的实例地址
	ActiveServer []string // 活跃的服务器
	WeightMap    map[string]string
}

func NewConsulConf(lb LoadBalance, format, serviceName, consulHost string, weightMap map[string]string) *ConsulConf {
	return &ConsulConf{
		lb:           lb,
		ServiceName:  serviceName,
		ConsulHost:   consulHost,
		Format:       format,
		ActiveServer: make([]string, 0),
		WeightMap:    weightMap,
	}
}

func (conf *ConsulConf) Attach(observer LoadBalance) {
	conf.lb = observer
}
func inArray(arr []string, tar string) bool {
	for _, v := range arr {
		if v == tar {
			return true
		}
	}
	return false
}

// 从consul获得活跃ip
func (conf *ConsulConf) GetActiveIPS() []string {
	client, err := capi.NewClient(defaultConfig(nil, cleanhttp.DefaultPooledTransport, conf.ConsulHost))
	if err != nil {
		return []string{}
	}
	agent := client.Agent()
	services, err := agent.ServicesWithFilter("Service == " + conf.ServiceName)
	if err != nil {
		return []string{}
	}
	ips := make([]string, 0)
	forbidIps := conf.lb.GetForbid()
	for _, service := range services {
		if inArray(forbidIps, service.Address) {
			continue
		}
		if conf.Format != "" {
			ips = append(ips, fmt.Sprintf(conf.Format, service.Address))
		} else {
			ips = append(ips, service.Address)
		}
	}
	conf.UpdateIPS(ips)
	return ips
}

func (conf *ConsulConf) UpdateIPS(addrs []string) {
	conf.ActiveServer = addrs
	conf.lb.Update(addrs, conf.WeightMap)
}

func defaultConfig(logger hclog.Logger, transportFn func() *http.Transport, addr string) *capi.Config {
	if logger == nil {
		logger = hclog.New(&hclog.LoggerOptions{
			Name: "consul-api",
		})
	}

	config := &capi.Config{
		Address:   addr,
		Scheme:    "http",
		Transport: transportFn(),
	}

	if addr := os.Getenv(capi.HTTPAddrEnvName); addr != "" {
		config.Address = addr
	}

	if tokenFile := os.Getenv(capi.HTTPTokenFileEnvName); tokenFile != "" {
		config.TokenFile = tokenFile
	}

	if token := os.Getenv(capi.HTTPTokenEnvName); token != "" {
		config.Token = token
	}

	if auth := os.Getenv(capi.HTTPAuthEnvName); auth != "" {
		var username, password string
		if strings.Contains(auth, ":") {
			split := strings.SplitN(auth, ":", 2)
			username = split[0]
			password = split[1]
		} else {
			username = auth
		}

		config.HttpAuth = &capi.HttpBasicAuth{
			Username: username,
			Password: password,
		}
	}

	if ssl := os.Getenv(capi.HTTPSSLEnvName); ssl != "" {
		enabled, err := strconv.ParseBool(ssl)
		if err != nil {
			logger.Warn(fmt.Sprintf("could not parse %s", capi.HTTPSSLEnvName), "error", err)
		}

		if enabled {
			config.Scheme = "https"
		}
	}

	if v := os.Getenv(capi.HTTPTLSServerName); v != "" {
		config.TLSConfig.Address = v
	}
	if v := os.Getenv(capi.HTTPCAFile); v != "" {
		config.TLSConfig.CAFile = v
	}
	if v := os.Getenv(capi.HTTPCAPath); v != "" {
		config.TLSConfig.CAPath = v
	}
	if v := os.Getenv(capi.HTTPClientCert); v != "" {
		config.TLSConfig.CertFile = v
	}
	if v := os.Getenv(capi.HTTPClientKey); v != "" {
		config.TLSConfig.KeyFile = v
	}
	if v := os.Getenv(capi.HTTPSSLVerifyEnvName); v != "" {
		doVerify, err := strconv.ParseBool(v)
		if err != nil {
			logger.Warn(fmt.Sprintf("could not parse %s", capi.HTTPSSLVerifyEnvName), "error", err)
		}
		if !doVerify {
			config.TLSConfig.InsecureSkipVerify = true
		}
	}

	if v := os.Getenv(capi.HTTPNamespaceEnvName); v != "" {
		config.Namespace = v
	}

	if v := os.Getenv(capi.HTTPPartitionEnvName); v != "" {
		config.Partition = v
	}

	return config
}
