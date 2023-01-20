package load_balance

import "math/rand"

type RandomLoadBalance struct {
	IPS       []string
	curIndex  int
	conf      LoadBalanceCof // 服务发现中心
	ForbidIps []string       // 屏蔽的ip
}

func (r *RandomLoadBalance) Add(param ...string) error {
	r.IPS = append(r.IPS, param[0])
	return nil
}
func (r *RandomLoadBalance) Get(addr string) string {
	r.conf.GetActiveIPS() // 从注册中心拉取最新的IP
	c := len(r.IPS)
	r.curIndex = rand.Intn(c)
	return r.IPS[r.curIndex]
}

func (r *RandomLoadBalance) Update(addrs []string, weightMap map[string]string) {
	r.IPS = addrs
	return
}

func (r *RandomLoadBalance) SetConf(conf LoadBalanceCof) {
	r.conf = conf
}

func (r *RandomLoadBalance) SetForbid(addr []string) {
	r.ForbidIps = addr
}
func (r *RandomLoadBalance) GetForbid() []string {
	return r.ForbidIps
}
