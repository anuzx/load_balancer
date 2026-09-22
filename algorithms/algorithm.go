package algorithms

import "github.com/anuzx/load_balancer/servers"

// interface because any algorithm that can select a server must provide a Next() method
type Selector interface {
	Next(servers []*servers.Server) *servers.Server
}
