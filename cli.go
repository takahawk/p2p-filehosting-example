package main

import (
	"fmt"
	"bufio"
	"os"
)


const PROMPT = "p2p-fh> "

func Interact() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(PROMPT)
		text, _ := reader.ReadString('\n')
		HandleCommand(text[:len(text) - 1])
	}
}

func HandleCommand(cmd string) {
	switch cmd {
	case "show peers":
		// TODO: implement communicating between goroutines
		fmt.Println("Peers:")
		fmt.Println(peerList)
	}
}
