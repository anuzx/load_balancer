package main

import (
	"os"

	"github.com/anuzx/load_balancer/servers"
)

func main() {
	port := os.Getenv("PORT")
	id := os.Getenv("SERVER_ID")

	servers.Start(port, id)
}
