package initial

import (
	"simple_gateway/global"
	"testing"
)

func TestInitMysql(t *testing.T) {
	InitConfigByPath(true, "D:\\simple_gateway\\conf\\dev\\base.yaml")
	InitMysql()
	err := global.Conn.Ping()
	if err != nil {
		t.Error("conn error:", err)
	}
}
