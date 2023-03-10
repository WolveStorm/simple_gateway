package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	capi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-hclog"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("请求来了")
	conn, _ := upgrader.Upgrade(w, r, nil)
	fmt.Println("client connected")
	w.Write([]byte("Hi, This is an example of http service in golang!"))
	conn.WriteMessage(websocket.TextMessage, []byte("hello client,this is message from websocket server"))
	Reader(conn)
}

func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
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

func main() {
	reg := &capi.AgentServiceRegistration{
		Name:    "real_server",
		ID:      "real_server",
		Tags:    []string{"real_server"},
		Address: "10.0.24.3:2003",
	}
	client, _ := capi.NewClient(defaultConfig(nil, cleanhttp.DefaultPooledTransport, "10.0.24.3:8500"))
	agent := client.Agent()
	if err := agent.ServiceRegister(reg); err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/test_http_service", handler)
	http.ListenAndServe("10.0.24.3:2003", nil)
}
