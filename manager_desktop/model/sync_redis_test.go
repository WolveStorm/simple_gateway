package model

import (
	"simple_gateway/initial"
	"testing"
)

func TestSyncToRedis(t *testing.T) {
	initial.InitConfigByPath(true, "D:\\simple_gateway\\gateway_server\\conf\\dev\\base.yaml")
	initial.InitRedis()
	initial.InitMysql()
	SyncToRedis()
}
