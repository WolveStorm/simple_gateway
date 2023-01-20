package main

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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("请求来了")
	fmt.Fprintf(w, "Hi, This is an example of http service in golang!")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of http service in golang2222!")
}

func main() {
	http.HandleFunc("/h2", handler2)
	http.HandleFunc("/test_http_service", handler)
	reg := &capi.AgentServiceRegistration{
		Name:    "real_server",
		ID:      "real_server",
		Tags:    []string{"real_server"},
		Address: "127.0.0.1:2003",
	}
	client, _ := capi.NewClient(defaultConfig(nil, cleanhttp.DefaultPooledTransport, "127.0.0.1:8500"))
	agent := client.Agent()
	if err := agent.ServiceRegister(reg); err != nil {
		fmt.Println(err)
	}
	err := http.ListenAndServeTLS(":2003", "D:\\simple_gateway\\real_server\\https\\ca\\server.crt", "D:\\simple_gateway\\real_server\\https\\ca\\server.key", nil)
	if err != nil {
		fmt.Println(err)
	}
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
