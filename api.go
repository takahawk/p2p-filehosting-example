package main


const REQ_CODE_ASSOC = 1  // association message, sent when nodes establish communication
const REQ_CODE_PROPAG = 2 // propagation message, sent when peer tell another peer about third one

// used for identification of received messages
type RequestMeta struct {
	Code int
}

type AssocRequest struct {
	Code int
	Peer Peer
}

type AssocResponse struct {
	// TODO: add public key for file decryption/encryption
	KnownPeers []Peer
}

type PropagationRequest struct {
	Code int
	Peer Peer
}

type PropagationResponse struct {}


func MakeAssocRequest(self Peer) AssocRequest {
	return AssocRequest { REQ_CODE_ASSOC, self }
}

func MakePropagationRequest(peer Peer) PropagationRequest {
	return PropagationRequest { REQ_CODE_PROPAG, peer }
}