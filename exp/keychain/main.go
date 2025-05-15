package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/moroz/bitcoin-rpc-go/config"
	"golang.org/x/crypto/argon2"
)

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

	fmt.Println("# For your local setup:")
	fmt.Printf("export WALLET_SEED=\"%s\"\n\n", base64.RawStdEncoding.EncodeToString(seed))

	fmt.Println("# For the server:")
	fmt.Printf("export WALLET_ORDER_BASE_PUBKEY=\"%s\"\n", key.String())
}
