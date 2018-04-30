package main

import (
	"crypto/rsa"
	"crypto/rand"
	"fmt"
)

const BIT_SIZE = 1024

func GenerateKey() *rsa.PrivateKey {
	reader := rand.Reader

	key, err := rsa.GenerateKey(reader, BIT_SIZE)
	if err != nil {
		fmt.Println("Error generating RSA key")
	}

	return key
}
