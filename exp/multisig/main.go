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
	"github.com/btcsuite/btcd/txscript"
	"github.com/moroz/bitcoin-rpc-go/config"
	"golang.org/x/crypto/argon2"
)

const ORDER_ID = "0196bdbf-9e66-741c-b0e2-20eccb5aa233"

func main() {
	uuid, err := hex.DecodeString(strings.ReplaceAll(ORDER_ID, "-", ""))
	if err != nil {
		log.Fatal(err)
	}
	privKeyBytes := argon2.IDKey(config.SECRET_KEY_BASE, uuid, config.ARGON2ID_TIME, config.ARGON2ID_MEMORY, config.ARGON2ID_PARALLELISM, 32)
	privKey, _ := btcec.PrivKeyFromBytes(privKeyBytes)

	serverPubKey, err := btcutil.NewAddressPubKey(privKey.PubKey().SerializeCompressed(), &chaincfg.RegressionNetParams)
	if err != nil {
		log.Fatal(err)
	}
	coldPubKey, _ := btcutil.NewAddressPubKey(config.COLD_PUBLIC_KEY, &chaincfg.RegressionNetParams)
	if err != nil {
		log.Fatal(err)
	}
	redeemScript, err := txscript.MultiSigScript([]*btcutil.AddressPubKey{coldPubKey, serverPubKey}, 2)

	fmt.Printf("%X (%d)\n", privKey.PubKey().SerializeCompressed(), len(privKey.PubKey().SerializeCompressed()))
	fmt.Printf("%X (%d)\n", config.COLD_PUBLIC_KEY, len(config.COLD_PUBLIC_KEY))
	fmt.Printf("%X (%d)\n", redeemScript, len(redeemScript))

	redeemProg := sha256.Sum256(redeemScript)

	addr, err := btcutil.NewAddressWitnessScriptHash(redeemProg[:], &chaincfg.RegressionNetParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addr)
}
