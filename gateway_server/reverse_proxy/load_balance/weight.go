package load_balance

import (
	"errors"
	"strconv"
)

type WeightNode struct {
	Addr            string
	Weight          int64
	EffectiveWeight int64
	CurrentWeight   int64
}

type WeightRoundLoadBalance struct {
	IPS       []*WeightNode
	curIndex  int
	conf      LoadBalanceCof // 服务发现中心
	ForbidIps []string       // 屏蔽的ip
}

func (w *WeightRoundLoadBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("params need more 2")
	}
	parseInt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return errors.New("convert to int error")
	}
	newNode := &WeightNode{
		Addr:            params[0],
		Weight:          parseInt,
		EffectiveWeight: parseInt,
		CurrentWeight:   parseInt,
	}
	w.IPS = append(w.IPS, newNode)
	return nil
}
func (w *WeightRoundLoadBalance) Get(addr string) string {
	w.conf.GetActiveIPS()
	var total int64 = 0
	var max int64 = 0
	var maxIndex = 0
	for i, v := range w.IPS {
		total += v.CurrentWeight
		w.IPS[i].CurrentWeight += v.EffectiveWeight
		if w.IPS[i].CurrentWeight > max {
			max = w.IPS[i].CurrentWeight
			maxIndex = i
		}
	}
	w.IPS[maxIndex].CurrentWeight -= total
	return w.IPS[maxIndex].Addr
}

func (w *WeightRoundLoadBalance) Update(addrs []string, weightMap map[string]string) {
	w.IPS = make([]*WeightNode, 0)
	for _, v := range addrs {
		if weight, ok := weightMap[v]; ok {
			w.Add(v, weight)
		} else {
			// 默认50
			w.Add(v, "50")
		}
	}
}

func (w *WeightRoundLoadBalance) SetConf(conf LoadBalanceCof) {
	w.conf = conf
}
func (w *WeightRoundLoadBalance) SetForbid(addr []string) {
	w.ForbidIps = addr
}
func (w *WeightRoundLoadBalance) GetForbid() []string {
	return w.ForbidIps
}
