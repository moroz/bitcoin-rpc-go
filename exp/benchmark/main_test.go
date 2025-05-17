package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"testing"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"golang.org/x/crypto/argon2"
)

var SECRET_KEY_BASE, _ = base64.StdEncoding.DecodeString("adkwd6hlANolxxSN/oV6BVkcC5wfeQy1hjO0Dcx8yRwzkZed6Gr7ljmY84whL/A+fgcxn9ZzkZc3cGWNw9cGWA==")

var params = []struct{ memory, time, parallelism uint32 }{
	{64, 3, 4},
	{46, 1, 8},
	{46, 1, 4},
	{46, 1, 1},
	{19, 2, 1},
	{12, 3, 1},
}

func BenchmarkArgon2id(b *testing.B) {
	for _, example := range params {
		b.Run(fmt.Sprintf("m=%d,t=%d,p=%d", example.memory, example.time, example.parallelism), func(b *testing.B) {
			for b.Loop() {
				argon2.IDKey(SECRET_KEY_BASE, []byte("wallet"), example.memory, example.time, uint8(example.parallelism), 32)
			}
		})
	}
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

func BenchmarkDeriveWallet(b *testing.B) {
	b.Run("Derive master from seed", func(b *testing.B) {
		seed := argon2.IDKey(SECRET_KEY_BASE, []byte("wallet"), 64, 3, 4, 32)
		for b.Loop() {
			_, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Derive public key from neutered key", func(b *testing.B) {
		seed := argon2.IDKey(SECRET_KEY_BASE, []byte("wallet"), 64, 3, 4, 32)
		master, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
		if err != nil {
			b.Fatal(err)
		}
		base, err := deriveKey(
			master,
			hdkeychain.HardenedKeyStart+84,
			hdkeychain.HardenedKeyStart,
			hdkeychain.HardenedKeyStart,
			0,
		)
		if err != nil {
			b.Fatal(err)
		}
		seg1b, _ := rand.Int(rand.Reader, big.NewInt(0x7FFFFFFF))
		seg2b, _ := rand.Int(rand.Reader, big.NewInt(0x7FFFFFFF))
		base, err = base.Neuter()
		if err != nil {
			b.Fatal(err)
		}

		seg1 := uint32(seg1b.Int64())
		seg2 := uint32(seg2b.Int64())

		for b.Loop() {
			_, err := deriveKey(base, seg1, seg2)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
