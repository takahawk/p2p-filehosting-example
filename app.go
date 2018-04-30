package main

import (
	"os"
	"fmt"
	"net"
	"encoding/json"
	"strconv"
	"time"
)

var peerList = make([]Peer, 0)

const SERVER_PORT = 8000
const BUFFER_SIZE = 1024
const CLI_STARTUP_DELAY = 500

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		StartClient(args[0], SendAssocRequest)
	}

	go StartServer(SERVER_PORT, HandleRequest)
	time.Sleep(CLI_STARTUP_DELAY * time.Millisecond)
	Interact()
}

func HandleRequest(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, BUFFER_SIZE)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Reading from socket error")
		return
	}

	var m RequestMeta
	err = json.Unmarshal(buf[:n], &m)

	if err != nil {
		fmt.Println("JSON decoding error")
		return
	}

	switch m.Code {

	case REQ_CODE_ASSOC:
		var m AssocRequest
		err := json.Unmarshal(buf[:n], &m)
		if err != nil {
			fmt.Println("JSON decoding error")
		}

		response := AssocResponse{append(peerList, GetSelf(conn.LocalAddr().String()))}
		bytes, err := json.Marshal(response)
		if err != nil {
			fmt.Println("JSON encoding error")
			return
		}
		conn.Write(bytes)
		AddAndPropagatePeer(&peerList, m.Peer, func(peer Peer) { StartClient(peer.IpAddress, GetPropagCallback(m.Peer))})
		fmt.Println(peerList)

	case REQ_CODE_PROPAG:
		fmt.Println("Propagation")
		var m PropagationRequest
		err := json.Unmarshal(buf[:n], &m)
		if err != nil {
			fmt.Println("JSON decoding error")
		}

		AddAndPropagatePeer(&peerList, m.Peer, func(peer Peer) { StartClient(peer.IpAddress, GetPropagCallback(m.Peer))})
		fmt.Println(peerList)
	}
}

func SendAssocRequest(conn net.Conn) {
	defer conn.Close()

	addr, _ := SplitAddress(conn.LocalAddr().String())
	addr = addr + ":" + strconv.Itoa(SERVER_PORT)

	buf := make([]byte, BUFFER_SIZE)
	request := MakeAssocRequest(GetSelf(addr))
	b, err := json.Marshal(request)
	if err != nil {
		fmt.Println("JSON encoding error")
		return
	}
	conn.Write(b)

	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Reading error")
		return
	}

	var m AssocResponse
	err = json.Unmarshal(buf[:n], &m)

	if err != nil {
		fmt.Println("JSON decoding error")
		return
	}

	for _, peer := range(m.KnownPeers) {
		AddPeer(&peerList, peer)
	}
}

func GetPropagCallback(peer Peer) func(net.Conn) {
	return func(conn net.Conn) {
		defer conn.Close()

		request := MakePropagationRequest(peer)
		buf := make([]byte, BUFFER_SIZE)
		b, err := json.Marshal(request)
		if err != nil {
			fmt.Println("JSON encoding error")
			return
		}
		conn.Write(b)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Reading error")
			return
		}

		var response PropagationResponse
		err = json.Unmarshal(buf[:n], &response)

		if err != nil {
			fmt.Println("JSON decoding error")
			return
		}
	}
}