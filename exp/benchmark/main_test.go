package main

import (
	"fmt"
	"testing"

	"github.com/moroz/bitcoin-rpc-go/config"
	"golang.org/x/crypto/argon2"
)

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
				argon2.IDKey(config.SECRET_KEY_BASE, []byte("wallet"), example.memory, example.time, uint8(example.parallelism), 32)
			}
		})
	}
}

func BenchmarkArgon2id_12_3_1(b *testing.B) {
	for b.Loop() {
		argon2.IDKey(config.SECRET_KEY_BASE, []byte("wallet"), 12, 3, 1, 32)
	}
}
