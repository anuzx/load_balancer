package main

import (
	"github.com/anuzx/load_balancer/algorithms"
	"github.com/anuzx/load_balancer/loadbalancer"
	"github.com/anuzx/load_balancer/servers"
)

func main() {
	servers := []*servers.Server{
		loadbalancer.NewServer("http://localhost:8081"),
		loadbalancer.NewServer("http://localhost:8082"),
		loadbalancer.NewServer("http://localhost:8083"),
	}

	selector := algorithms.NewRoundRobin()

	//we are giving load balancer the list of servers and the algo to use
	lb := loadbalancer.NewLoadBalancer(
		servers,
		selector,
	)

	lb.Start()
}
