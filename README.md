# Load Balancer

A simple load balancer built in Go to explore concurrency, reverse proxying, health checks, load-balancing algorithms

## Tech & Concepts Used

* **Go**
* **HTTP Server** — `net/http`
* **Reverse Proxy** — `net/http/httputil`
* **Round Robin Load Balancing**
* **Interfaces** — algorithm abstraction using `Selector`
* **Goroutines** — concurrent health checks
* **Channels** — used through Go concurrency primitives
* **Mutexes** — `sync.Mutex` and `sync.RWMutex`
* **Context** — cancellation and lifecycle management
* **Health Checks** — periodic backend health monitoring
* **Environment Variables**
* **Makefile**

## System Architecture

<img width="752" height="466" alt="image" src="https://github.com/user-attachments/assets/62b438f0-2f51-4695-9873-1b84217eb6dd" />


<!--
![System Architecture](./docs/system-architecture.png)
-->

## Project Structure

```text
.
├── algorithms
│   ├── algorithm.go
│   └── roundrobin.go
├── cmd
│   ├── loadbalancer
│   │   └── main.go
│   └── server
│       └── main.go
├── loadbalancer
│   ├── loadbalancer.go
│   └── health.go
├── servers
│    └── server.go
├── Makefile
├── README.md
└── go.mod
```

## Running Locally

### 1. Clone the repository

```bash
git clone <repository-url>
cd <repository-name>
```

### 2. Sync Go dependencies

```bash
go mod tidy
```

### 3. Run the backend servers

Open separate terminals and run:

```bash
make run-server-1
```

```bash
make run-server-2
```

```bash
make run-server-3
```

### 4. Run the load balancer

In another terminal:

```bash
make run-lb
```

The load balancer runs on:

```text
http://localhost:8080
```

The backend servers run on:

```text
http://localhost:8081
http://localhost:8082
http://localhost:8083
```

## Load Balancing

The load balancer currently uses **Round Robin** selection.

```text
Request 1 → Server 1
Request 2 → Server 2
Request 3 → Server 3
Request 4 → Server 1
...
```

Unhealthy servers are skipped by the load-balancing algorithm.

## Health Checking

Backend servers expose:

```text
GET /health
```

The load balancer periodically checks backend health and updates each server's health state.

## Future Improvements

* Retry/failover when a selected backend fails
* Additional algorithms such as Random and Least Connections
* Graceful shutdown
* More comprehensive tests
* Improved Docker service discovery
* Metrics and observability

