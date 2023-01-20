package load_balance

import (
	"fmt"
	"testing"
)

func TestWeight(t *testing.T) {
	list := make([]*WeightNode, 0)
	balance := WeightRoundLoadBalance{
		IPS:      list,
		curIndex: 0,
	}
	balance.Add("A", "4")
	balance.Add("B", "3")
	balance.Add("C", "2")
	fmt.Println(balance.Get(""))
	fmt.Println(balance.Get(""))
	fmt.Println(balance.Get(""))
	fmt.Println(balance.Get(""))
	fmt.Println(balance.Get(""))
	fmt.Println(balance.Get(""))
	fmt.Println(balance.Get(""))

}
