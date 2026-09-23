.PHONY: run-lb run-server-1 run-server-2 run-server-3

run-lb:
	go run ./cmd/loadbalancer

run-server-1:
	PORT=8081 SERVER_ID=server-1 go run ./cmd/server

run-server-2:
	PORT=8082 SERVER_ID=server-2 go run ./cmd/server

run-server-3:
	PORT=8083 SERVER_ID=server-3 go run ./cmd/server
