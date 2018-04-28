package main

import (
	"fmt"
	"net"
)

const SERVER_PORT = 8000

func main() {
	StartServer(SERVER_PORT, HandleRequest)
}

func HandleRequest(conn net.Conn) {
	// TODO: impl
	fmt.Println("Connection accepted")
}