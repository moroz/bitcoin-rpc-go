package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/moroz/bitcoin-rpc-go/config"
	"golang.org/x/crypto/argon2"

	//nolint:staticcheck // RIPEMD160 required by Bitcoin HASH160
	"golang.org/x/crypto/ripemd160"
)

const ORDER_ID = "0196bdbf-9e66-741c-b0e2-20eccb5aa233"

func main() {
	uuid, err := hex.DecodeString(strings.ReplaceAll(ORDER_ID, "-", ""))
	if err != nil {
		log.Fatal(err)
	}
	privKeyBytes := argon2.IDKey(config.SECRET_KEY_BASE, uuid, config.ARGON2ID_TIME, config.ARGON2ID_MEMORY, config.ARGON2ID_PARALLELISM, 32)
	privKey, _ := btcec.PrivKeyFromBytes(privKeyBytes)
	pubKey := privKey.PubKey().SerializeCompressed()
	h := sha256.Sum256(pubKey)
	r := ripemd160.New()
	r.Write(h[:])
	pubKeyHash := r.Sum(nil)

	addr, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addr.EncodeAddress())

	fmt.Printf("%X\n", privKeyBytes)
	fmt.Printf("%X\n", pubKey)
}
