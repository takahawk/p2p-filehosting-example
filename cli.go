package main

import (
	"fmt"
)


const PROMPT = "p2p-fh> "

func Interact() {
	for {
		fmt.Print(PROMPT)
		var input string
		fmt.Scanln(&input)
		HandleCommand(input)
	}
}

func HandleCommand(cmd string) {
	switch cmd {
	case "show peers":
		// TODO: implement communicating between goroutines
		fmt.Println("Peers:")
		fmt.Println(peerList)
	default:
		fmt.Println()
	}
}
