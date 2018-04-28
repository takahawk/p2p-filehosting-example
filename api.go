package main


const MSG_CODE_ASSOC = 1 // association message, sent when nodes establish communication

type Message struct {
	Code int
	Payload []byte
}

const RESPONSE_CODE_OK = 0