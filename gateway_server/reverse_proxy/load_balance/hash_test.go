package load_balance

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	balance := NewHashConsistentLoadBalance(5, nil)
	balance.Add("5")
	balance.Add("6")
	fmt.Println(balance.Get("1"))
	fmt.Println(balance.Get("2"))
	fmt.Println(balance.Get("3"))
	fmt.Println(balance.Get("4"))
}
