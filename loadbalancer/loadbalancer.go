package loadbalancer

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server struct {
	URL string
}

var servers = []Server{
	{
		URL: "http://localhost:8081",
	},
	{
		URL: "http://localhost:8082",
	},
	{
		URL: "http://localhost:8083",
	},
}

var current = 0

func getNextServer() Server {
	server := servers[current]

	current = (current + 1) % len(servers)

	return server
}

func handler(w http.ResponseWriter, r *http.Request) {
	server := getNextServer()

	target, err := url.Parse(server.URL)

	if err != nil {
		http.Error(w, "invalid server URL", http.StatusInternalServerError)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ServeHTTP(w, r)
}

func Start() {
	http.HandleFunc("/", handler)

	fmt.Println("Load balancer running on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("load balancer error:", err)
	}
}
