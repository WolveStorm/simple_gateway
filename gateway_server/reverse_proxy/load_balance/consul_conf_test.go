package load_balance

import (
	"fmt"
	"testing"
)

func TestConsulConf_GetActiveIPS(t *testing.T) {
	conf := NewConsulConf(GenerateLoadBalance(1), "", "real_server", "127.0.0.1:8500")
	fmt.Println(conf.GetActiveIPS())
}
