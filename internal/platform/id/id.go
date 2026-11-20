package id

import (
	"crypto/rand"
	"encoding/hex"
)

// New returns an opaque, prefixed identifier such as "acc_9f..." with 128 bits
// of entropy. IDs never contain personal data.
func New(prefix string) string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic("crypto/rand unavailable: " + err.Error())
	}
	return prefix + "_" + hex.EncodeToString(b[:])
}
