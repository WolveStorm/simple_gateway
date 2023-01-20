package initial

import (
	"bufio"
	"gateway_server/global"
	"github.com/spf13/viper"
	"log"
	"os"
)

func getDevBasePath() string {
	path := GetCurrentPath()
	return path + "/conf/dev/base.yaml"
}

func getProBasePath() string {
	path := GetCurrentPath()
	return path + "/conf/pro/base.yaml"
}

func getProProxyPath() string {
	path := GetCurrentPath()
	return path + "/conf/pro/proxy.yaml"
}

func getDevProxyPath() string {
	path := GetCurrentPath()
	return path + "/conf/dev/proxy.yaml"
}

func InitConfigByPath(dev bool, path string) {
	v := viper.New()
	v.SetConfigType("yaml")
	open, err := os.Open(path)
	if err != nil {
		log.Fatal("[Navi Gateway]	open file error,details:", err.Error())
		return
	}
	defer open.Close()
	err = v.ReadConfig(bufio.NewReader(open))
	if err != nil {
		log.Fatal("[Navi Gateway]	read config error,details:", err.Error())
		return
	}
	err = v.Unmarshal(global.DebugFullConfig)
	if err != nil {
		log.Fatal("[Navi Gateway]	unmarshal error,details:", err.Error())
		return
	}
}

func InitBaseConfig(dev bool) {
	v := viper.New()
	path := getProBasePath()
	if dev {
		path = getDevBasePath()
	}
	open, err := os.Open(path)
	if err != nil {
		log.Fatal("[Navi Gateway]	open file error,details:", err.Error())
		return
	}
	defer open.Close()
	v.SetConfigType("yaml")
	err = v.ReadConfig(bufio.NewReader(open))
	if err != nil {
		log.Fatal("[Navi Gateway]	read config error,details:", err.Error())
		return
	}
	err = v.Unmarshal(global.DebugFullConfig)
	if err != nil {
		log.Fatal("[Navi Gateway]	unmarshal error,details:", err.Error())
		return
	}
}

func InitProxyConfig(dev bool) {
	v := viper.New()
	path := getProProxyPath()
	if dev {
		path = getDevProxyPath()
	}
	open, err := os.Open(path)
	if err != nil {
		log.Fatal("[Navi Gateway]	open file error,details:", err.Error())
		return
	}
	defer open.Close()
	v.SetConfigType("yaml")
	err = v.ReadConfig(bufio.NewReader(open))
	if err != nil {
		log.Fatal("[Navi Gateway]	read config error,details:", err.Error())
		return
	}
	err = v.Unmarshal(global.ProxyFullConfig)
	if err != nil {
		log.Fatal("[Navi Gateway]	unmarshal error,details:", err.Error())
		return
	}
}

func InitConfig(dev bool) {
	InitBaseConfig(dev)
	InitProxyConfig(dev)
}
