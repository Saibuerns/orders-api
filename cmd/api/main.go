package main

import (
	"orders.api/cmd/api/server"
)

const serverPort = ":8080"

func main() {
	urlMapping := server.NewMapping()
	_ = server.NewGinServer(urlMapping).Run(serverPort)
}
