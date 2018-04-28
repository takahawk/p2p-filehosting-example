package main


const MSG_CODE_ASSOC = 1 // association message, sent when nodes establish communication
const MSG_CODE_PROPAG = 2 // propagation message, sent when peer tell another peer about third one
type Message struct {
	Code int
	Payload []byte
}

const RESPONSE_CODE_OK = 0