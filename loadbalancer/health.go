package loadbalancer

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/anuzx/load_balancer/servers"
)

type HealthChecker struct {
	Servers  []*servers.Server
	Interval time.Duration
	Client   *http.Client
}

func NewHealthChecker(
	servers []*servers.Server,
	interval time.Duration,
) *HealthChecker {
	return &HealthChecker{
		Servers:  servers,
		Interval: interval,
		Client: &http.Client{
			Timeout: 2 * time.Second,
			//if the health check doesn't complete within 2 seconds, consider it failed
		},
	}
}

//imagine a server alive but stuck ,without a timeout our health checker could wait forever

// connection error -> unhealthy
// response but anything other than 200 -> unhealthy
func (hc *HealthChecker) check(server *servers.Server) {
	url := server.URL.String() + "/health"

	resp, err := hc.Client.Get(url)

	if err != nil {
		server.SetHealthy(false)

		fmt.Printf("%s unhealthy: %v\n", url, err)

		return
	}

	//when this fn finishes, close the response body
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		server.SetHealthy(true)
		fmt.Printf("%s healthy\n", url)
		return
	}

	server.SetHealthy(false)

	fmt.Printf(
		"%s unhealthy: status %d\n",
		url,
		resp.StatusCode,
	)
}

func (hc *HealthChecker) checkAll() {
	for _, server := range hc.Servers {
		go hc.check(server)
	}
}

func (hc *HealthChecker) Start(ctx context.Context) {
	ticker := time.NewTicker(hc.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			hc.checkAll()

		case <-ctx.Done():
			return
		}
	}
}
