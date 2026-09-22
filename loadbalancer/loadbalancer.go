package loadbalancer

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/anuzx/load_balancer/algorithms"
	"github.com/anuzx/load_balancer/servers"
)

//our lb doesn't says round robin anywhere ,lb only knows "give me something that can select a server"

type LoadBalancer struct {
	Servers       []*servers.Server
	Selector      algorithms.Selector
	HealthChecker *HealthChecker
}

func NewLoadBalancer(
	servers []*servers.Server,
	selector algorithms.Selector,
) *LoadBalancer {
	return &LoadBalancer{
		Servers:  servers,
		Selector: selector,
		HealthChecker: NewHealthChecker(
			servers,
			2*time.Second,
		),
	}
}

func (lb *LoadBalancer) handler(w http.ResponseWriter, r *http.Request) {
	server := lb.Selector.Next(lb.Servers)

	if server == nil {
		http.Error(
			w,
			"No healthy servers available",
			http.StatusServiceUnavailable,
		)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(server.URL)

	proxy.ServeHTTP(w, r)
}

func (lb *LoadBalancer) Start(ctx context.Context) {
	http.HandleFunc("/", lb.handler)

	go lb.HealthChecker.Start(ctx)

	fmt.Println("Load balancer running on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("load balancer error:", err)
	}
}

func NewServer(rawURL string) *servers.Server {
	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		panic(err)
	}

	server := &servers.Server{
		URL: parsedURL,
	}

	server.SetHealthy(true)

	return server
}
