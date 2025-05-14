package main

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/moroz/bitcoin-rpc-go/config"
	"golang.org/x/crypto/argon2"
)

func main() {
	seed := argon2.IDKey(config.SECRET_KEY_BASE, []byte("wallet"), config.ARGON2ID_MEMORY, config.ARGON2ID_TIME, config.ARGON2ID_PARALLELISM, hdkeychain.RecommendedSeedLen)

	priv, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}

	// m/84'
	key, err := priv.Derive(hdkeychain.HardenedKeyStart + 84)
	// m/84'/0'
	key, err = key.Derive(hdkeychain.HardenedKeyStart)
	// m/84'/0'/0'
	key, err = key.Derive(hdkeychain.HardenedKeyStart)
	parent := key
	// m/84'/0'/0'/0
	key, err = key.Derive(0)
	// m/84'/0'/0'/0/0
	key, err = key.Derive(0)

	pubKey, _ := key.ECPubKey()
	privKey, _ := key.ECPrivKey()
	witnessProg := btcutil.Hash160(pubKey.SerializeCompressed())
	addr, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%X\n", pubKey.SerializeCompressed())
	fmt.Println(parent.String())
	fmt.Println(addr.EncodeAddress())
	fmt.Println(base58.Encode(privKey.Serialize()))

	wif, err := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, true)
	fmt.Println(wif.String())
}
