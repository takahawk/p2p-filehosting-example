package main

import (
	"fmt"
	"strconv"
)

type Peer struct {
	address string
	propagated bool
}
type PeerList struct {
	peers []Peer
}

type peerCallback func(string)

func AddPeer(list *PeerList, address string) {
	// TODO: implement checking if it contained already


	addr, _ := SplitAddress(address)
	addr = addr + ":" + strconv.Itoa(SERVER_PORT)

	fmt.Println("Address: " + addr)
	for _, n := range list.peers {
		fmt.Println("Peer: "+ n.address)
		if n.address == addr {
			return
		}
	}

	fmt.Println("Adding new peer to list: " + addr)
	list.peers = append(list.peers, Peer { addr, false })
}

func PropagatePeer(list *PeerList, address string, callback peerCallback) {
	AddPeer(list, address)
	addr, _ := SplitAddress(address)
	address = addr + ":" + strconv.Itoa(SERVER_PORT)

	for _, n := range list.peers {
		if n.address == address && !n.propagated {
			n.propagated = true
			callback(address)
		}
	}

}

