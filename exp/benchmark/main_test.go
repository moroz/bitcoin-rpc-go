package main

import (
	"testing"

	"github.com/moroz/bitcoin-rpc-go/config"
	"golang.org/x/crypto/argon2"
)

func BenchmarkArgon2id(b *testing.B) {
	for b.Loop() {
		argon2.IDKey(config.SECRET_KEY_BASE, []byte("wallet"), config.ARGON2ID_MEMORY, config.ARGON2ID_TIME, config.ARGON2ID_PARALLELISM, 32)
	}
}
