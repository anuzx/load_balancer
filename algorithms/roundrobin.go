package algorithms

import (
	"sync"

	"github.com/anuzx/load_balancer/servers"
)

// algo owns its own state
type RoundRobin struct {
	current int //its not the property of loadbalancer ,its the property of roundrobin algo
	mu      sync.Mutex
}

// factory fn
func NewRoundRobin() *RoundRobin {
	return &RoundRobin{}
}

// *RoundRobin instead of RoundRobin because next modifies rr.current (we want to modify actual round robin object)
func (rr *RoundRobin) Next(pool []*servers.Server) *servers.Server {
	//because multiple http reqs can call
	rr.mu.Lock()
	defer rr.mu.Unlock()

	if len(pool) == 0 {
		return nil
	}

	for range pool {
		index := rr.current % len(pool)
		server := pool[index]

		rr.current++

		if server.IsHealthy() {
			return server
		}
	}

	return nil

}
