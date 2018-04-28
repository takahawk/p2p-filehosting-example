package main

import (
	"os"
	"fmt"
	"net"
	"encoding/json"
)

var peerList = PeerList { make([]Peer, 0) }

const SERVER_PORT = 8000

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		StartClient(args[0], SendAssocRequest)
	}

	StartServer(SERVER_PORT, HandleRequest)
}

func HandleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Reading from socket error")
		return
	}
	fmt.Println(buf[:n])
	var m Message
	err = json.Unmarshal(buf[:n], &m)

	if err != nil {
		fmt.Println("JSON decoding error")
		return
	}

	fmt.Println(m)
	switch m.Code {

	case MSG_CODE_ASSOC:
		addr := conn.RemoteAddr().String()
		AddPeer(&peerList, addr)
		conn.Write([]byte { RESPONSE_CODE_OK })
		conn.Close()
		PropagatePeer(&peerList, addr, func(peerAddr string) { StartClient(peerAddr, GetPropapagCallback(addr))})

	case MSG_CODE_PROPAG:
		fmt.Println("Propagation")
		addr := string(m.Payload)
		AddPeer(&peerList, addr)
		conn.Write([]byte { RESPONSE_CODE_OK })
		conn.Close()
		PropagatePeer(&peerList, addr, func(peerAddr string) { StartClient(peerAddr, GetPropapagCallback(addr))})
	}
	conn.Close()
}

func SendAssocRequest(conn net.Conn) {
	buf := make([]byte, 1)
	request := &Message { MSG_CODE_ASSOC, make([]byte, 0)}
	b, err := json.Marshal(request)
	if err != nil {
		fmt.Println("JSON encoding error")
		return
	}
	conn.Write(b)
	conn.Read(buf)

	if buf[0] == RESPONSE_CODE_OK {
		addr := conn.RemoteAddr().String()
		AddPeer(&peerList, addr)
	}
}

func GetPropapagCallback(addr string) func(net.Conn) {
	return func(conn net.Conn) {
		buf := make([]byte, 1)
		request := &Message { MSG_CODE_PROPAG, []byte(addr)}
		b, err := json.Marshal(request)
		if err != nil {
			fmt.Println("JSON encoding error")
			return
		}
		conn.Write(b)
		conn.Read(buf)

		if buf[0] == RESPONSE_CODE_OK {
			AddPeer(&peerList, addr)
		}
	}
}