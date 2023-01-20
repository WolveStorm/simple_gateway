package initial

import (
	"go.uber.org/zap"
	"testing"
)

func TestZapInit(t *testing.T) {
	InitConfigByPath(true, "D:\\simple_gateway\\conf\\dev\\base.yaml")
	InitAllZap()
	t.Run("error", func(t *testing.T) {
		zap.S().Error("这是一条测试日志")
	})
	t.Run("other", func(t *testing.T) {
		zap.S().Infof("这是一条测试日志")
		zap.S().Warnf("这是一条测试日志")
	})
}
