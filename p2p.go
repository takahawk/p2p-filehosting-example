package main

type PeerList struct {
	peers []string
}

func AddPeer(list PeerList, address string) {
	// TODO: implement checking if it contained already
	list.peers = append(list.peers, address)
}