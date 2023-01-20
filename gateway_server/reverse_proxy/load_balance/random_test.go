package load_balance

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	balance := RandomLoadBalance{
		IPS:      []string{"1", "2", "3", "4"},
		curIndex: 0,
	}
	balance.Add("5")
	fmt.Println(balance.Get(""))
	fmt.Println(balance.Get(""))
	fmt.Println(balance.Get(""))
	fmt.Println(balance.Get(""))
}
