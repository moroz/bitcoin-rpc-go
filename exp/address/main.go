package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/moroz/bitcoin-rpc-go/config"
	"golang.org/x/crypto/argon2"
)

var ORDER_IDS = []string{
	"0196bdbf-9e66-741c-b0e2-20eccb5aa233",
	"0196d381-0c2b-74e8-9dc5-be8ce0962e96",
	"0196d3a1-5b06-7a59-bbce-6b2b1037dfc7",
}

var HMAC_WALLET_DERIVATION_KEY = argon2.IDKey(config.SECRET_KEY_BASE, []byte("hmac"), config.ARGON2ID_MEMORY, config.ARGON2ID_TIME, config.ARGON2ID_PARALLELISM, 32)

func walletIndexFromBytes(bytes []byte) uint32 {
	var val uint32
	binary.Decode(bytes, binary.LittleEndian, &val)
	return val & 0x7F_FF_FF_FF
}

func calculateHMAC(bytes []byte) []byte {
	mac := hmac.New(sha256.New, HMAC_WALLET_DERIVATION_KEY)
	mac.Write(bytes)
	return mac.Sum(nil)
}

func walletPath(orderID []byte) (uint32, uint32) {
	sum := calculateHMAC(orderID)
	return walletIndexFromBytes(sum[:4]), walletIndexFromBytes(sum[4:8])
}

func mustParseUUID(uuid string) []byte {
	parsed, err := hex.DecodeString(strings.ReplaceAll(uuid, "-", ""))
	if err != nil {
		panic(err)
	}
	return parsed
}

func pubKeyToSegWitAddress(key *btcec.PublicKey) (string, error) {
	witnessProg := btcutil.Hash160(key.SerializeCompressed())
	addr, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}
	return addr.EncodeAddress(), nil
}

func deriveKey(baseKey *hdkeychain.ExtendedKey, paths ...uint32) (key *hdkeychain.ExtendedKey, err error) {
	key = baseKey
	for _, segment := range paths {
		key, err = key.Derive(segment)
		if err != nil {
			return
		}
	}
	return
}

func main() {
	master, err := hdkeychain.NewMaster(config.WALLET_SEED, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.WALLET_ORDER_BASE_PUBKEY)
	baseKey, err := hdkeychain.NewKeyFromString(config.WALLET_ORDER_BASE_PUBKEY)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HMAC Key: %X\n", HMAC_WALLET_DERIVATION_KEY)

	for _, id := range ORDER_IDS {
		uuid := mustParseUUID(id)
		fmt.Printf("\nUUID: %X\n", uuid)
		fmt.Printf("HMAC: %X\n", calculateHMAC(uuid))

		seg1, seg2 := walletPath(uuid)
		fmt.Printf("Derivation path: m/84'/0'/0'/0/%d/%d\n", seg1, seg2)

		fromMaster, err := deriveKey(
			master,
			hdkeychain.HardenedKeyStart+84,
			hdkeychain.HardenedKeyStart,
			hdkeychain.HardenedKeyStart,
			0,
			seg1, seg2,
		)
		if err != nil {
			log.Fatal(err)
		}

		pubKeyFromMaster, err := fromMaster.ECPubKey()
		if err != nil {
			log.Fatal(err)
		}

		key, err := deriveKey(baseKey, seg1, seg2)
		if err != nil {
			log.Fatal(err)
		}

		pubKey, err := key.ECPubKey()
		if err != nil {
			log.Fatal(err)
		}

		if pubKey.IsEqual(pubKeyFromMaster) {
			fmt.Println("Derived key from master works :+1:")
		} else {
			fmt.Printf("From master: %X\nFrom neutered: %X\n", pubKey.SerializeCompressed(), pubKeyFromMaster.SerializeCompressed())
		}

		fmt.Printf("Public key: %s\n", key.String())
		address, err := pubKeyToSegWitAddress(pubKey)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Address: %s\n", address)
	}

}
