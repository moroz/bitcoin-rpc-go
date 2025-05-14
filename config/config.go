package config

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"os"
)

func MustGetenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Environment variable %s is not set!", key)
	}
	return val
}

func MustGetenvBase64(key string) []byte {
	val := MustGetenv(key)
	bytes, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		log.Fatalf("Failed to decode environment variable %s from Base64: %s", key, err)
	}
	return bytes
}

func MustGetenvHex(key string) []byte {
	val := MustGetenv(key)
	bytes, err := hex.DecodeString(val)
	if err != nil {
		log.Fatalf("Failed to decode environment variable %s from Hex: %s", key, err)
	}
	return bytes
}

var DATABASE_URL = MustGetenv("DATABASE_URL")
var SECRET_KEY_BASE = MustGetenvBase64("SECRET_KEY_BASE")
var COLD_PUBLIC_KEY = MustGetenvHex("COLD_PUBLIC_KEY")

// These parameters are required to generate secret keys for orders
// Do not change these once in production, otherwise the keys will change
const ARGON2ID_MEMORY = 46 * 1024
const ARGON2ID_TIME = 1
const ARGON2ID_PARALLELISM = 1
