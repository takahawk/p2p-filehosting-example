package main

import (
	"fmt"
)

type Peer struct {
	// TODO: add public key
	IpAddress  string
}

type PeerList struct {
	peers []Peer
}

type peerCallback func(Peer)

func AddPeer(list *[]Peer, peer Peer) {
	fmt.Println("Adding new peer to list: " + peer.IpAddress)
	*list = append(*list, peer)
}

func IsPeerKnown(list *[]Peer, peer Peer) bool {
	for _, p := range *list {
		// FIXME: check by public key
		if p.IpAddress == peer.IpAddress {
			return true
		}
	}
	return false
}

func AddAndPropagatePeer(list *[]Peer, peer Peer, callback peerCallback) {
	// if peer is already known - we don't propagate it further
	if IsPeerKnown(list, peer) || peer.IpAddress == self.IpAddress {
		return
	}
	prevList := *list
	AddPeer(list, peer)

	for _, knownPeer := range prevList {
			callback(knownPeer)
	}
}

func WithSelf(list PeerList, useIpAddr string) PeerList {
	self.IpAddress = useIpAddr
	return PeerList {append(list.peers, self ) }
}

func GetSelf(useIpAddr string) Peer {
	self.IpAddress = useIpAddr
	return self
}

var self = Peer {
	// TODO: good place for public key
}