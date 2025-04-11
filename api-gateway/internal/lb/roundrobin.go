package lb

import (
	"sync/atomic"
)

// RoundRobin holds a list of target URLs and an atomic counter.
type RoundRobin struct {
	targets []string
	idx     uint64
}

// New creates a RoundRobin with the given targets.
func New(targets []string) *RoundRobin {
	return &RoundRobin{targets: targets}
}

// Next returns the next target in round‑robin fashion.
func (r *RoundRobin) Next() string {
	i := atomic.AddUint64(&r.idx, 1)
	return r.targets[(int(i)-1)%len(r.targets)]
}
