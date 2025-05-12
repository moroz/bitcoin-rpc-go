package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
)

func main() {
	privKeyBytes := make([]byte, 32)
	rand.Read(privKeyBytes)

	privKey, _ := btcec.PrivKeyFromBytes(privKeyBytes)
	pubKey := privKey.PubKey().SerializeCompressed()

	fmt.Printf("export COLD_PRIVATE_KEY=%s\n", base64.StdEncoding.EncodeToString(privKeyBytes))
	fmt.Printf("export COLD_PUBLIC_KEY=%X\n", pubKey)
}
