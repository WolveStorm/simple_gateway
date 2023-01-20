package initial

import (
	"simple_gateway/global"
	"testing"
)

func TestInitDebug(t *testing.T) {
	InitConfigByPath(true, "D:\\simple_gateway\\conf\\dev\\base.yaml")
	config := global.DebugFullConfig
	if !(config.ZapConfig.MaxSize == 300) {
		t.Error("data error")
	}
	if !(config.ZapConfig.OtherPath == "/logs/other") {
		t.Error("data error")
	}
	if !(config.ZapConfig.ErrorPath == "/logs/error") {
		t.Error("data error")
	}
	if !(config.ZapConfig.MaxAge == 30) {
		t.Error("data error")
	}
	if !(config.ZapConfig.MaxBackup == 5) {
		t.Error("data error")
	}
}
