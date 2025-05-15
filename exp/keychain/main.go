package main

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/moroz/bitcoin-rpc-go/config"
	"golang.org/x/crypto/argon2"
)

func pubKeyToSegWitAddress(key *btcec.PublicKey) (string, error) {
	witnessProg := btcutil.Hash160(key.SerializeCompressed())
	addr, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}
	return addr.EncodeAddress(), nil
}

func deriveNeuteredBaseWalletFromMaster(master *hdkeychain.ExtendedKey) (key *hdkeychain.ExtendedKey, err error) {
	// m/84'
	if key, err = master.Derive(hdkeychain.HardenedKeyStart + 84); err != nil {
		return
	}
	// m/84'/0'
	if key, err = key.Derive(hdkeychain.HardenedKeyStart); err != nil {
		return
	}
	// m/84'/0'/0'
	if key, err = key.Derive(hdkeychain.HardenedKeyStart); err != nil {
		return
	}
	// m/84'/0'/0'/0
	if key, err = key.Derive(0); err != nil {
		return
	}
	return key.Neuter()
}

func main() {
	seed := argon2.IDKey(config.SECRET_KEY_BASE, []byte("wallet"), config.ARGON2ID_MEMORY, config.ARGON2ID_TIME, config.ARGON2ID_PARALLELISM, hdkeychain.RecommendedSeedLen)

	priv, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}

	key, err := deriveNeuteredBaseWalletFromMaster(priv)
	if err != nil {
		log.Fatal(err)
	}

	pubKey, _ := key.ECPubKey()
	addr, err := pubKeyToSegWitAddress(pubKey)
	fmt.Println("Derived from public:", addr)
}
