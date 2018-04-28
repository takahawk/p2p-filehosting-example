package main

import (
	"fmt"
	"net"
	"strconv"
)

const CONN_TYPE = "tcp"
const CONN_HOST = "0.0.0.0"

type handler func(conn net.Conn)

func StartServer(port int, h handler) {
	l, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + strconv.Itoa(port))
	if err != nil {
		fmt.Println("Starting server error: ", err.Error())
		return
	}
	defer l.Close()

	fmt.Println("Server started")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			return
		}

		h(conn)
	}
}