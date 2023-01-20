package load_balance

type RoundLoadBalance struct {
	IPS       []string
	curIndex  int
	conf      LoadBalanceCof // 服务发现中心
	ForbidIps []string       // 屏蔽的ip
}

func (r *RoundLoadBalance) Add(param ...string) error {
	r.IPS = append(r.IPS, param[0])
	return nil
}
func (r *RoundLoadBalance) Get(addr string) string {
	r.conf.GetActiveIPS() // 从注册中心拉取最新的IP
	// 防止正好服务发现中心下线了一个实例
	if r.curIndex >= len(r.IPS) {
		r.curIndex = 0
	}
	ret := r.IPS[r.curIndex]
	r.curIndex = (r.curIndex + 1) % len(r.IPS)
	return ret
}
func (r *RoundLoadBalance) Update(addrs []string, weightMap map[string]string) {
	r.IPS = addrs
	return
}

func (r *RoundLoadBalance) SetConf(conf LoadBalanceCof) {
	r.conf = conf
}
func (r *RoundLoadBalance) SetForbid(addr []string) {
	r.ForbidIps = addr
}
func (r *RoundLoadBalance) GetForbid() []string {
	return r.ForbidIps
}
