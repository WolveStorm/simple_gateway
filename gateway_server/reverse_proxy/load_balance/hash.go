package load_balance

import (
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type Hash func([]byte) uint32

type Uint32Slice []uint32

func (s Uint32Slice) Len() int {
	return len(s)
}
func (s Uint32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Uint32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type HashConsistentLoadBalance struct {
	h          Hash
	round      Uint32Slice
	m          map[uint32]string
	replicates int // 避免hash偏斜，创建副本
	mux        sync.Mutex
	conf       LoadBalanceCof // 服务发现中心
	ForbidIps  []string       // 屏蔽的ip
}

func NewHashConsistentLoadBalance(replicas int, hash Hash) *HashConsistentLoadBalance {
	c := &HashConsistentLoadBalance{
		h:          hash,
		replicates: replicas,
		m:          make(map[uint32]string),
		round:      make([]uint32, 0, 100),
	}

	if c.h == nil {
		//保证是一个2^32 - 1的一个环
		c.h = crc32.ChecksumIEEE
	}
	return c
}
func (h *HashConsistentLoadBalance) Add(param ...string) error {
	for i := 0; i < h.replicates; i++ {
		key := h.h([]byte(strconv.Itoa(i) + param[0]))
		h.round = append(h.round, key)
		h.m[key] = param[0]
	}
	sort.Sort(h.round)
	return nil
}
func (h *HashConsistentLoadBalance) Get(addr string) string {
	h.conf.GetActiveIPS() // 从注册中心拉取最新的IP
	hash := h.h([]byte(addr))
	idx := sort.Search(len(h.round), func(i int) bool {
		return h.round[i] > hash
	})
	if idx == len(h.round) {
		idx = 0
	}
	h.mux.Lock()
	defer h.mux.Unlock()
	val := h.m[h.round[idx]]
	return val
}

func (h *HashConsistentLoadBalance) Update(addrs []string, weightMap map[string]string) {
	h.m = make(map[uint32]string)
	h.round = make([]uint32, 0)
	for _, v := range addrs {
		h.Add(v)
	}
}

func (h *HashConsistentLoadBalance) SetConf(conf LoadBalanceCof) {
	h.conf = conf
}

func (h *HashConsistentLoadBalance) SetForbid(addr []string) {
	h.ForbidIps = addr
}
func (h *HashConsistentLoadBalance) GetForbid() []string {
	return h.ForbidIps
}
